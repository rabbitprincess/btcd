package pebbledb

import (
	"runtime"

	"github.com/btcsuite/btcd/database/engine"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/bloom"
)

var _ engine.Engine = (*DB)(nil)

type DB struct {
	*pebble.DB
}

func (d *DB) Init(create bool, dbPath string) error {
	opts := &pebble.Options{
		Cache:                    pebble.NewCache(64 * 1024 * 1024), // 64 MB
		ErrorIfExists:            create,                            // Fail if the database exists and create is true
		MaxOpenFiles:             64,
		MaxConcurrentCompactions: runtime.NumCPU,
		Levels: []pebble.LevelOptions{
			{TargetFileSize: 2 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 4 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 8 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 16 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 32 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 64 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
			{TargetFileSize: 128 * 1024 * 1024, FilterPolicy: bloom.FilterPolicy(10)},
		},
	}
	opts.Experimental.ReadSamplingMultiplier = -1
	dbEngine, err := pebble.Open(dbPath, opts)
	if err != nil {
		return err
	}

	d.DB = dbEngine
	return nil
}

func (d *DB) Transaction() (engine.Transaction, error) {
	return NewTransaction(d.DB.NewBatch()), nil
}

func (d *DB) Snapshot() (engine.Snapshot, error) {
	return NewSnapshot(d.DB.NewSnapshot()), nil
}

func (d *DB) Close() error {
	return d.DB.Close()
}
