package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

func TestETHGasStation(t *testing.T) {
	tracker := &GasTracker{}
	opts := &rpc.MockOptions{
		ETHBalance:    big.NewInt(300000),
		Nonce:         1,
		GasPrice:      big.NewInt(7000000000),
		TokenBalance:  big.NewInt(0),
		Top50Requests: []*big.Int{},
	}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "ethGas_test"))
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.WithValue(context.Background(), zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.DBContextKey, DB)
	err = tracker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := DB.Get(db.GasKey)
	if err != nil {
		t.Fatal(err)
	}

	dec, err := hexutil.DecodeBig(string(v))
	if err != nil {
		t.Fatalf("Error in decoding gas price: %v", err)
	}

	t.Logf("Gas Price stored: %v\n", string(v))

	assert.IsType(t, big.NewInt(0), dec)
	assert.Equal(t, "GasTracker", tracker.String())
}

// test the suggested gas price client function
// this is based on the fallback logic when the ETH Gas Station is down
func TestGas(t *testing.T) {
	opts := &rpc.MockOptions{
		ETHBalance:    big.NewInt(300000),
		Nonce:         1,
		GasPrice:      big.NewInt(7000000000),
		TokenBalance:  big.NewInt(0),
		Top50Requests: []*big.Int{},
	}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_gas"))
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(context.Background(), zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.DBContextKey, DB)
	gas, err := client.SuggestGasPrice(ctx)
	if err != nil {
		t.Fatalf("Error in getting the mocked gas price.")
	}

	enc := hexutil.EncodeBig(gas)
	DB.Put(db.GasKey, []byte(enc))

	v, err := DB.Get(db.GasKey)
	if err != nil {
		t.Fatal(err)
	}
	b, err := hexutil.DecodeBig(string(v))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Gas Price stored: %v\n", string(v))

	assert.Equal(t, opts.GasPrice, b)
}
