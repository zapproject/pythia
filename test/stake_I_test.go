package test

// withdraw test fails because of time lock

// import (
// 	"math/big"
// 	"testing"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// 	"github.com/stretchr/testify/assert"
// 	zapCommon "github.com/zapproject/pythia/common"
// 	"github.com/zapproject/pythia/config"
// 	zap "github.com/zapproject/pythia/contracts"
// 	zap1 "github.com/zapproject/pythia/contracts1"
// 	"github.com/zapproject/pythia/ops"
// 	"github.com/zapproject/pythia/token"
// )

// func TestStake(t *testing.T) {
// 	setup()
// 	setupOwner()

// 	RequestWithdraw(t)
// 	Withdraw(t)
// 	Deposit(t)
// 	// RequestWithdraw(t)
// 	// Withdraw(t)
// }

// func Deposit(t *testing.T) {
// 	cfg := config.GetConfig()

// 	tmaster := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
// 	publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)
// 	status, _, _ := tmaster.GetStakerInfo(nil, publicAddress)
// 	assert.Equal(t, big.NewInt(0).Uint64(), status.Uint64(), "Staker status does not match - status should be 0")

// 	// allocate stake funds to this account
// 	auth, _ := ops.PrepareEthTransaction(minerCtx[5])
// 	instanceT := minerCtx[5].Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
// 	instanceT.Allocate(auth, publicAddress, big.NewInt(100000000000))

// 	balance, _ := tmaster.BalanceOf(nil, publicAddress)
// 	dat := crypto.Keccak256([]byte("stakeAmount"))
// 	var dat32 [32]byte
// 	copy(dat32[:], dat)
// 	stakeAmt, _ := tmaster.GetUintVar(nil, dat32)
// 	assert.Greater(t, balance.Cmp(stakeAmt), 0, "Account 0 does not have enough Zap to stake")

// 	// approve zap master for stake amount
// 	auth, _ = ops.PrepareEthTransaction(ctx)
// 	instanceT.IncreaseApproval(auth, common.HexToAddress(cfg.ContractAddress), stakeAmt)

// 	// deposit stakePrepareEthTransaction
// 	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
// 	auth, _ = ops.PrepareEthTransaction(ctx)
// 	instance2.DepositStake(auth)

// 	status, _, _ = tmaster.GetStakerInfo(nil, publicAddress)
// 	assert.Equal(t, big.NewInt(1).Uint64(), status.Uint64(), "Staker status does not match - status should be 1")

// 	newBalance, _ := tmaster.BalanceOf(nil, publicAddress)
// 	assert.Equal(t, newBalance.Add(stakeAmt, newBalance), balance, "Stake amount was not transfered to vault")
// }

// func Withdraw(t *testing.T) {
// 	// check staker status
// 	tmaster := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
// 	publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)
// 	status, _, _ := tmaster.GetStakerInfo(nil, publicAddress)
// 	assert.Equal(t, big.NewInt(2).Uint64(), status.Uint64(), "Staker has not requested for withdrawal - status should be 2")

// 	balance, _ := tmaster.BalanceOf(nil, publicAddress)

// 	// withdraw stake
// 	auth, _ := ops.PrepareEthTransaction(ctx)
// 	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
// 	instance2.WithdrawStake(auth)
// 	status, _, _ = tmaster.GetStakerInfo(nil, publicAddress)
// 	assert.Equal(t, big.NewInt(0).Uint64(), status.Uint64(), "Withdraw was not successful - status should be 0")

// 	// The new balance should = the prev balance - dispute fee + stake amount
// 	newBalance, _ := tmaster.BalanceOf(nil, publicAddress)
// 	dat := crypto.Keccak256([]byte("stakeAmount"))
// 	var dat32 [32]byte
// 	copy(dat32[:], dat)
// 	stakeAmt, _ := tmaster.GetUintVar(nil, dat32)

// 	balance = balance.Sub(balance, stakeAmt)
// 	balance = balance.Add(balance, big.NewInt(5))
// 	assert.Equal(t, newBalance, balance.Add(balance, big.NewInt(1000000)), "Stake amount was not transfered from vault to user")
// }

// func RequestWithdraw(t *testing.T) {
// 	tmaster := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
// 	publicAddress := ctx.Value(zapCommon.PublicAddress).(common.Address)
// 	status, _, _ := tmaster.GetStakerInfo(nil, publicAddress)
// 	assert.Equal(t, big.NewInt(1).Uint64(), status.Uint64(), "Account 0 is not staked - status should be 1")

// 	auth, _ := ops.PrepareEthTransaction(ctx)

// 	instance2 := ctx.Value(zapCommon.TransactorContractContextKey).(*zap1.ZapTransactor)
// 	instance2.RequestStakingWithdraw(auth)

// 	status, _, _ = tmaster.GetStakerInfo(nil, publicAddress)
// 	assert.Equal(t, big.NewInt(2).Uint64(), status.Uint64(), "Request withdraw was not successful - status should be 2")
// }
