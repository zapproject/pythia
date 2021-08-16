package main

import (
	// "bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	// "math/big"
	"os"
	// "os/signal"
	// "time"

	// "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	cli "github.com/jawher/mow.cli"
	ZapCommon "github.com/zapproject/pythia/common"
	config "github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/contracts2"
	db "github.com/zapproject/pythia/db"

	// "github.com/zapproject/pythia/ops"
	"github.com/zapproject/pythia/rpc"
	token "github.com/zapproject/pythia/token"

	// "github.com/zapproject/pythia/util"
	"github.com/zapproject/pythia/vault"
	"github.com/zapproject/pythia/webview"
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
		masterInstance, _ := contracts.NewZapMaster(contractAddress, client)
		transactorInstance, _ := contracts1.NewZapTransactor(contractAddress, client)
		newZapInstance, _ := contracts2.NewZap(contractAddress, client)
		newTransactorInstance, _ := contracts2.NewZapTransactor(contractAddress, client)
		tokenInstance, _ := token.NewZapTokenBSCTransactor(tokenAddress, client)
		vaultInstance, _ := vault.NewVaultTransactor(vaultAddress, client)

		ctx = context.WithValue(context.Background(), ZapCommon.ClientContextKey, client)
		ctx = context.WithValue(ctx, ZapCommon.ContractAddress, contractAddress)
		ctx = context.WithValue(ctx, ZapCommon.MasterContractContextKey, masterInstance)
		ctx = context.WithValue(ctx, ZapCommon.TransactorContractContextKey, transactorInstance)
		ctx = context.WithValue(ctx, ZapCommon.TokenTransactorContractContextKey, tokenInstance)
		ctx = context.WithValue(ctx, ZapCommon.NewZapContractContextKey, newZapInstance)
		ctx = context.WithValue(ctx, ZapCommon.NewTransactorContractContextKey, newTransactorInstance)
		ctx = context.WithValue(ctx, ZapCommon.VaultTransactorContractContextKey, vaultInstance)

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

// var GitTag string
// var GitHash string

// const versionMessage = `
// 	The official Pythia %s (%s)
// 	-----------------------------------------
// 	Website: https://Zap.org
// 	Github:  https://github.com/zapproject/pythia
// 	`

// func App() *cli.Cli {
// 	app := cli.App("Pythia", "The Zap.org official miner")

// 	//app wide config options
// 	configPath := app.StringOpt("config", "config.json", "Path to the primary JSON config file")
// 	logPath := app.StringOpt("logConfig", "loggingConfig.json", "Path to a JSON logging config file")

// 	//this will get run before any of the commands
// 	app.Before = func() {
// 		ErrorHandler(config.ParseConfig(*configPath), "parsing config file")
// 		ErrorHandler(util.ParseLoggingConfig(*logPath), "parsing log file")
// 		ErrorHandler(buildContext(), "building context")
// 	}

// 	versionMessage := fmt.Sprintf(versionMessage, GitTag, GitHash)
// 	app.Version("version", versionMessage)

// 	app.Command("stake", "\U0001F510 staking operations", stakeCmd)
// 	app.Command("transfer", "\U0001F381 send ZAP to address", moveCmd(ops.Transfer))
// 	app.Command("approve", "\U00002705 approve ZAP to address", moveCmd(ops.Approve))
// 	app.Command("balance", "\U0001F440 check balance of address", balanceCmd)
// 	app.Command("dispute", "\U00002696 dispute operations", disputeCmd)
// 	app.Command("mine", "\U000026CF  mine for ZAP", mineCmd)
// 	app.Command("dataserver", "\U0001F5C4  start an independent dataserver", dataserverCmd)
// 	app.Command("gui", "\U0001F5C4  start an independent dataserver", guiCmd)

// 	return app

// }

// func stakeCmd(cmd *cli.Cmd) {
// 	cmd.Command("deposit", "\U0001F512 deposit ZAP stake", simpleCmd(ops.Deposit))
// 	cmd.Command("withdraw", "\U0001F511 withdraw ZAP stake", simpleCmd(ops.WithdrawStake))
// 	cmd.Command("request", "\U000023F2 request to withdraw ZAP stake", simpleCmd(ops.RequestStakingWithdraw))
// 	cmd.Command("status", "\U0001F52E show current staking status", simpleCmd(ops.ShowStatus))
// }

// func simpleCmd(f func(context.Context) error) func(*cli.Cmd) {
// 	return func(cmd *cli.Cmd) {
// 		cmd.Action = func() {
// 			ErrorHandler(f(ctx), "")
// 		}
// 	}
// }

// func moveCmd(f func(common.Address, *big.Int, context.Context) error) func(*cli.Cmd) {
// 	return func(cmd *cli.Cmd) {
// 		amt := ZAPAmount{}
// 		addr := ETHAddress{}
// 		cmd.VarArg("AMOUNT", &amt, "amount to transfer")
// 		cmd.VarArg("ADDRESS", &addr, "ethereum public address")
// 		cmd.Action = func() {
// 			ErrorHandler(f(addr.addr, amt.Int, ctx), "move")
// 		}
// 	}
// }

// func balanceCmd(cmd *cli.Cmd) {
// 	addr := ETHAddress{}
// 	cmd.VarArg("ADDRESS", &addr, "binance public address")
// 	cmd.Spec = "[ADDRESS]"
// 	cmd.Action = func() {
// 		var zero [20]byte
// 		if bytes.Compare(addr.addr.Bytes(), zero[:]) == 0 {
// 			addr.addr = ctx.Value(ZapCommon.PublicAddress).(common.Address)
// 		}
// 		ErrorHandler(ops.Balance(ctx, addr.addr), "checking balance")
// 	}
// }

// func disputeCmd(cmd *cli.Cmd) {
// 	cmd.Command("vote", "\U00002696 vote on an active dispute", voteCmd)
// 	cmd.Command("new", "\U0001F4C4 start a new dispute", newDisputeCmd)
// 	cmd.Command("show", "\U0001F4CA show existing disputes", simpleCmd(ops.List))
// }

// func voteCmd(cmd *cli.Cmd) {
// 	disputeID := EthereumInt{}
// 	cmd.VarArg("DISPUTE_ID", &disputeID, "dispute id")
// 	supports := cmd.BoolArg("SUPPORT", false, "do you support the dispute? (true|false)")
// 	cmd.Action = func() {
// 		ErrorHandler(ops.Vote(disputeID.Int, *supports, ctx), "vote")
// 	}
// }

// func newDisputeCmd(cmd *cli.Cmd) {
// 	requestID := EthereumInt{}
// 	timestamp := EthereumInt{}
// 	minerIndex := EthereumInt{}
// 	cmd.VarArg("REQUEST_ID", &requestID, "request id")
// 	cmd.VarArg("TIMESTAMP", &timestamp, "timestamp")
// 	cmd.VarArg("MINER_INDEX", &minerIndex, "miner to dispute (0-4)")
// 	cmd.Action = func() {
// 		ErrorHandler(ops.Dispute(requestID.Int, timestamp.Int, minerIndex.Int, ctx), "new dispute")
// 	}
// }

// func mineCmd(cmd *cli.Cmd) {
// 	remoteDS := cmd.BoolOpt("remote r", false, "connect to remote dataserver")
// 	cmd.Action = func() {
// 		//create os kill sig listener
// 		c := make(chan os.Signal)
// 		signal.Notify(c, os.Interrupt)
// 		exitChannels := make([]*chan os.Signal, 0)

// 		cfg := config.GetConfig()
// 		var ds *ops.DataServerOps
// 		if !cfg.EnablePoolWorker {
// 			ErrorHandler(AddDBToCtx(*remoteDS), "\U0001F5C4 initializing database \U0001F5C4")
// 			if !*remoteDS {
// 				ch := make(chan os.Signal)
// 				exitChannels = append(exitChannels, &ch)

// 				var err error
// 				ds, err = ops.CreateDataServerOps(ctx, ch)
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				//start and wait for it to be ready
// 				ds.Start(ctx)
// 				<-ds.Ready()
// 			}
// 		}
// 		//start miner
// 		ch := make(chan os.Signal)
// 		exitChannels = append(exitChannels, &ch)
// 		miner, err := ops.CreateMiningManager(ctx, ch, ops.NewSubmitter())
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		miner.Start(ctx)

// 		//now we wait for kill sig
// 		<-c
// 		//and then notify exit channels
// 		for _, ch := range exitChannels {
// 			*ch <- os.Interrupt
// 		}
// 		cnt := 0
// 		start := time.Now()
// 		for {
// 			cnt++
// 			dsStopped := false
// 			minerStopped := false

// 			if ds != nil {
// 				dsStopped = !ds.Running
// 			} else {
// 				dsStopped = true
// 			}

// 			if miner != nil {
// 				minerStopped = !miner.Running
// 			} else {
// 				minerStopped = true
// 			}

// 			if !dsStopped && !minerStopped && cnt > 60 {
// 				fmt.Printf("\U000026A0 Taking longer than expected to stop operations. Waited %v so far\n", time.Now().Sub(start))
// 			} else if dsStopped && minerStopped {
// 				break
// 			}
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 		fmt.Printf("Main shutdown complete \U00002622\n")
// 	}
// }

// func dataserverCmd(cmd *cli.Cmd) {
// 	cmd.Action = func() {
// 		//create os kill sig listener
// 		c := make(chan os.Signal)
// 		signal.Notify(c, os.Interrupt)

// 		var ds *ops.DataServerOps
// 		ErrorHandler(AddDBToCtx(true), "\U0001F5C4 initializing database \U0001F5C4")
// 		ch := make(chan os.Signal)
// 		var err error
// 		ds, err = ops.CreateDataServerOps(ctx, ch)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		//start and wait for it to be ready
// 		ds.Start(ctx)
// 		<-ds.Ready()

// 		//now we wait for kill sig
// 		<-c
// 		//and then notify exit channels
// 		ch <- os.Interrupt

// 		cnt := 0
// 		start := time.Now()
// 		for {
// 			cnt++
// 			dsStopped := false

// 			if ds != nil {
// 				dsStopped = !ds.Running
// 			} else {
// 				dsStopped = true
// 			}

// 			if !dsStopped && cnt > 60 {
// 				fmt.Printf("\U000026A0 Taking longer than expected to stop operations. Waited %v so far\n", time.Now().Sub(start))
// 			} else if dsStopped {
// 				break
// 			}
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 		fmt.Printf("Main shutdown complete \U00002622\n")
// 	}

// }

// func guiCmd(cmd *cli.Cmd) {
// 	webview.LoadWebview()
// }

// func listenTransfers(client rpc.ETHClient, cfg *config.Config) {
// 	tokenAddress := common.HexToAddress(cfg.TokenAddress)
// 	query := ethereum.FilterQuery{
// 		Addresses: []common.Address{tokenAddress},
// 	}

// 	logs := make(chan types.Log)
// 	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for {
// 		select {
// 		case err := <-sub.Err():
// 			log.Fatal(err)
// 		case vLog := <-logs:
// 			fmt.Println(vLog)
// 		}
// 	}
// }

func main() {
	//see, programming is easy. Just create an App() and run it!!!!!
	// app := App()
	// err := app.Run(os.Args)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "\U0001F6AB app.Run failed: %v\n", err)
	// }

	webview.Start()

}
