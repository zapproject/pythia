package routes

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/contracts"
	"github.com/zapproject/zap-miner/contracts1"
	"github.com/zapproject/zap-miner/contracts2"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/ops"
	"github.com/zapproject/zap-miner/rest"
	"github.com/zapproject/zap-miner/rpc"
	"github.com/zapproject/zap-miner/tracker"
	"github.com/zapproject/zap-miner/util"
)

// var DB db.DB
// var err error
// var ctx context.Context
// var cfg *config.Config
// var currentChallenge [32]byte
// var requestID *big.Int
// var difficulty *big.Int
// var queryString string
// var granularity *big.Int
// var totalTip *big.Int

// var instance *contracts.ZapMaster
// var fromAddress common.Address
var client rpc.ETHClient
var status *big.Int

func prep_I(t *testing.T) {
	exitCh := make(chan os.Signal)
	submitter := ops.NewSubmitter()

	cfg = config.GetConfig()
	DB, err = db.Open(cfg.DBFile)

	//delete any request id
	DB.Delete(db.RequestIdKey)

	if err != nil {
		t.Fatal(err)
	}

	proxy, err := db.OpenRemoteDB(DB)
	if err != nil {
		t.Fatal(err)
	}

	client, err = rpc.NewClient(cfg.NodeURL)
	if err != nil {
		t.Fatal(err)
	}

	contractAddress := common.HexToAddress(cfg.ContractAddress)

	ctx = context.WithValue(context.Background(), zapCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, zapCommon.ClientContextKey, client)

	masterInstance := ctx.Value(zapCommon.MasterContractContextKey)
	if masterInstance == nil {
		// TODO create error state flag for mock client
		masterInstance, err = contracts.NewZapMaster(contractAddress, client)
		if err != nil {
			t.Errorf("Problem creating zap master instance: %v\n", err)
		}
		ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
	}

	transactor1Instance, err := contracts1.NewZapTransactor(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem with initializing the ZapTransactor: %v\n", err)
	}

	transactor2Instance, err := contracts2.NewZapTransactor(contractAddress, client)
	if err != nil {
		t.Fatalf("Problem with initializing the ZapTransactor: %v\n", err)
	}

	ctx = context.WithValue(ctx, zapCommon.DataProxyKey, proxy)
	proxy = ctx.Value(zapCommon.DataProxyKey).(db.DataServerProxy)
	ctx = context.WithValue(ctx, zapCommon.TransactorContractContextKey, transactor1Instance)
	ctx = context.WithValue(ctx, zapCommon.NewTransactorContractContextKey, transactor2Instance)
	reqData := ops.CreateDataRequester(exitCh, submitter, 2, proxy)

	server, err := rest.Create(ctx, cfg.ServerHost, cfg.ServerPort)
	if err != nil {
		t.Fatalf("Unable to create remote rest server: %v", err)
	}

	server.Start()

	cfg.RequestData = 1
	DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(big.NewInt(0))))
	reqData.Start(ctx)
	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	instance := ctx.Value(zapCommon.MasterContractContextKey).(*contracts.ZapMaster)

	currentChallenge, requestID, difficulty, queryString, granularity, totalTip, err = instance.GetCurrentVariables(nil)
	if err != nil {
		fmt.Println("Current Variables Retrieval Error")
	}

	status, _, err = instance.GetStakerInfo(nil, fromAddress)
	if err != nil {
		t.Fatalf("Error in getting dispute status: %v", err)
	}
}

func setup_I(t *testing.T) {
	err := config.ParseConfig("../../config.json")
	if err != nil {
		fmt.Errorf("Can't parse config for test.")
	}
	path := "../../testConfig.json"
	err = util.ParseLoggingConfig(path)
	if err != nil {
		fmt.Errorf("Can't parse logging config for test.")
	}

	prep_I(t)
}

func Test_I_CurrentChanllengeHandler(t *testing.T) {
	ch := &CurrentChallengeHandler{}
	if cfg == nil {
		setup_I(t)
	}

	DB.Put(db.CurrentChallengeKey, currentChallenge[:])

	code, payload := ch.Incoming(ctx, nil)
	t.Logf("Current challenge payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}
}

func Test_I_DifficultyHandler(t *testing.T) {
	dh := &DifficultyHandler{}
	if cfg == nil {
		setup_I(t)
	}

	enc := hexutil.EncodeBig(difficulty)
	DB.Put(db.DifficultyKey, []byte(enc))

	code, payload := dh.Incoming(ctx, nil)
	t.Logf("Difficulty payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}
}

func Test_I_DisputeStatusHandler(t *testing.T) {
	dsh := &DisputeStatusHandler{}
	if cfg == nil {
		setup_I(t)
	}

	trkr := &tracker.DisputeTracker{}
	err = trkr.Exec(ctx)

	if err != nil {
		t.Fatal(err)
	}

	code, payload := dsh.Incoming(ctx, nil)
	t.Logf("Dispute status payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}
}

func Test_I_GasHandler(t *testing.T) {
	gh := &GasHandler{}
	if cfg == nil {
		setup_I(t)
	}

	trkr := &tracker.GasTracker{}

	err = trkr.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}

	code, payload := gh.Incoming(ctx, nil)
	t.Logf("Gas Price payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}
}

func Test_I_GranularityHandler(t *testing.T) {
	gh := &GranularityHandler{}
	if cfg == nil {
		setup_I(t)
	}

	enc := hexutil.EncodeBig(granularity)
	DB.Put(db.GranularityKey, []byte(enc))

	code, payload := gh.Incoming(ctx, nil)
	t.Logf("Granularity payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}
}

func Test_I_MiningStatusHandler(t *testing.T) {
	msh := &MiningStatusHandler{}
	if cfg == nil {
		setup_I(t)
	}

	ms := []byte{0}
	if myStatus {
		ms = []byte{1}
	}
	DB.Put(db.MiningStatusKey, ms)

	code, payload := msh.Incoming(ctx, nil)
	t.Logf("Mining status payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}
}

func Test_I_QueryStringHandler(t *testing.T) {
	qsh := &QueryStringHandler{}
	if cfg == nil {
		setup_I(t)
	}

	DB.Put(db.QueryStringKey, []byte(queryString))

	code, payload := qsh.Incoming(ctx, nil)
	t.Logf("Query string payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}
}

func Test_I_RequestIdHandler(t *testing.T) {
	ridh := &RequestIdHandler{}
	if cfg == nil {
		setup_I(t)
	}

	enc := hexutil.EncodeBig(requestID)
	DB.Put(db.RequestIdKey, []byte(enc))

	code, payload := ridh.Incoming(ctx, nil)
	t.Logf("Request Id payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}
}

func Test_I_TotalTipHandler(t *testing.T) {
	tth := &TotalTipHandler{}
	if cfg == nil {
		setup_I(t)
	}

	enc := hexutil.EncodeBig(totalTip)
	DB.Put(db.TotalTipKey, []byte(enc))

	code, payload := tth.Incoming(ctx, nil)
	t.Logf("Total tip payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}

	_, cancel := context.WithCancel(ctx)

	cancel()
}

func Test_I_ZapBalanceHandler(t *testing.T) {
	tbh := &TokenBalanceHandler{}
	if cfg == nil {
		setup_I(t)
	}

	trkr := &tracker.GasTracker{}

	err = trkr.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}

	code, payload := tbh.Incoming(ctx, nil)
	t.Logf("Token balance payload: %s\n", payload)
	if code != 200 {
		if !strings.Contains(payload, "error") {
			t.Fatal("Expected non-200 code to contain error message")
		}
	}
}
