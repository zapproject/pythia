package webview

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ZapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/ops"
	"github.com/zapproject/pythia/setup"
	"github.com/zapproject/pythia/util"
)

func configWallet(cfg string) (bool, error) {
	// fmt.Println(req.TokenAddress )
	err := config.ParseConfigBytes([]byte(cfg))
	if err != nil {
		fmt.Println(err.Error())
	}

	return setup.App()
}

func getPubKey() string {
	addr := setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)

	return addr.String()[0:10]
}

func getBalance() (string, error) {
	addr := setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)
	instance := setup.CTX.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	zapBalance, err := instance.BalanceOf(nil, addr)

	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}

	return zapBalance.String(), nil

}

func stakeStatus() (int64, error) {

	tmaster := setup.CTX.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)

	publicAddress := setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)
	status, _, err := tmaster.GetStakerInfo(nil, publicAddress)
	if err != nil {
		return 0, fmt.Errorf("failed to get stake status: %s", err.Error())
	}

	return status.Int64(), nil

}

func depositStake() (bool, error) {
	err := ops.Deposit(setup.CTX)
	if err != nil {
		return false, err
	}
	return true, nil
}

func requestWithdraw() (bool, error) {
	err := ops.RequestStakingWithdraw(setup.CTX)
	if err != nil {
		return false, err
	}
	return true, nil
}

func withdrawStatus() (string, bool, error) {
	tmaster := setup.CTX.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	publicAddress := setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)
	_, startTime, err := tmaster.GetStakerInfo(nil, publicAddress)
	if err != nil {
		return "", false, err
	}
	startedRound := startTime.Int64()
	startedRound = ((startedRound + 86399) / 86400) * 86400
	target := time.Unix(startedRound, 0)
	timePassed := time.Now().Sub(target)
	delta := timePassed - (time.Hour * 24 * 7)
	// delta := 1

	if delta > 0 {
		return "Stake has been eligbile to withdraw for " + fmt.Sprint(delta), true, nil
	} else {
		return "Stake will be eligible to withdraw in " + fmt.Sprint(-delta), false, nil
	}

}

func canWithdraw() (bool, error) {

	_, canWithdraw, err := withdrawStatus()
	if err != nil {
		return false, err
	}
	return canWithdraw, nil

}

func withdrawMsg() (string, error) {

	msg, _, err := withdrawStatus()
	if err != nil {
		return "", err
	}
	return msg, nil

}

func withdraw() (bool, error) {
	err := ops.WithdrawStake(setup.CTX)
	if err != nil {
		return false, err
	}
	return true, nil
}

var mineCh chan string
var mining bool = false
var initializing bool = false

func startMine() {

	if mining == false && initializing == false {
		go func() {
			initializing = true
			file, _ := os.OpenFile("./webview/public/miningLogs.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

			os.Stdout = file
			os.Stderr = file
			log.SetOutput(file)
			loggingCfgs := util.GetLoggingConfig()
			util.InitLoggers(loggingCfgs)

			//create os kill sig listener
			mineCh = make(chan string)
			// signal.Notify(c, os.Interrupt)
			exitChannels := make([]*chan os.Signal, 0)

			cfg := config.GetConfig()
			var ds *ops.DataServerOps
			if !cfg.EnablePoolWorker {
				setup.AddDBToCtx(false)
				ch := make(chan os.Signal)
				exitChannels = append(exitChannels, &ch)

				var err error
				ds, err = ops.CreateDataServerOps(setup.CTX, ch)
				if err != nil {
					log.Fatal(err)
				}
				//start and wait for it to be ready
				ds.Start(setup.CTX)
				<-ds.Ready()
			}
			//start miner
			ch := make(chan os.Signal)
			exitChannels = append(exitChannels, &ch)
			miner, err := ops.CreateMiningManager(setup.CTX, ch, ops.NewSubmitter())
			if err != nil {
				log.Fatal(err)
			}
			miner.Start(setup.CTX)
			initializing = false
			mining = true
			//now we wait for kill sig
			<-mineCh
			//and then notify exit channels
			for _, ch := range exitChannels {
				*ch <- os.Interrupt
			}
			cnt := 0
			start := time.Now()
			for {
				cnt++
				dsStopped := false
				minerStopped := false

				if ds != nil {
					dsStopped = !ds.Running
				} else {
					dsStopped = true
				}

				if miner != nil {
					minerStopped = !miner.Running
				} else {
					minerStopped = true
				}

				if !dsStopped && !minerStopped && cnt > 60 {
					fmt.Printf("\U000026A0 Taking longer than expected to stop operations. Waited %v so far\n", time.Now().Sub(start))
				} else if dsStopped && minerStopped {
					break
				}
				time.Sleep(500 * time.Millisecond)
			}
			fmt.Printf("Main shutdown complete \U00002622\n")
			mining = false
		}()
	}

}

func stopMine() bool {

	if mining == true {
		mineCh <- "Stopping miner"
		return true
	}
	return false

}

func isMining() bool {
	return mining
}

func isInit() bool {
	return initializing
}

func showLogs(index int) string {
	var content string
	file, _ := os.Open("./webview/public/miningLogs.log")
	reader := bufio.NewScanner(file)

	reader.Split(bufio.ScanLines)

	ind := 0
	for reader.Scan() {
		ind = ind + 1

		if ind <= index {
			continue
		}

		content = content + reader.Text() + "\n"
	}

	return content
}

func transfer(address string, amount string) (bool, error) {

	n := new(big.Int)
	n, ok := n.SetString(amount, 10)
	if !ok {
		fmt.Println("SetString: error")
		return false, errors.New("SetString: error")
	}
	addr := common.HexToAddress("70997970C51812dc3A010C7d01b50e0d17dc79C8")

	err := ops.Transfer(addr, n, setup.CTX)

	if err != nil {

		fmt.Println(err.Error())
		return false, err

	}
	return true, nil

}
