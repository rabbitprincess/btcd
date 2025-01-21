package engine

type Engine interface {
	Transaction() (Transaction, error)
	Snapshot() (Snapshot, error)
	Close() error
}

type Transaction interface {
	Put(key, value []byte) error
	Delete(key []byte) error
	Commit() error
	Discard()
}
