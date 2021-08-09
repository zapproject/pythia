package local

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/prometheus/tsdb/testutil"
	"github.com/zapproject/pythia/common"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/db"
	"github.com/zapproject/pythia/util"
)

type PriceResult struct {
	Pair  string `json:"pair"`
	Value string `json:"value"`
}

func setup(t *testing.T) error {
	err := config.ParseConfig("../config.json")
	if err != nil {
		return fmt.Errorf("Can't parse config for test: %s", err)
	}
	path := "../loggingConfig.json"
	err = util.ParseLoggingConfig(path)
	if err != nil {
		return fmt.Errorf("Can't parse logging config for test: %s", err)
	}

	return nil
}

func start(s *http.Server) {
	go func() {
		fmt.Printf("Starting server on %+v\n", s.Addr)
		// returns ErrServerClosed on graceful close
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			// NOTE: there is a chance that next line won't have time to run,
			// as main() doesn't wait for this goroutine to stop. don't use
			// code with race conditions like these for production. see post
			// comments below on more discussion on how to handle this.
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()
}

func TestIndexHander(t *testing.T) {
	err := setup(t)
	if err != nil {
		t.Errorf("Issue with setting up test\n%s", err)
	}

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_index_route"))
	if err != nil {
		t.Fatalf("Error with opening DB for test: %v", err)
	}

	// proxy, err := db.OpenRemoteDB(DB)
	// if err != nil {
	// 	t.Fatalf("Error in opening remote proxy to set up rest server: ", err)
	// }

	ctx := context.WithValue(context.Background(), common.DBContextKey, DB)
	server := CreateLocal(ctx)
	server.Start()

	t.Logf("Spinning up server...")
	time.Sleep(5 * time.Second)

	pairs := []string{
		"/xcd",
		"/ttd",
		"/jmd",
		"/bbd",
		"/bsd",
		"/bzd",
		"/htg",
		"/srd",
		"/bmd",
		"/kyd",
		"/aud",
		"/gbp",
		"/cad",
		"/dkk",
		"/eur",
		"/hkd",
		"/jpy",
		"/nzd",
		"/nok",
		"/zar",
		"/sek",
		"/chf",
		"/try",
		"/aed",
	}

	for _, pair := range pairs {
		endpoint := "http://localhost:5002" + pair
		resp, err := http.Get(endpoint)
		if err != nil {
			t.Fatalf("Error in getting the XCD/USD price: %v", err)
		}

		defer resp.Body.Close()

		var price PriceResult
		t.Logf("This is the response: %v", resp.Body)

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&price)

		if !strings.Contains(price.Pair, "0x") || !strings.Contains(price.Value, "0x") {
			testutil.Ok(t, err)
		}

		t.Logf("Decoded Response: %s\n", price)
	}

	t.Logf("Winding down server...")
	server.Stop()
	time.Sleep(5 * time.Second)

}
