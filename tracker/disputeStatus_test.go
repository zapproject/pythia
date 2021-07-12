package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/stretchr/testify/assert"
	"github.com/zapproject/zap-miner/common"

	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

func TestDisputeString(t *testing.T) {

	// Gets the DisputeTracker ID
	tracker := &DisputeTracker{}

	// Converts the tracker ID to a string
	disputeTrackerStr := tracker.String()

	// Asserts disputeTrackerStr equals "DisputeTracker"
	assert.Equal(t, disputeTrackerStr, "DisputeTracker", "should return 'DisputeTracker' string")
}

func TestDisputeStatus(t *testing.T) {

	startBal := big.NewInt(356000)

	// MockOptions are config options for the mock client
	opts := &rpc.MockOptions{ETHBalance: startBal,
		Nonce:         1,
		GasPrice:      big.NewInt(700000000),
		TokenBalance:  big.NewInt(0),
		Top50Requests: []*big.Int{},
		DisputeStatus: big.NewInt(1)}

	// NewMockClientWithValues creates a mock client with default values to return for calls
	client := rpc.NewMockClientWithValues(opts)

	// Mock DB
	DB, err := db.Open(filepath.Join(os.TempDir(), "test_disputeStatus"))

	// Path test_disputeStatus
	testDisputePath := filepath.Join(os.TempDir(), "test_disputeStatus")

	// DisputeStatus tracker
	tracker := &DisputeTracker{}

	ctx := context.WithValue(context.Background(), common.ClientContextKey, client)
	ctx = context.WithValue(ctx, common.DBContextKey, DB)
	trackerErr := tracker.Exec(ctx)

	// Get the dispute Status from the test DB as a bytes array
	getStatusKey, getStatusKeyErr := DB.Get(db.DisputeStatusKey)

	// Converts the bytes array to hex string
	// Converts the hexString to a BigNumber
	statusKey, statusKeyErr := hexutil.DecodeBig(string(getStatusKey))

	// 0 = Not Staked
	// The lowest status key
	statusKeyZero := big.NewInt(0)

	// 3 = InDispute
	// The maximum status key
	statusKeyThree := big.NewInt(3)

	t.Logf("Dispute Status stored: %v\n", statusKey)

	// Asserts err has no value
	assert.Nil(t, err, err)

	// Asserts DB has a value
	assert.NotNil(t, DB)

	// Asserts the testDisputePath exists
	assert.DirExists(t, testDisputePath)

	// Asserts trackerErr has no value
	assert.Nil(t, trackerErr, trackerErr)

	// Asserts getStatusKeyErr has no value
	assert.Nil(t, getStatusKeyErr)

	// Asserts getStatusKey has a value
	assert.NotNil(t, getStatusKey)

	// Asserts statusKeyErr has no value
	assert.Nil(t, statusKeyErr, statusKeyErr)

	// Asserts the statusKey has a value
	assert.NotNil(t, statusKey)

	// Asserts the status key is greater than or equal to 0
	// Keeps the range between 0-3
	assert.GreaterOrEqual(t, statusKey.Uint64(), statusKeyZero.Uint64(),

		"Status key not within range")

	// Asserts the status key is less than or equal to 3
	// Keeps the range between 0-3
	assert.LessOrEqual(t, statusKey.Uint64(), statusKeyThree.Uint64(),

		"Status key not within range")

	DB.Close()

}
