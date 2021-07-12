package routes

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/contracts"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
	"github.com/zapproject/zap-miner/tracker"
	"github.com/zapproject/zap-miner/util"
)

var DB db.DB
var err error
var ctx context.Context
var opts *rpc.MockOptions
var cfg *config.Config
var currentChallenge [32]byte
var requestID *big.Int
var difficulty *big.Int
var queryString string
var granularity *big.Int
var totalTip *big.Int
var myStatus bool

func prep(t *testing.T) {
	cfg = config.GetConfig()
	DB, err = db.Open(cfg.DBFile)
	if err != nil {
		t.Fatal(err)
	}

	startBal := big.NewInt(356000)

	hash := math.PaddedBigBytes(big.NewInt(256), 32)
	var b32 [32]byte
	for i, v := range hash {
		b32[i] = v
	}
	queryStr := "json(https://coinbase.com)"
	chal := &rpc.CurrentChallenge{
		ChallengeHash: b32,
		RequestID:     big.NewInt(1),
		Difficulty:    big.NewInt(500),
		QueryString:   queryStr,
		Granularity:   big.NewInt(1000),
		Tip:           big.NewInt(0),
	}
	opts = &rpc.MockOptions{
		ETHBalance:       startBal,
		Nonce:            1,
		GasPrice:         big.NewInt(700000000),
		TokenBalance:     big.NewInt(0),
		MiningStatus:     true,
		Top50Requests:    []*big.Int{},
		CurrentChallenge: chal,
		DisputeStatus:    big.NewInt(1),
	}

	client := rpc.NewMockClientWithValues(opts)

	ctx = context.WithValue(context.Background(), zapCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, zapCommon.ClientContextKey, client)

	masterInstance := ctx.Value(zapCommon.MasterContractContextKey)
	if masterInstance == nil {
		contractAddress := common.HexToAddress(cfg.ContractAddress)
		// TODO create error state flag for mock client
		masterInstance, err = contracts.NewZapMaster(contractAddress, client)
		if err != nil {
			t.Errorf("Problem creating zap master instance: %v\n", err)
		}
		ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
	}

	_fromAddress := cfg.PublicAddress

	//convert to address
	fromAddress := common.HexToAddress(_fromAddress)

	instance := ctx.Value(zapCommon.MasterContractContextKey).(*contracts.ZapMaster)
	currentChallenge, requestID, difficulty, queryString, granularity, totalTip, err = instance.GetCurrentVariables(nil)
	if err != nil {
		fmt.Println("Current Variables Retrieval Error")
	}

	myStatus, err = instance.DidMine(nil, currentChallenge, fromAddress)
	if err != nil {
		t.Fatalf("Could not retrieve mining status: %v", err)
	}

}

func setup(t *testing.T) {
	err := config.ParseConfig("../../config.json")
	if err != nil {
		fmt.Errorf("Can't parse config for test.")
	}
	path := "../../testConfig.json"
	err = util.ParseLoggingConfig(path)
	if err != nil {
		fmt.Errorf("Can't parse logging config for test.")
	}

	prep(t)
}

func TestCurrentChanllengeHandler(t *testing.T) {
	ch := &CurrentChallengeHandler{}
	if cfg == nil {
		setup(t)
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

func TestDifficultyHandler(t *testing.T) {
	dh := &DifficultyHandler{}
	if cfg == nil {
		setup(t)
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

func TestDisputeStatusHandler(t *testing.T) {
	dsh := &DisputeStatusHandler{}
	if cfg == nil {
		setup(t)
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

func TestGasHandler(t *testing.T) {
	gh := &GasHandler{}
	if cfg == nil {
		setup(t)
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

func TestGranularityHandler(t *testing.T) {
	gh := &GranularityHandler{}
	if cfg == nil {
		setup(t)
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

func TestMiningStatusHandler(t *testing.T) {
	msh := &MiningStatusHandler{}
	if cfg == nil {
		setup(t)
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

func TestQueryStringHandler(t *testing.T) {
	qsh := &QueryStringHandler{}
	if cfg == nil {
		setup(t)
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

func TestRequestIdHandler(t *testing.T) {
	ridh := &RequestIdHandler{}
	if cfg == nil {
		setup(t)
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

func TestTotalTipHandler(t *testing.T) {
	tth := &TotalTipHandler{}
	if cfg == nil {
		setup(t)
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
}

func TestZapBalanceHandler(t *testing.T) {
	tbh := &TokenBalanceHandler{}
	if cfg == nil {
		setup(t)
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
