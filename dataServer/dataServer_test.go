package dataServer

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	zapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/contracts2"
	"github.com/zapproject/pythia/db"
	"github.com/zapproject/pythia/rpc"
	"github.com/zapproject/pythia/util"
)

func setup() {
	err := config.ParseConfig("../config.json")
	if err != nil {
		fmt.Errorf("Can't parse config for test.")
	}
	path := "../testConfig.json"
	err = util.ParseLoggingConfig(path)
	if err != nil {
		fmt.Errorf("Can't parse logging config for test.")
	}
}

func TestDataServer(t *testing.T) {
	exitCh := make(chan int)
	setup()
	cfg := config.GetConfig()
	if len(cfg.DBFile) == 0 {
		log.Fatal("Missing dbFile config setting")
	}

	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		log.Fatal(err)
	}
	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(cfg.ContractAddress)
	masterInstance, err := contracts.NewZapMaster(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem creating zap master instance: %v\n", err)
	}

	proxy, err := db.OpenRemoteDB(DB)
	if err != nil {
		t.Fatalf("Problem creating remote db: %v\n", err)
	}

	instance, err := contracts2.NewZap(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem with initializing contracts2: %v\n", err)
	}

	ctx := context.WithValue(context.Background(), zapCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
	ctx = context.WithValue(ctx, zapCommon.DataProxyKey, proxy)
	ctx = context.WithValue(ctx, zapCommon.NewZapContractContextKey, instance)
	ds, err := CreateServer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	ds.Start(ctx, exitCh)

	time.Sleep(5000 * time.Millisecond)

	isReady := <-ds.Ready()

	if !isReady {
		t.Fatal("Data Server is not ready")
	}

	exitCh <- 1
	time.Sleep(1 * time.Second)
	if !ds.Stopped {
		t.Fatal("Did not stop server")
	}
}
