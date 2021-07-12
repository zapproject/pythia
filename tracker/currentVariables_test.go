package tracker

import (
	"context"
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/common"
	zapCommon "github.com/zapproject/zap-miner/common"
	"github.com/zapproject/zap-miner/config"
	zap "github.com/zapproject/zap-miner/contracts"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/zapproject/zap-miner/db"
	"github.com/zapproject/zap-miner/rpc"
)

func TestCurrentVarableString(t *testing.T) {

	// Gets the CurrentVariablesTracker string
	// Type is *tracker.CurrentVariablesTracker
	tracker := &CurrentVariablesTracker{}

	// Converts the tracker data type to a string
	currentVariablesStr := tracker.String()

	// Asserts currentVariablesStr has a value
	assert.NotNil(t, currentVariablesStr, "Should return 'CurrentVariablesTracker' string")

	// Asserts currentVariablesStr is equal to "CurrentVariablesTracker"
	assert.Equal(t, currentVariablesStr, "CurrentVariablesTracker", "Should return 'CurrentVariablesTracker' string")

}

func TestCurrentVariables(t *testing.T) {

	startBal := big.NewInt(356000)

	// Creates a hash
	hash := math.PaddedBigBytes(big.NewInt(256), 32)

	// Creates an array with the size of 32 bytes
	var b32 [32]byte

	// Iterates through hash
	for i, v := range hash {
		b32[i] = v

	}

	// Query String
	queryStr := "json(https://coinbase.com)"

	chal := &rpc.CurrentChallenge{ChallengeHash: b32, RequestID: big.NewInt(1),

		// Sets the difficulty to 500
		Difficulty: big.NewInt(500),

		// Sets the QueryString to "json(https://coinbase.com)"
		QueryString: queryStr,

		Granularity: big.NewInt(1000),
		Tip:         big.NewInt(0)}

	opts := &rpc.MockOptions{ETHBalance: startBal, Nonce: 1,

		GasPrice: big.NewInt(700000000),

		TokenBalance: big.NewInt(0),

		MiningStatus: true,

		Top50Requests: []*big.Int{},

		CurrentChallenge: chal}

	client := rpc.NewMockClientWithValues(opts)

	DB, err := db.Open(filepath.Join(os.TempDir(), "test_currentVariables"))
	assert.Nil(t, err)

	cfg := config.GetConfig()
	tracker := &CurrentVariablesTracker{}

	ctx := context.WithValue(context.Background(), zapCommon.ClientContextKey, client)
	ctx = context.WithValue(ctx, zapCommon.DBContextKey, DB)
	assert.Equal(t, DB, ctx.Value(zapCommon.DBContextKey))

	masterInstance := ctx.Value(zapCommon.MasterContractContextKey)

	if masterInstance == nil {

		contractAddress := common.HexToAddress(cfg.ContractAddress)

		// TODO create error state flag for mock client
		masterInstance, err = zap.NewZapMaster(contractAddress, client)
		assert.Nil(t, err, "Problem creating Zap Master instance")

		ctx = context.WithValue(ctx, zapCommon.MasterContractContextKey, masterInstance)
	}

	err = tracker.Exec(ctx)
	assert.Nil(t, err)

	// Gets the request ID as a bytes array
	v, err := DB.Get(db.RequestIdKey)
	assert.Nil(t, err)

	b, err := hexutil.DecodeBig(string(v))
	assert.Nil(t, err)
	assert.Equal(t, big.NewInt(1), b, "Current Request ID from client did not match what should have been stored in DB.")

	// Gets the QueryStringKey from the DB as a bytes array
	v, err = DB.Get(db.QueryStringKey)
	assert.Nil(t, err)

	// Converts the QueryStringKey from a bytes array to a string
	// Checks if the QueryStringKey returned from the DB is equivalent to the queryStr
	assert.Equalf(t, queryStr, string(v), "Expected query string to match test input: %s != %s\n", string(v), queryStr)
}
