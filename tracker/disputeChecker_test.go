package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

func TestDisputeCheckerInRange(t *testing.T) {
	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	disputeChecker := &disputeChecker{lastCheckedBlock: 500}

	DB, err := db.Open(filepath.Join(os.TempDir(), "disputeChecker_test"))
	assert.Nil(t, err)

	client := rpc.NewMockClientWithValues(opts)
	ctx := context.WithValue(context.Background(), zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.DBContextKey, DB)
	assert.Equal(t, DB, ctx.Value(zapCommon.DBContextKey))

	BuildIndexTrackers()

	ethUSDPairs := indexes["ETH/USD"]

	execEthUsdPsrs(ctx, t, ethUSDPairs)

	time.Sleep(2 * time.Second)

	execEthUsdPsrs(ctx, t, ethUSDPairs)

	ctx = context.WithValue(ctx, zapCommon.ContractAddress, common.Address{0x0000000000000000000000000000000000000000})
	assert.Equal(t, common.Address{0x0000000000000000000000000000000000000000}, ctx.Value(zapCommon.ContractAddress), "Contract address was not saved properlly")

	err = disputeChecker.Exec(ctx)
	assert.Nil(t, err)

	DB.Close()
}

func TestDisputeCheckerOutOfRange(t *testing.T) {
	cfg := config.GetConfig()
	cfg.DisputeThreshold = 0.000000001
	opts := &rpc.MockOptions{ETHBalance: big.NewInt(300000), Nonce: 1, GasPrice: big.NewInt(7000000000),
		TokenBalance: big.NewInt(0), Top50Requests: []*big.Int{}}
	disputeChecker := &disputeChecker{lastCheckedBlock: 500}
	DB, err := db.Open(filepath.Join(os.TempDir(), "disputeChecker_test"))
	assert.Nil(t, err)

	client := rpc.NewMockClientWithValues(opts)
	ctx := context.WithValue(context.Background(), zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.DBContextKey, DB)
	assert.Equal(t, DB, ctx.Value(zapCommon.DBContextKey))

	BuildIndexTrackers()
	ethUSDPairs := indexes["ETH/USD"]
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	time.Sleep(2 * time.Second)
	execEthUsdPsrs(ctx, t, ethUSDPairs)
	ctx = context.WithValue(ctx, zapCommon.ContractAddress, common.Address{0x0000000000000000000000000000000000000000})
	assert.Equal(t, common.Address{0x0000000000000000000000000000000000000000}, ctx.Value(zapCommon.ContractAddress), "Contract address was not saved properlly")

	err = disputeChecker.Exec(ctx)
	assert.Nil(t, err)

	DB.Close()
}

func execEthUsdPsrs(ctx context.Context, t *testing.T, psrs []*IndexTracker) {
	for _, psr := range psrs {
		err := psr.Exec(ctx)
		assert.Nilf(t, err, "failed to execute psr: %v", err)
		assert.NotEmpty(t, psr.String())
	}
}
