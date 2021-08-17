package ops

import (
	"context"
	"log"
	"math/big"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/contracts2"
	"github.com/zapproject/pythia/db"
	"github.com/zapproject/pythia/rest"
	"github.com/zapproject/pythia/rpc"
)

var requestID *big.Int

func TestRequestDataOps(t *testing.T) {
	err := setup()
	if err != nil {
		t.Fatalf("Error in setting up test: %v", err)
	}
	exitCh := make(chan os.Signal)
	cfg := config.GetConfig()

	submitter := NewSubmitter()
	DB, err := db.Open(cfg.DBFile)
	if err != nil {
		t.Fatalf("Error with opening DB: %v", err)
	}
	defer DB.Close()

	//delete any request id
	DB.Delete(db.RequestIdKey)

	if err != nil {
		log.Fatal(err)
	}

	proxy, err := db.OpenRemoteDB(DB)
	if err != nil {
		log.Fatal(err)
	}

	client, err := rpc.NewClient(cfg.NodeURL)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(cfg.ContractAddress)
	transactor1Instance, _ := contracts1.NewZapTransactor(contractAddress, client)
	transactor2Instance, _ := contracts2.NewZapTransactor(contractAddress, client)
	masterInstance, _ := contracts.NewZapMaster(contractAddress, client)

	ctx := context.WithValue(context.Background(), zapCommon.DBContextKey, DB)
	ctx = context.WithValue(ctx, zapCommon.DataProxyKey, proxy)
	proxy = ctx.Value(zapCommon.DataProxyKey).(db.DataServerProxy)
	ctx = context.WithValue(ctx, zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.TransactorContractContextKey, transactor1Instance)
	ctx = context.WithValue(ctx, zapCommon.NewTransactorContractContextKey, transactor2Instance)
	ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
	reqData := CreateDataRequester(exitCh, submitter, 2, proxy)

	server, err := rest.Create(ctx, cfg.ServerHost, cfg.ServerPort)
	if err != nil {
		t.Fatal(err)

	}
	server.Start()
	defer server.Stop()

	//it should not request data if not configured to do it
	cfg.RequestData = 0
	reqData.Start(ctx)
	time.Sleep(300 * time.Millisecond)
	assert.Equal(t, false, reqData.submittingRequests, "they should be equal")

	cfg.RequestData = 1
	DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(big.NewInt(0))))
	reqData.Start(ctx)
	time.Sleep(2500 * time.Millisecond)
	reqIdBytes, err := DB.Get(db.RequestIdKey)
	assert.Nil(t, err, "Error in getting Request ID")

	reqIdInt, _ := strconv.Atoi(string(reqIdBytes))
	requestID = big.NewInt(int64(reqIdInt))
	log.Printf("Request ID: %d", requestID)
	assert.NotNil(t, requestID, "Should have requested data")

	requestID = nil
	DB.Put(db.RequestIdKey, []byte(hexutil.EncodeBig(big.NewInt(1))))
	time.Sleep(2500 * time.Millisecond)
	assert.Nil(t, requestID, "Should not have requested data when a challenge request is in progress")

	exitCh <- os.Kill
	time.Sleep(300 * time.Millisecond)
	assert.False(t, reqData.submittingRequests, "Should not be submitting requests after exit sig")
}
