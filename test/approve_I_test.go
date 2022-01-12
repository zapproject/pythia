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
	addr1 := common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC")
	amt1 := big.NewInt(1000)
	instance := ctx.Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instance.Approve(auth, addr1, amt1)

	master := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	approval1, _ := master.Allowance(nil, owner, addr1)
	assert.Equal(t, amt1, approval1, "Incorrect approved amount")
}
