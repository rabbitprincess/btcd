package engine

type Engine interface {
	NewTransaction() Transaction
	NewSnapshot() Snapshot
	Close() error
}

type Transaction interface {
	Put(key, value []byte) error
	Delete(key []byte) error
	Commit() error
}
