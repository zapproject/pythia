package test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/pythia/common"
	zap "github.com/zapproject/pythia/contracts"
	"github.com/zapproject/pythia/util"
)

func TestBalance(t *testing.T) {
	setupOwner()

	addr := minerCtx[5].Value(zapCommon.PublicAddress).(common.Address)
	instance := minerCtx[5].Value(zapCommon.MasterContractContextKey).(*zap.ZapMaster)
	zapBalance, _ := instance.BalanceOf(nil, addr)
	assert.Equal(t, "531000000.00", util.FormatERC20Balance(zapBalance), "Incorrect Balance")
}
