package ops

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	zapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/contracts"
	zap1 "github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/contracts2"
	"github.com/zapproject/pythia/db"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/util"
)

func setup() error {
	err := config.ParseConfig("../config.json")
	if err != nil {
		return fmt.Errorf("Error with parsing config: %v", err)
	}

	path := "../testConfig.json"
	err = util.ParseLoggingConfig(path)
	if err != nil {
		return err
	}

	return nil
}

func TestDataServerOps(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatalf("Error in setting up test: %v", err)
	}
	exitCh := make(chan os.Signal)
	cfg := config.GetConfig()

	if len(cfg.DBFile) == 0 {
		log.Fatal("Missing dbFile config setting")
	}

	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(cfg.ContractAddress)
	masterInstance, err := contracts.NewZapMaster(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem creating zap master instance: %v\n", err)
	}
	proxy, err := db.OpenLocalProxy(DB)
	if err != nil {
		t.Fatal(err)
	}
	var transactorInstance *zap1.ZapTransactor
	// if err != nil {
	// 	t.Fatalf("Problem with initializing the ZapTransactor: %v\n", err)
	// }

	instance, err := contracts2.NewZap(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem with initializing contracts2: %v\n", err)
	}

	ctx := context.WithValue(context.Background(), zapCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
	ctx = context.WithValue(ctx, zapCommon.DataProxyKey, proxy)
	ctx = context.WithValue(ctx, zapCommon.PublicAddress, common.BytesToAddress([]byte(cfg.PublicAddress)))
	ctx = context.WithValue(ctx, zapCommon.TransactorContractContextKey, transactorInstance)
	ctx = context.WithValue(ctx, zapCommon.NewZapContractContextKey, instance)
	ops, err := CreateDataServerOps(ctx, exitCh)
	if err != nil {
		t.Fatal(err)
	}
	ops.Start(ctx)
	time.Sleep(2 * time.Second)
	exitCh <- os.Interrupt
	time.Sleep(27 * time.Second)
	if ops.Running {
		t.Fatal("data server is still running after stopping")
	}
}
