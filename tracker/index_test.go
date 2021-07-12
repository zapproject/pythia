package tracker

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/assert"
	"github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/db"
)

//./runTest.sh TestPSR tracker

func TestPSR(t *testing.T) {

	//client.Timeout = 5

	// db = Open the database using the given DB file as its data store
	db, dbErr := db.Open(filepath.Join(os.TempDir(), "test_psrFetch"))

	// File path for test_psrFetch
	testPsrFilePath := filepath.Join(os.TempDir(), "test_psrFetch")

	ctx := context.Background()
	ctx = context.WithValue(ctx, common.DBContextKey, db)

	// Gets the array of price index trackers
	psr, psrErr := BuildIndexTrackers()

	// psr array length
	psrLength := len(psr)

	for idx := range psr {

		psrIdxErr := psr[idx].Exec(ctx)

		psrStr := psr[idx].String()

		assert.Nil(t, psrIdxErr, "failed to execute psr: %s %v", psrStr, psrIdxErr)
	}

	// Gets the value for request ID 1
	reqId1Val, valErr := db.Get(fmt.Sprintf("qv_%d", 1))

	// Converts reqId1Val from a bytes array to an integer
	intVal, intErr := hexutil.DecodeBig(string(reqId1Val))

	fmt.Println("DB value", intVal)

	// Asserts db has a value
	assert.NotNil(t, db, dbErr)

	// Asserts the filepath for test_psrFetch exists
	assert.DirExists(t, testPsrFilePath, "Error getting the test_psrFetch path")

	// Asserts dbErr has no value
	assert.Nil(t, dbErr, dbErr)

	// Asserts psr length is greater than 0
	assert.Greater(t, len(psr), 0)

	// Asserts the psr length is valid
	assert.Len(t, psr, psrLength)

	// Asserts psrErr has no value
	assert.Nil(t, psrErr, psrErr)

	// Asserts valErr has no value
	assert.Nil(t, valErr, valErr)

	// Asserts reqId1Val has a value
	assert.NotNil(t, reqId1Val, fmt.Errorf("Expected a value stored for request ID 1"))

	// Asserts intErr has no value
	assert.Nil(t, intErr, intErr)

}
