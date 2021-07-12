// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package db

import (
	"io/ioutil"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/zapproject/zap-miner/config"
)

func OpenTestDB(cfg *config.Config) (DB, func() error, error) {
	tmpdir, err := ioutil.TempDir("", "test")
	if err != nil {
		return nil, nil, err
	}

	db, err := Open(tmpdir)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() error {
		if err := db.Close(); err != nil {
			if err != leveldb.ErrClosed {
				if err != nil {
					return err
				}
			}
		}
		if err := os.RemoveAll(tmpdir); err != nil {
			return err
		}
		return nil

	}
	return db, cleanup, nil
}
