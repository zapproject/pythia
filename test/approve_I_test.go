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

func TestApprove(t *testing.T) {
	setup()

	auth, _ := ops.PrepareEthTransaction(ctx)
	owner := ctx.Value(zapCommon.PublicAddress).(common.Address)
	addr1 := common.HexToAddress("70997970C51812dc3A010C7d01b50e0d17dc79C8")
	amt1 := big.NewInt(1000)
	instance := ctx.Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instance.Approve(auth, addr1, amt1)

	master := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	approval1, _ := master.Allowance(nil, owner, addr1)
	assert.Equal(t, amt1, approval1, "Incorrect approved amount")
}
