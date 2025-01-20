package engine

type Snapshot interface {
	Get(key []byte) ([]byte, error)
	Has(key []byte) (bool, error)

	NewIterator(start, end []byte) (Iterator, error)

	Releaser
}
