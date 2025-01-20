package pebbledb

import "github.com/cockroachdb/pebble"

func NewTransaction(batch *pebble.Batch) *Transaction {
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

func (t *Transaction) Commit() error {
	return t.Batch.Commit(pebble.Sync)
}
