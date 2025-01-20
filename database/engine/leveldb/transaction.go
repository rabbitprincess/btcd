package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
)

func NewTransaction(tx *leveldb.Transaction) *Transaction {
	return &Transaction{Transaction: tx}
}

type Transaction struct {
	*leveldb.Transaction
}

func (t *Transaction) Put(key, value []byte) error {
	return t.Transaction.Put(key, value, nil)
}

func (t *Transaction) Delete(key []byte) error {
	return t.Transaction.Delete(key, nil)
}

func (t *Transaction) Commit() error {
	return t.Transaction.Commit()
}
