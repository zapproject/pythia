package routes

import (
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/prometheus/tsdb/testutil"
	"github.com/zapproject/pythia/db"
)

type BalResult struct {
	Balance string `json:"balance"`
}

func TestBalanceHandler(t *testing.T) {
	setup(t)
	h := &BalanceHandler{}

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_balance_rest"))
	if err != nil {
		t.Fatal(err)
	}

	bigBal := big.NewInt(350000)
	raw_bal := hexutil.EncodeBig(bigBal)
	DB.Put(db.BalanceKey, []byte(raw_bal))

	router := NewRouter(DB)
	router.AddRoute("/balance", h)

	http.Handle("/", router)

	go func() { http.ListenAndServe("localhost:5001", nil) }()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost:5001/balance") // not sure how to make requests for routes.router
	if err != nil {
		t.Fatalf("Error getting balance from endpoint: %v", err)
	}
	defer resp.Body.Close()

	var bal BalResult
	t.Logf("Balance from response: %v", resp.Body)
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&bal)
	if !strings.Contains(bal.Balance, "0x") {
		testutil.Ok(t, errors.New("Missing balance in response"))
	} else {
		t.Logf("Retrieved balance from server: %+v\n", bal)
	}

	// resBal := bal.Balance

	// if resBal.Cmp(bigBal) != 0 {
	// 	t.Fatal("Starting and result balances did not match")
	// } else {
	// 	t.Logf("Ending balance: %s\n", res.Balance)
	// }
}
