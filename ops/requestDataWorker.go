package ops

import (
	"context"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	zapCommon "github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/db"
	"github.com/zapproject/pythia/util"
)

//DataRequester responsible for submitting tips to request data periodically if configured to do so
type DataRequester struct {
	submittingRequests bool
	exitCh             chan os.Signal
	log                *util.Logger
	checkInterval      time.Duration
	proxy              db.DataServerProxy
	submitter          zapCommon.TransactionSubmitter
}

const (
	statusWaitNext = iota + 1
	statusFailure
	statusSuccess
)

//CreateDataRequester creates a requester instance
func CreateDataRequester(exitCh chan os.Signal, submitter zapCommon.TransactionSubmitter, checkIntervalSeconds time.Duration, proxy db.DataServerProxy) *DataRequester {
	if checkIntervalSeconds == 0 {
		checkIntervalSeconds = 300
	}
	return &DataRequester{exitCh: exitCh, submitter: submitter, proxy: proxy, checkInterval: checkIntervalSeconds, log: util.NewLogger("ops", "DataRequester")}
}

//Start kicks of go routines to periodically submit tips if configured to do so
func (r *DataRequester) Start(ctx context.Context) error {
	cfg := config.GetConfig()

	//if we're not configured to request anything.
	if cfg.RequestData == 0 {
		return nil
	}

	r.submittingRequests = true
	ticker := time.NewTicker(r.checkInterval)
	go func() {
		for {
			select {
			case _ = <-r.exitCh:
				{
					r.log.Info("Stopping data requester")
					r.submittingRequests = false
					ticker.Stop()
					return
				}
			case _ = <-ticker.C:
				{
					r.maybeRequestData(ctx)
				}
			}
		}
	}()
	return nil
}

//IsRunning checks whether this requester is requesting data
func (r *DataRequester) IsRunning() bool {
	return r.submittingRequests
}

func (r *DataRequester) reqDataCallback(ctx context.Context, contract zapCommon.ContractInterface) (*types.Transaction, error) {
	cfg := config.GetConfig()

	//if we're not configured to request anything.
	if cfg.RequestData == 0 {
		r.log.Warn("Getting asked to request data when not configured to do so")
		return nil, nil
	}

	masterInstance := ctx.Value(zapCommon.MasterContractContextKey).(*contracts.ZapMaster)

	keys := []string{
		db.RequestIdKey,
	}

	m, err := r.proxy.BatchGet(keys)
	if err != nil {
		r.log.Error("Could not get data from data proxy, cannot continue at all")
		return nil, nil
	}
	r.log.Debug("Received data: %v", m)

	reqID, stat := r.getInt(m[db.RequestIdKey])
	if stat == statusWaitNext || stat == statusFailure {
		return nil, nil
	}
	zapBalance, err := masterInstance.BalanceOf(nil, common.HexToAddress(cfg.PublicAddress))
	if err != nil {
		r.log.Error("Error in getting balance to add tip from: ", err)
	}

	b, _ := new(big.Int).SetString("1000", 10)
	c := big.NewInt(0).Sub(zapBalance, b)

	if c.Cmp(big.NewInt(cfg.RequestTips*int64(1e18))) < 0 {
		r.log.Info("Not enough tokens to requestData with this tip: ", cfg.RequestTips, c)
		return nil, nil
	}
	if reqID.Cmp(big.NewInt(0)) != 0 {
		r.log.Info("There is a challenge being mined right now so will not request data")
		return nil, nil
	}
	r.log.Info("Submitting tip for requestID: %v\n", cfg.RequestData)
	return contract.AddTip(big.NewInt(int64(cfg.RequestData)), big.NewInt(cfg.RequestTips*int64(1e18)))
}

func (r *DataRequester) maybeRequestData(ctx context.Context) {
	r.log.Info("Checking whether to submit data request...")
	err := r.submitter.PrepareTransaction(ctx, r.proxy, "AddTip", r.reqDataCallback)
	if err != nil {
		r.log.Error("Problem preparing contract transaction: %v\n", err)
	}
	r.log.Info("Finished checking whether to submit data")
}

func (r *DataRequester) getInt(data []byte) (*big.Int, int) {
	if data == nil || len(data) == 0 {
		return nil, statusWaitNext
	}

	val, err := hexutil.DecodeBig(string(data))
	if err != nil {
		r.log.Error("Problem decoding int: %v", err)
		return nil, statusFailure
	}
	return val, statusSuccess
}
