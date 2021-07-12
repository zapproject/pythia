package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	cli "github.com/jawher/mow.cli"
	ZapCommon "github.com/zapproject/zap-miner/common"
	config "github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/contracts"
	"github.com/zapproject/zap-miner/contracts1"
	"github.com/zapproject/zap-miner/contracts2"
	db "github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/ops"
	"github.com/zapproject/zap-miner/rpc"
	token "github.com/zapproject/zap-miner/token"
	"github.com/zapproject/zap-miner/util"
	"github.com/zapproject/zap-miner/vault"
)

var ctx context.Context

func ErrorHandler(err error, operation string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s failed: %s\n", operation, err.Error())
		cli.Exit(-1)
	}
}

func buildContext() error {
	cfg := config.GetConfig()

	if !cfg.EnablePoolWorker {
		//create an rpc client
		client, err := rpc.NewClient(cfg.NodeURL)
		if err != nil {
			log.Fatal(err)
		}

		//create an instance of the Zap master contract for on-chain interactions
		tokenAddress := common.HexToAddress(cfg.TokenAddress)
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		vaultAddress := common.HexToAddress(cfg.VaultAddress)
		masterInstance, err := contracts.NewZapMaster(contractAddress, client)
		transactorInstance, err := contracts1.NewZapTransactor(contractAddress, client)
		newZapInstance, err := contracts2.NewZap(contractAddress, client)
		newTransactorInstance, err := contracts2.NewZapTransactor(contractAddress, client)
		tokenInstance, err := token.NewZapTokenBSCTransactor(tokenAddress, client)
		// tokenListener, err := token.NewZapTokenBSCFilterer(tokenAddress, client)
		vaultInstance, _ := vault.NewVaultTransactor(vaultAddress, client)
		if err != nil {
			log.Fatal(err)
		}

		ctx = context.WithValue(context.Background(), ZapCommon.ClientContextKey, client)
		ctx = context.WithValue(ctx, ZapCommon.ContractAddress, contractAddress)
		ctx = context.WithValue(ctx, ZapCommon.MasterContractContextKey, masterInstance)
		ctx = context.WithValue(ctx, ZapCommon.TransactorContractContextKey, transactorInstance)
		ctx = context.WithValue(ctx, ZapCommon.TokenTransactorContractContextKey, tokenInstance)
		ctx = context.WithValue(ctx, ZapCommon.NewZapContractContextKey, newZapInstance)
		ctx = context.WithValue(ctx, ZapCommon.NewTransactorContractContextKey, newTransactorInstance)
		// ctx = context.WithValue(ctx, ZapCommon.TokenFilterContextKey, tokenListener)
		ctx = context.WithValue(ctx, ZapCommon.VaultTransactorContractContextKey, vaultInstance)

		// start event listener
		// tokenListener.ParseTransfer()
		// go listenTransfers(client, cfg)

		privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
		if err != nil {
			return fmt.Errorf("problem getting private key: %s", err.Error())
		}
		ctx = context.WithValue(ctx, ZapCommon.PrivateKey, privateKey)

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("error casting public key to ECDSA")
		}

		publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		ctx = context.WithValue(ctx, ZapCommon.PublicAddress, publicAddress)

		//Issue #55, halt if client is still syncing with Ethereum network
		s, err := client.IsSyncing(ctx)
		if err != nil {
			return fmt.Errorf("could not determine if Ethereum client is syncing: %v\n", err)
		}
		if s {
			return fmt.Errorf("ethereum node is still sycning with the network")
		}
	}
	return nil
}

func AddDBToCtx(remote bool) error {
	cfg := config.GetConfig()
	//create a db instance
	os.RemoveAll(cfg.DBFile)
	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		return err
	}

	var dataProxy db.DataServerProxy
	if remote {
		proxy, err := db.OpenRemoteDB(DB)
		if err != nil {
			log.Fatal(err)
		}
		dataProxy = proxy
	} else {
		proxy, err := db.OpenLocalProxy(DB)
		if err != nil {
			log.Fatal(err)
		}
		dataProxy = proxy
	}
	ctx = context.WithValue(ctx, ZapCommon.DataProxyKey, dataProxy)
	ctx = context.WithValue(ctx, ZapCommon.DBContextKey, DB)
	return nil
}

