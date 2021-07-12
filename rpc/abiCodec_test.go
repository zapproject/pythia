package rpc

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/zapproject/zap-miner/config"
	"github.com/zapproject/zap-miner/util"
)

func setup() {
	err := config.ParseConfig("../config.json")
	if err != nil {
		fmt.Errorf("Can't parse config for test.")
	}
	path := "../testConfig.json"
	err = util.ParseLoggingConfig(path)
	if err != nil {
		fmt.Errorf("Can't parse logging config for test.")
	}
}

func TestABICodec(t *testing.T) {
	setup()
	codec, err := BuildCodec()
	if err != nil {
		t.Fatal(err)
	}
	m := codec.methods["0xe1eee6d6"] // method signature
	if m == nil {
		t.Fatal("Missing expected method matching test sig")
	}
	if m.Name != "getRequestVars" {
		t.Fatalf("Method name is unexpected. %s != getRequestVars", m.Name)
	}

	//string, string, bytes32,  uint, uint, uint
	var hash [32]byte
	copy([]byte("12345"), hash[:])
	data, err := m.Outputs.Pack("someQueryString", "ETH/USD", hash, big.NewInt(1000), big.NewInt(0), big.NewInt(0))
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(data); i += 32 {
		hex := hexutil.Encode(data[i : i+32])
		fmt.Println(hex)
	}
	//fmt.Printf("Encoded payload: %s\n", hexutil.Encode(data))
}
