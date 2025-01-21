package engine

type Engine interface {
	Init(create bool, dbPath string) error

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
