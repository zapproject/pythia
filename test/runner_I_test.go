package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/contracts"
	"github.com/zapproject/zap-miner/contracts1"
	"github.com/zapproject/zap-miner/contracts2"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/ops"
	"github.com/zapproject/zap-miner/rpc"
	"github.com/zapproject/zap-miner/tracker"
	"github.com/zapproject/zap-miner/util"
)

func Test_I_Runner(t *testing.T) {
	exitCh := make(chan int)
	osExitCh := make(chan os.Signal)
	err := config.ParseConfig("../config.json")
	err = util.ParseLoggingConfig("../loggingConfig.json")
	cfg := config.GetConfig()

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_leveldb"))
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		t.Fatalf("Unable to create rpc client: %v", err)
	}

	submitter := ops.NewSubmitter()

	proxy, err := db.OpenRemoteDB(DB)
	if err != nil {
		t.Fatal(err)
	}

	contractAddress := common.HexToAddress(cfg.ContractAddress)

	transactor1Instance, err := contracts1.NewZapTransactor(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem with initializing the ZapTransactor: %v\n", err)
	}

	transactor2Instance, err := contracts2.NewZapTransactor(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem with initializing the ZapTransactor: %v\n", err)
	}

	ctx := context.WithValue(context.Background(), zapCommon.DataProxyKey, proxy)
	ctx = context.WithValue(ctx, zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.TransactorContractContextKey, transactor1Instance)
	ctx = context.WithValue(ctx, zapCommon.NewTransactorContractContextKey, transactor2Instance)

	masterInstance := ctx.Value(zapCommon.MasterContractContextKey)
	if masterInstance == nil {
		// TODO create error state flag for mock client
		masterInstance, err = contracts.NewZapMaster(contractAddress, client)
		if err != nil {
			t.Errorf("Problem creating zap master instance: %v\n", err)
		}
		ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
	}

	proxy = ctx.Value(zapCommon.DataProxyKey).(db.DataServerProxy)
	requester := ops.CreateDataRequester(osExitCh, submitter, 2, proxy)
	requester.Start(ctx)

	runner, _ := tracker.NewRunner(client, DB)

	runner.Ready()
	runner.Start(ctx, exitCh)
	fmt.Println("runner done")
	time.Sleep(2 * time.Second)
	close(exitCh)
	time.Sleep(1 * time.Second)
}
