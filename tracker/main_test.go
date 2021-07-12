package tracker

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zapproject/zap-miner/apiOracle"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/util"
)

// TODO: Set threshold low and test the  "out of range" failure
var configJSON = `{
	"publicAddress":"0000000000000000000000000000000000000000",
	"privateKey":"1111111111111111111111111111111111111111111111111111111111111111",
	"contractAddress":"0xCD8a1C3ba11CF5ECfa6267617243239504a98d90",
	"trackers": ["indexers", "balance", "currentVariables", "disputeStatus", "gas", "disputeChecker"],
	"IndexFolder": "..",
	"disputeThreshold": 1.0, 
	"disputeTimeDelta": "50s"
}
`

func TestMain(m *testing.M) {
	err := config.ParseConfigBytes([]byte(configJSON))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to parse mock config: %v\n", err)
		os.Exit(-1)
	}
	util.ParseLoggingConfig("")
	apiOracle.EnsureValueOracle()
	os.Exit(m.Run())
}

func TestMockedConfig(t *testing.T) {
	// assert for configs
	cfg := config.GetConfig()
	assert.Equal(t, "0000000000000000000000000000000000000000", cfg.PublicAddress)
	assert.Equal(t, "1111111111111111111111111111111111111111111111111111111111111111", cfg.PrivateKey)
	assert.Equal(t, "0x724D1B69a7Ba352F11D73fDBdEB7fF869cB22E19", cfg.ContractAddress)
	assert.Equal(t, "..", cfg.IndexFolder)
	assert.Equal(t, 1.0, cfg.DisputeThreshold)
	assert.Equal(t, config.Duration(config.Duration{Duration: 50000000000}), cfg.DisputeTimeDelta)

	// assert trackers individually
	assert.Equal(t, "indexers", cfg.Trackers[0])
	assert.Equal(t, "balance", cfg.Trackers[1])
	assert.Equal(t, "currentVariables", cfg.Trackers[2])
	assert.Equal(t, "disputeStatus", cfg.Trackers[3])
	assert.Equal(t, "gas", cfg.Trackers[4])
	assert.Equal(t, "disputeChecker", cfg.Trackers[5])
}
