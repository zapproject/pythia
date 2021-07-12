package test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	zap "github.com/zapproject/zap-miner/contracts"
	zap1 "github.com/zapproject/zap-miner/contracts1"
	"github.com/zapproject/zap-miner/ops"
	"github.com/zapproject/zap-miner/token"
	"github.com/zapproject/zap-miner/vault"
)

func TestStake(t *testing.T) {
	setup()
	setupOwner()

	// tmaster := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	// publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)
	// status, _, _ := tmaster.GetStakerInfo(nil, publicAddress)

	// assert.Equal(t, big.NewInt(1).Uint64(), status.Uint64(), "Staker status should be staked a inital miner 0 - status should be 1")
	RequestWithdraw(t)
	Withdraw(t)
	Deposit(t)
	// RequestWithdraw(t)
	// Withdraw(t)
}

func Deposit(t *testing.T) {
	cfg := config.GetConfig()

	tmaster := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)
	balance, _ := tmaster.BalanceOf(nil, publicAddress)
	status, _, _ := tmaster.GetStakerInfo(nil, publicAddress)
	assert.Equal(t, big.NewInt(0).Uint64(), status.Uint64(), "Staker status does not match - status should be 0")

	dat := crypto.Keccak256([]byte("stakeAmount"))
	var dat32 [32]byte
	copy(dat32[:], dat)
	stakeAmt, _ := tmaster.GetUintVar(nil, dat32)
	assert.Greater(t, balance.Cmp(stakeAmt), 0, "Account 0 does not have enough Zap to stake")

	// allocate stake funds to this account
	auth, _ := ops.PrepareEthTransaction(minerCtx[5])
	instanceT := minerCtx[5].Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instanceT.Allocate(auth, publicAddress, big.NewInt(100000000000))

	balance, _ = tmaster.BalanceOf(nil, publicAddress)

	// call vault locksmith
	instanceV := ctx.Value(zapCommon.VaultTransactorContractContextKey).(*vault.VaultTransactor)
	auth, _ = ops.PrepareEthTransaction(ctx)
	instanceV.LockSmith(auth, publicAddress, common.HexToAddress(cfg.ContractAddress))

	// approve zap master for stake amount
	auth, _ = ops.PrepareEthTransaction(ctx)
	instanceT.IncreaseApproval(auth, common.HexToAddress(cfg.ContractAddress), stakeAmt)

	// deposit stake
	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	auth, _ = ops.PrepareEthTransaction(ctx)
	instance2.DepositStake(auth)

	status, _, _ = tmaster.GetStakerInfo(nil, publicAddress)
	assert.Equal(t, big.NewInt(1).Uint64(), status.Uint64(), "Staker status does not match - status should be 1")

	newBalance, _ := tmaster.BalanceOf(nil, publicAddress)
	assert.Equal(t, newBalance.Add(stakeAmt, newBalance), balance, "Stake amount was not transfered to vault")
}

func Withdraw(t *testing.T) {
	// check staker status
	tmaster := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)
	status, _, _ := tmaster.GetStakerInfo(nil, publicAddress)
	assert.Equal(t, big.NewInt(2).Uint64(), status.Uint64(), "Staker has not requested for withdrawal - status should be 2")

	balance, _ := tmaster.BalanceOf(nil, publicAddress)

	// withdraw stake
	auth, _ := ops.PrepareEthTransaction(ctx)
	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	instance2.WithdrawStake(auth)
	status, _, _ = tmaster.GetStakerInfo(nil, publicAddress)
	assert.Equal(t, big.NewInt(0).Uint64(), status.Uint64(), "Withdraw was not successful - status should be 0")

	// The new balance should = the prev balance - dispute fee + stake amount
	newBalance, _ := tmaster.BalanceOf(nil, publicAddress)
	dat := crypto.Keccak256([]byte("stakeAmount"))
	var dat32 [32]byte
	copy(dat32[:], dat)
	stakeAmt, _ := tmaster.GetUintVar(nil, dat32)

	balance = balance.Sub(balance, stakeAmt)
	assert.Equal(t, newBalance, balance.Add(balance, big.NewInt(1000000)), "Stake amount was not transfered from vault to user")
}

func RequestWithdraw(t *testing.T) {
	tmaster := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)
	status, _, _ := tmaster.GetStakerInfo(nil, publicAddress)
	assert.Equal(t, big.NewInt(1).Uint64(), status.Uint64(), "Account 0 is not staked - status should be 1")

	auth, _ := ops.PrepareEthTransaction(ctx)

	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
	instance2.RequestStakingWithdraw(auth)

	status, _, _ = tmaster.GetStakerInfo(nil, publicAddress)
	assert.Equal(t, big.NewInt(2).Uint64(), status.Uint64(), "Request withdraw was not successful - status should be 2")
}
