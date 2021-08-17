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
	addr1 := common.HexToAddress("0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199")
	amt1 := big.NewInt(1000)

	master := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	prevBalance, _ := master.BalanceOf(nil, addr1)
	instance := ctx.Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instance.Transfer(auth, addr1, amt1)
	newBalance, _ := master.BalanceOf(nil, addr1)

	assert.Equal(t, prevBalance.Add(prevBalance, amt1), newBalance, "Transfer was not successful")
}
