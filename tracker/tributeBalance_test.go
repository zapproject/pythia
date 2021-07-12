package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

func TestTokenBalance(t *testing.T) {
	startBal := big.NewInt(456000)
	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1, GasPrice: big.NewInt(700000000),
		TokenBalance: startBal, Top50Requests: []*big.Int{}}
	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_Tokenbalance"))
	if err != nil {
		t.Fatal(err)
	}
	tracker := &TokenTracker{}
	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	err = tracker.Exec(ctx)
	if err != nil {
		t.Fatal(err)
	}
	v, err := DB.Get(db.TokenBalanceKey)
	if err != nil {
		t.Fatal(err)
	}
	b, err := hexutil.DecodeBig(string(v))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Token Balance stored: %v\n", b)
	if b.Cmp(startBal) != 0 {
		t.Fatalf("Balance from client did not match what should have been stored in DB. %s != %s", b, startBal)
	}
}