var GitTag string
var GitHash string

const versionMessage = `Zap
    The official Zap Miner %s (%s)
    -----------------------------------------
	Website: https://Zap.org
	Github:  https://github.com/zapproject/zap-miner
`

func App() *cli.Cli {
	app := cli.App("ZapMiner", "The Zap.org official miner")

	//app wide config options
	configPath := app.StringOpt("config", "config.json", "Path to the primary JSON config file")
	logPath := app.StringOpt("logConfig", "loggingConfig.json", "Path to a JSON logging config file")

	//this will get run before any of the commands
	app.Before = func() {
		ErrorHandler(config.ParseConfig(*configPath), "parsing config file")
		ErrorHandler(util.ParseLoggingConfig(*logPath), "parsing log file")
		ErrorHandler(buildContext(), "building context")
	}

	versionMessage := fmt.Sprintf(versionMessage, GitTag, GitHash)
	app.Version("version", versionMessage)

	app.Command("stake", "staking operations", stakeCmd)
	app.Command("transfer", "send ZAP to address", moveCmd(ops.Transfer))
	app.Command("approve", "approve ZAP to address", moveCmd(ops.Approve))
	app.Command("balance", "check balance of address", balanceCmd)
	app.Command("dispute", "dispute operations", disputeCmd)
	app.Command("mine", "mine for ZAP", mineCmd)
	app.Command("dataserver", "start an independent dataserver", dataserverCmd)
	return app
}

func stakeCmd(cmd *cli.Cmd) {
	cmd.Command("deposit", "deposit ZAP stake", simpleCmd(ops.Deposit))
	cmd.Command("withdraw", "withdraw ZAP stake", simpleCmd(ops.WithdrawStake))
	cmd.Command("request", "request to withdraw ZAP stake", simpleCmd(ops.RequestStakingWithdraw))
	cmd.Command("status", "show current staking status", simpleCmd(ops.ShowStatus))
}

func simpleCmd(f func(context.Context) error) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Action = func() {
			ErrorHandler(f(ctx), "")
		}
	}
}

func moveCmd(f func(common.Address, *big.Int, context.Context) error) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		amt := ZAPAmount{}
		addr := ETHAddress{}
		cmd.VarArg("AMOUNT", &amt, "amount to transfer")
		cmd.VarArg("ADDRESS", &addr, "ethereum public address")
		cmd.Action = func() {
			ErrorHandler(f(addr.addr, amt.Int, ctx), "move")
		}
	}
}

func balanceCmd(cmd *cli.Cmd) {
	addr := ETHAddress{}
	cmd.VarArg("ADDRESS", &addr, "ethereum public address")
	cmd.Spec = "[ADDRESS]"
	cmd.Action = func() {
		var zero [20]byte
		if bytes.Compare(addr.addr.Bytes(), zero[:]) == 0 {
			addr.addr = ctx.Value(ZapCommon.PublicAddress).(common.Address)
		}
		ErrorHandler(ops.Balance(ctx, addr.addr), "checking balance")
	}
}

func disputeCmd(cmd *cli.Cmd) {
	cmd.Command("vote", "vote on an active dispute", voteCmd)
	cmd.Command("new", "start a new dispute", newDisputeCmd)
	cmd.Command("show", "show existing disputes", simpleCmd(ops.List))
}

func voteCmd(cmd *cli.Cmd) {
	disputeID := EthereumInt{}
	cmd.VarArg("DISPUTE_ID", &disputeID, "dispute id")
	supports := cmd.BoolArg("SUPPORT", false, "do you support the dispute? (true|false)")
	cmd.Action = func() {
		ErrorHandler(ops.Vote(disputeID.Int, *supports, ctx), "vote")
	}
}

