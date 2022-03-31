package tracker

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zapproject/pythia/apiOracle"
	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/util"
)

// TODO: Set threshold low and test the  "out of range" failure
var configJSON = `{
	"zapTokenAddress": "0x5fbdb2315678afecb367f032d93f642f64180aa3",
    "contractAddress": "0x724D1B69a7Ba352F11D73fDBdEB7fF869cB22E19",
    "vaultAddress": "0xa513E6E4b8f2a923D98304ec87F64353C4D5C853",
    "nodeURL": "http://127.0.0.1:8545/",
    "publicAddress": "0000000000000000000000000000000000000000",
    "privateKey": "1111111111111111111111111111111111111111111111111111111111111111",
    "serverHost": "localhost",
    "serverPort": 5001,
    "ethClientTimeout": 3000,
    "trackerCycle": 10,
    "serverWhitelist": [
        "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266", 
        "0xcd3B766CCDd6AE721141F452C550Ca635964ce71",
        "0x2546BcD3c84621e976D8185a91A922aE77ECEc30",
        "0xbDA5747bFD65F08deb54cb465eB87D40e51B197E",
        "0xdD2FD4581271e230360230F9337D5c0430Bf44C0",
        "0x8626f6940E2eb28930eFb4CeF49B2d1F2C9C1199"
    ],
    "gasMultiplier": 1,
    "gasMax":30,
    "trackers": ["indexers", "balance", "currentVariables", "disputeStatus", "gas", "disputeChecker"],
	"IndexFolder": "..",
    "dbFile": "zapDB",
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
