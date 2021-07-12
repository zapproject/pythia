package test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/zap-miner/common"
	zap "github.com/zapproject/zap-miner/contracts"
	"github.com/zapproject/zap-miner/ops"
	token "github.com/zapproject/zap-miner/token"
)

func TestApprove(t *testing.T) {
	setup()

	auth, _ := ops.PrepareEthTransaction(ctx)
	owner := ctx.Value(zapCommon.PublicAddress).(common.Address)
	addr1 := common.HexToAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3")
	amt1 := big.NewInt(1000)
	instance := ctx.Value(zapCommon.TokenTransactorContractContextKey).(*token.ZapTokenBSCTransactor)
	instance.Approve(auth, addr1, amt1)

	master := ctx.Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	approval1, _ := master.Allowance(nil, owner, addr1)
	assert.Equal(t, amt1, approval1, "Incorrect approved amount")
}
