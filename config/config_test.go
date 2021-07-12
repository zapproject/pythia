package config

import (
	"testing"

	"github.com/zapproject/zap-miner/util"
)

func setup() error {
	err := ParseConfig("../config.json")
	if err != nil {
		return err
	}

	path := "../testConfig.json"
	err = util.ParseLoggingConfig(path)
	if err != nil {
		return err
	}

	return nil
}

func TestConfig(t *testing.T) {
	setup()
	cfg := GetConfig()
	if len(cfg.ContractAddress) == 0 {
		t.Fatal("Config did not parse correctly")
	}
}
