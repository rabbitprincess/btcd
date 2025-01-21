package pebbledb

import (
	"github.com/btcsuite/btcd/database/engine"
	"github.com/cockroachdb/pebble"
)

func NewTransaction(batch *pebble.Batch) engine.Transaction {
	return &Transaction{Batch: batch}
}

type Transaction struct {
	*pebble.Batch
}

func (t *Transaction) Put(key, value []byte) error {
	return t.Batch.Set(key, value, pebble.NoSync)
}

func (t *Transaction) Delete(key []byte) error {
	return t.Batch.Delete(key, pebble.NoSync)
}

func (t *Transaction) Discard() {
	t.Batch.Close()
}

func (t *Transaction) Commit() error {
	return t.Batch.Commit(pebble.Sync)
}
