package test

import (
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
	setup()
	auth, _ := ops.PrepareEthTransaction(ctx)
	addr1 := common.HexToAddress("70997970C51812dc3A010C7d01b50e0d17dc79C8")
	amt1 := big.NewInt(1000)

	master := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	prevBalance, _ := master.BalanceOf(nil, addr1)
	instance := ctx.Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instance.Transfer(auth, addr1, amt1)
	newBalance, _ := master.BalanceOf(nil, addr1)

	assert.Equal(t, prevBalance.Add(prevBalance, amt1), newBalance, "Transfer was not successful")
}
