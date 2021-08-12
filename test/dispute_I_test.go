package test

import (
	"math/big"
	"strconv"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/pythia/common"
	zap "github.com/zapproject/pythia/contracts"
	zap1 "github.com/zapproject/pythia/contracts1"
	"github.com/zapproject/pythia/ops"
	"github.com/zapproject/pythia/rpc"
	token "github.com/zapproject/pythia/token"
	"github.com/zapproject/pythia/vault"
)

func TestDispute(t *testing.T) {
	setupMiners()
	setup()
	setupOwner()

	// [show] check if there are existing disputes - should be 0
	logs := Show()
	assert.Equal(t, 0, logs, "There should be 0 disputes")

	// start miners
	timestamp := StartMiners(t)
	assert.Greater(t, timestamp.Int64(), big.NewInt(1).Int64(), "Timestamp is invalid")

	// [new] dispute last miner as first miner
	MakeNew(t, timestamp)

	// check updated logs - should be 1
	logs = Show()
	assert.Equal(t, 1, logs, "There should be 1 dispute")

	// vote true as miner 2
	Vote(t)
}

func Vote(t *testing.T) {
	instance2 := minerCtx[1].Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	auth, _ := ops.PrepareEthTransaction(minerCtx[1])
	instance2.Vote(auth, big.NewInt(1), true)

	// check if vote passed
	instance := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	addr := minerCtx[1].Value(zapCommon.PublicAddress).(common.Address)
	did, _ := instance.DidVote(nil, big.NewInt(1), addr)
	assert.Equal(t, true, did, "Vote did not go through")
}

func MakeNew(t *testing.T, timestamp *big.Int) {
	auth, _ := ops.PrepareEthTransaction(minerCtx[0])
	instance2 := minerCtx[0].Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	addr1 := common.HexToAddress("0xb7278a61aa25c888815afc32ad3cc52ff24fe575")
	addr2 := common.HexToAddress("0x5f3f1dBD7B74C6B46e8c44f98792A1dAf8d69154")
	amt1 := big.NewInt(1000000000)
	instance := minerCtx[0].Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instance.IncreaseApproval(auth, addr1, amt1)

	auth, _ = ops.PrepareEthTransaction(minerCtx[0])
	instance.IncreaseApproval(auth, addr2, amt1)
	// allocate some funds for dispute fee
	auth, _ = ops.PrepareEthTransaction((minerCtx[5]))
	instance = minerCtx[5].Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instance.Allocate(auth, minerCtx[0].Value(zapCommon.PublicAddress).(common.Address), big.NewInt(1000000000))

	auth, _ = ops.PrepareEthTransaction(minerCtx[0])
	// dispute miner 5 as miner 1
	instance2.BeginDispute(auth, big.NewInt(1), timestamp, big.NewInt(4))

	// check if miner 5 is in dispute
	array := [32]byte{}
	data := []byte("requestId")
	data = crypto.Keccak256(data)
	for i := 0; i < 32; i++ {
		array[i] = data[i]
	}
	instancem := minerCtx[0].Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	RID, _ := instancem.GetDisputeUintVars(nil, big.NewInt(1), array)
	assert.Equal(t, big.NewInt(1), RID)
}

func Show() int {
	tokenAbi, _ := abi.JSON(strings.NewReader(zap1.ZapDisputeABI))
	contractAddress := ctx.Value(zapCommon.ContractAddress).(common.Address)
	client := ctx.Value(zapCommon.ClientContextKey).(rpc.ETHClient)
	header, _ := client.HeaderByNumber(ctx, nil)
	startBlock := big.NewInt(54) //big.NewInt(10e3 * 14)
	startBlock.Sub(header.Number, startBlock)
	newDisputeID := tokenAbi.Events["NewDispute"].ID
	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   nil,
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{newDisputeID}},
	}

	logs, _ := client.FilterLogs(ctx, query)
	return len(logs)
}

func StartMiners(t *testing.T) *big.Int {
	// make requestData
	auth, _ := ops.PrepareEthTransaction(minerCtx[5])

	// approve Zap & ZapMaster contracts
	addr1 := minerCtx[5].Value(zapCommon.ContractAddress).(common.Address)
	amt1 := big.NewInt(100000000000)
	instancet := minerCtx[5].Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instancet.IncreaseApproval(auth, addr1, amt1)

	auth, _ = ops.PrepareEthTransaction(minerCtx[5])
	instanceT := minerCtx[5].Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instanceT.Allocate(auth, minerCtx[5].Value(zapCommon.ContractAddress).(common.Address), big.NewInt(100000000000))

	// deposit stake for each miner
	for i := 0; i < 5; i++ {
		auth, _ = ops.PrepareEthTransaction((minerCtx[i]))
		instancet = minerCtx[i].Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
		instancet.IncreaseApproval(auth, minerCtx[i].Value(zapCommon.ContractAddress).(common.Address), big.NewInt(100000000000))

		auth, _ = ops.PrepareEthTransaction((minerCtx[i]))
		instanceV := minerCtx[i].Value(zapCommon.VaultTransactorContractContextKey).(*vault.VaultTransactor)
		instanceV.LockSmith(auth, minerCtx[i].Value(zapCommon.PublicAddress).(common.Address), minerCtx[i].Value(zapCommon.ContractAddress).(common.Address))

		auth, _ = ops.PrepareEthTransaction((minerCtx[i]))
		instanceZ := minerCtx[i].Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
		instanceZ.DepositStake(auth)
	}

	// requestData with tip
	auth, _ = ops.PrepareEthTransaction(minerCtx[5])
	instance1 := minerCtx[5].Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	instance1.RequestData(auth,
		"json(https://api.binance.com/api/v1/klines?symbol=BTCUSDT&interval=1d&limit=1).0.4", "BTC/USD",
		new(big.Int).SetInt64(10000), new(big.Int).SetInt64(1))

	// allocate 10000 zap for ZapMaster contract (bug - normally ZapMaster has 6000 zap initially)
	auth, _ = ops.PrepareEthTransaction(minerCtx[5])
	// instance1.UpdateBalanceAtNow(auth, addr1, big.NewInt(10000000))

	instance := minerCtx[5].Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	// check challenge
	challenge, requestID, difficulty, queryString, granularity, _, _ := instance.GetCurrentVariables(nil)
	assert.NotNil(t, challenge, "Challenge should not be nil")
	assert.Equal(t, big.NewInt(1), requestID, "RequestID should be 1")
	assert.GreaterOrEqual(t, difficulty.Int64(), big.NewInt(0).Int64(), "Difficulty should be at least 0")
	assert.Equal(t, "json(https://api.binance.com/api/v1/klines?symbol=BTCUSDT&interval=1d&limit=1).0.4", queryString, "Query string does not match")
	assert.Equal(t, big.NewInt(10000).Int64(), granularity.Int64(), "Granularity is incorrect")
	// assert.Equal(t, big.NewInt(1).Int64(), totalTip.Int64(), "Tips should be 1")
	// submit solution for 5 miners
	for i := 0; i < 5; i++ {
		value := int64(100) + int64(0)
		authx, _ := ops.PrepareEthTransaction(minerCtx[i])
		instancex := minerCtx[i].Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
		instancex.SubmitMiningSolution(authx, strconv.Itoa(58+i), big.NewInt(1), big.NewInt(value))
	}

	array := [32]byte{}
	data := []byte("timeOfLastNewValue")
	data = crypto.Keccak256(data)
	for i := 0; i < 32; i++ {
		array[i] = data[i]
	}

	uvar, _ := instance.GetUintVar(nil, array)
	return uvar
	// return nil
}