func newDisputeCmd(cmd *cli.Cmd) {
	requestID := EthereumInt{}
	timestamp := EthereumInt{}
	minerIndex := EthereumInt{}
	cmd.VarArg("REQUEST_ID", &requestID, "request id")
	cmd.VarArg("TIMESTAMP", &timestamp, "timestamp")
	cmd.VarArg("MINER_INDEX", &minerIndex, "miner to dispute (0-4)")
	cmd.Action = func() {
		ErrorHandler(ops.Dispute(requestID.Int, timestamp.Int, minerIndex.Int, ctx), "new dispute")
	}
}

func mineCmd(cmd *cli.Cmd) {
	remoteDS := cmd.BoolOpt("remote r", false, "connect to remote dataserver")
	cmd.Action = func() {
		//create os kill sig listener
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		exitChannels := make([]*chan os.Signal, 0)

		cfg := config.GetConfig()
		var ds *ops.DataServerOps
		if !cfg.EnablePoolWorker {
			ErrorHandler(AddDBToCtx(*remoteDS), "initializing database")
			if !*remoteDS {
				ch := make(chan os.Signal)
				exitChannels = append(exitChannels, &ch)

				var err error
				ds, err = ops.CreateDataServerOps(ctx, ch)
				if err != nil {
					log.Fatal(err)
				}
				//start and wait for it to be ready
				ds.Start(ctx)
				<-ds.Ready()
			}
		}
		//start miner
		DB := ctx.Value(ZapCommon.DataProxyKey).(db.DataServerProxy)
		//DB := ctx.Value(ZapCommon.DBContextKey).(db.DB)
		v, err := DB.Get(db.DisputeStatusKey)
		if err != nil {
			fmt.Println("ignoring --- could not get dispute status.  Check if staked")
		}
		status, _ := hexutil.DecodeBig(string(v))
		if status.Cmp(big.NewInt(1)) != 0 {
			log.Fatalf("Miner is not able to mine with status %v. Stopping all mining immediately", status)
		}
		ch := make(chan os.Signal)
		exitChannels = append(exitChannels, &ch)
		miner, err := ops.CreateMiningManager(ctx, ch, ops.NewSubmitter())
		if err != nil {
			log.Fatal(err)
		}
		miner.Start(ctx)

		//now we wait for kill sig
		<-c
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
				fmt.Printf("Taking longer than expected to stop operations. Waited %v so far\n", time.Now().Sub(start))
			} else if dsStopped && minerStopped {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Printf("Main shutdown complete\n")
	}
}

func dataserverCmd(cmd *cli.Cmd) {
	cmd.Action = func() {
		//create os kill sig listener
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)

		var ds *ops.DataServerOps
		ErrorHandler(AddDBToCtx(true), "initializing database")
		ch := make(chan os.Signal)
		var err error
		ds, err = ops.CreateDataServerOps(ctx, ch)
		if err != nil {
			log.Fatal(err)
		}
		//start and wait for it to be ready
		ds.Start(ctx)
		<-ds.Ready()

		//now we wait for kill sig
		<-c
		//and then notify exit channels
		ch <- os.Interrupt

		cnt := 0
		start := time.Now()
		for {
			cnt++
			dsStopped := false

			if ds != nil {
				dsStopped = !ds.Running
			} else {
				dsStopped = true
			}

			if !dsStopped && cnt > 60 {
				fmt.Printf("Taking longer than expected to stop operations. Waited %v so far\n", time.Now().Sub(start))
			} else if dsStopped {
				break
			}
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Printf("Main shutdown complete\n")
	}

}

func listenTransfers(client rpc.ETHClient, cfg *config.Config) {
	tokenAddress := common.HexToAddress(cfg.TokenAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{tokenAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}

func main() {
	//see, programming is easy. Just create an App() and run it!!!!!
	app := App()
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "app.Run failed: %v\n", err)
	}
}