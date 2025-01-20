package pebbledb

import (
	"runtime"

	"github.com/btcsuite/btcd/database/engine"
	"github.com/cockroachdb/pebble"
	"github.com/cockroachdb/pebble/bloom"
)

func NewDBEngine(create bool, dbPath string) (engine.Engine, error) {
	opts := &pebble.Options{
		Cache:                    pebble.NewCache(64 * 1024 * 1024), // 16 MB
		ErrorIfExists:            create,
		MaxOpenFiles:             64, // Fail if the database exists and create is true
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
		return nil, err
	}
	return &DB{DB: dbEngine}, nil
}

var _ engine.Engine = (*DB)(nil)

type DB struct {
	*pebble.DB
}

func (d *DB) NewTransaction() engine.Transaction {
	return NewTransaction(d.DB.NewBatch())
}

func (d *DB) NewSnapshot() engine.Snapshot {
	return NewSnapshot(d.DB.NewSnapshot())
}

func (d *DB) Close() error {
	return d.DB.Close()
}
