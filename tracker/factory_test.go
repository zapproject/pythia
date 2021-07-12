package tracker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBalanceTracker(t *testing.T) {

	// Gets the BalanceTracker string
	// Type is *tracker.BalanceTracker
	balanceTracker, _ := createTracker("balance")

	// Converts the tracker data type to a string
	balanceTrackerStr := balanceTracker[0].String()

	// Asserts balanceTrackerStr has a value
	assert.NotNil(t, balanceTrackerStr)

	// Asserts balanceTrackerStr equals "BalanceTracker"
	assert.Equal(t, balanceTrackerStr, "BalanceTracker",

		"Expected BalanceTracker but got %s", balanceTrackerStr)

}

func TestCreateCurrentVariablesTracker(t *testing.T) {

	// Gets the CurrentVariablesTracker string
	// Type is *tracker.CurrentVariablesTracker
	currentVariablesTracker, _ := createTracker("currentVariables")

	// Converts the tracker data type to a string
	currentVariablesStr := currentVariablesTracker[0].String()

	// Asserts currentVariablesStr has a value
	assert.NotNil(t, currentVariablesStr)

	// Asserts currentVariablesStr is equal to "CurrentVariablesTracker"
	assert.Equal(t, currentVariablesStr, "CurrentVariablesTracker",

		"Expected CurrentVariablesTracker but got %s", currentVariablesStr)

}

func TestCreateDisputeStatusTracker(t *testing.T) {

	// Gets the DisputeStatusTracker string
	// Type is *tracker.DisputeTracker
	disputeStatusTracker, _ := createTracker("disputeStatus")

	// Converts the tracker data type to a string
	disputeStatusStr := disputeStatusTracker[0].String()

	// Asserts disputeStatusStr has a value
	assert.NotNil(t, disputeStatusStr)

	// Asserts disputeStatusStr is equal to DisputeTracker
	assert.Equal(t, disputeStatusStr, "DisputeTracker",

		"Expected DisputeTracker but got %s", disputeStatusStr)

}

func TestCreateGasTracker(t *testing.T) {

	// Gets the GasTracker string
	// Type is *tracker.GasTracker
	gasTracker, _ := createTracker("gas")

	// Converts the tracker data type to a string
	gasTrackerStr := gasTracker[0].String()

	// Asserts gasTrackerStr has a value
	assert.NotNil(t, gasTrackerStr)

	// Assert gasTrackerStr equals "GasTracker"
	assert.Equal(t, gasTrackerStr, "GasTracker",

		"Expected GasTracker but got %s", gasTrackerStr)

}

func TestCreateTokenBalanceTracker(t *testing.T) {

	// Gets the tokenBalanceTracker string
	// Type is *tracker.TokenTracker
	tokenBalanceTracker, _ := createTracker("tokenBalance")

	// Converts the tracker data type to a string
	tokenBalanceStr := tokenBalanceTracker[0].String()

	// Asserts tokenBalanceStr has a value
	assert.NotNil(t, tokenBalanceStr)

	// Asserts tokenBalanceStr equals "TokenTracker"
	assert.Equal(t, tokenBalanceStr, "TokenTracker",

		"Expected TokenTracker but got %s", tokenBalanceStr)

}

func TestCreateIndexesTracker(t *testing.T) {

	// Gets the array of price query strings
	indexersTracker, err := createTracker("indexers")

	// Asserts err has no value
	assert.Nil(t, err, "Could not build IndexTracker")

	// Asserts indexersTracker length is not 0
	assert.NotEqual(t, len(indexersTracker), 0,

		"Could not build all IndexTrackers: only tracking %d indexes", len(indexersTracker))

}

func TestCreateDisputeCheckerTracker(t *testing.T) {

	// Gets the DisputeChecker string
	disputeChecker, _ := createTracker("disputeChecker")

	// Converts the tracker data type to a string
	disputeCheckerStr := disputeChecker[0].String()

	// Asserts disputeCheckerStr has a value
	assert.NotNil(t, disputeCheckerStr)

	// Asserts disputeCheckerStr equals "DisputeChecker"
	assert.Equal(t, disputeCheckerStr, "DisputeChecker",

		"Expected DisputeChecker but got %s", disputeCheckerStr)

}

func TestCreateBadTracker(t *testing.T) {

	// Creates a bad tracker as an empty array
	badTracker, err := createTracker("badTracker")

	// Asserts err has a value
	assert.NotNil(t, err,
		"Expected error but instead received this tracker: %s", badTracker)

	// Asserts the empty badTracker array length equals 0
	assert.Len(t, badTracker, 0)
}
