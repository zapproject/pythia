package  webview
import (
	// "time"
	"fmt"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/setup"
	// "github.com/zapproject/pythia/ops"
	ZapCommon "github.com/zapproject/pythia/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/util"
	"github.com/zapproject/pythia/ops"

	"log"
	"os"
	// "os/signal"
	// "syscall"
	"time"
)


var mineCh chan string

func configWallet(cfg string) (bool, error){
	// fmt.Println(req.TokenAddress )
	err := config.ParseConfigBytes([]byte(cfg))
	if err != nil{
		fmt.Println(err.Error())
	} 

	return setup.App()
}


func getBalance() (string, error) {
	// setup.App()
	addr := setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)
	instance := setup.CTX.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	zapBalance, err := instance.BalanceOf(nil, addr)

	if err != nil{
		fmt.Print(err.Error())
		return "",err
	}

	return zapBalance.String(),nil

}

func stakeStatus() (int64, error) {
	  
	tmaster := setup.CTX.Value(ZapCommon.MasterContractContextKey).(*contracts.ZapMaster)

	publicAddress := setup.CTX.Value(ZapCommon.PublicAddress).(common.Address)
	status, _, err := tmaster.GetStakerInfo(nil, publicAddress)
	if err != nil {
		return 0,fmt.Errorf("failed to get stake status: %s", err.Error())
	}

	fmt.Println(status)

	return status.Int64(),nil

}

func startMine() {

	go func(){
		file, _ := os.OpenFile("./webview/public/miningLogs.html", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

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
	}()
	

}

func stopMine() {
	
	mineCh <- "Stahp it"

}