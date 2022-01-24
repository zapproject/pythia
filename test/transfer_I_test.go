package test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/pythia/common"
	zap "github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/ops"
	token "github.com/zapproject/pythia/token"
)

func TestTransfer(t *testing.T) {
	setupMiners()
	setupOwner()

	addr1 := minerCtx[1].Value(zapCommon.PublicAddress).(common.Address)
	amt1 := big.NewInt(1000)

	// allocate 1000 zap to owner
	auth, _ := ops.PrepareEthTransaction(minerCtx[5])
	instance := minerCtx[5].Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instance.Allocate(auth, minerCtx[0].Value(zapCommon.PublicAddress).(common.Address), amt1)

	// recipient previous balance
	master := minerCtx[1].Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	prevBalance, _ := master.BalanceOf(nil, addr1)
	fmt.Println(prevBalance)

	// transfer to recipient
	auth, _ = ops.PrepareEthTransaction(minerCtx[0])
	instance = minerCtx[0].Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instance.Transfer(auth, addr1, amt1)

	// recipient new balance
	newBalance, _ := master.BalanceOf(nil, addr1)
	fmt.Println(newBalance)

	assert.Equal(t, prevBalance.Add(prevBalance, amt1), newBalance, "Transfer was not successful")
}
