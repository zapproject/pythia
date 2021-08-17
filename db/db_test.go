package db

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/zapproject/pythia/config"
	"github.com/zapproject/pythia/util"
)

var db DB
var err error

func setup(t *testing.T) {
	db, err = Open(filepath.Join(os.TempDir(), "test_db"))
	err := config.ParseConfig("../config.json")
	if err != nil {
		t.Error(err)
	}
	path := "../testConfig.json"
	err = util.ParseLoggingConfig(path)
	if err != nil {
		fmt.Errorf("Can't parse logging config for test.")
	}
	config.GetConfig()
}

func TestDB(t *testing.T) {
	if db == nil {
		setup(t)
	}

	defer db.Close()
	err = db.Put("sample", []byte("sample_value"))
	if err != nil {
		t.Error(err)
	}
	b, err := db.Has("sample")
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error("Expected key to be present in DB")
	}
	v, _ := db.Get("sample")
	s := string(v)
	if s != "sample_value" {
		t.Error("Get value doesn't match original: " + s + " != sample_value")
	}
	t.Log("Retrieved " + s)
}

func TestKeys(t *testing.T) {
	if db == nil {
		setup(t)
	}
	defer db.Close()
	if !isKnownKey("top_50_requestIds") {
		t.Error("Error with getting known keys")
	}

	if isKnownKey("rubbish") {
		t.Error("Shouldn't be have retrieved this key")
	}
}

func TestLocalDataProxy(t *testing.T) {
	if db == nil {
		setup(t)
	}
	defer db.Close()

	local_proxy, err := OpenLocalProxy(db)
	if err != nil {
		t.Error("Error with creating local proxy: ", err)
	}

	local_proxy.Put("sample", []byte("sample_value"))
	b, err := db.Has("sample")
	if err != nil {
		t.Error(err)
	}
	if !b {
		t.Error("Expected key to be present in DB")
	}

	v, _ := local_proxy.Get("sample")
	s := string(v)
	if s != "sample_value" {
		t.Error("Get value doesn't match original: " + s + " != sample_value")
	}
	t.Log("Retrieved " + s)
}
