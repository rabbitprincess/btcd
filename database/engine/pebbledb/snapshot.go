package pebbledb

import (
	"github.com/btcsuite/btcd/database/engine"
	"github.com/cockroachdb/pebble"
)

func NewSnapshot(snapshot *pebble.Snapshot) *Snapshot {
	return &Snapshot{Snapshot: snapshot}
}

var _ engine.Snapshot = (*Snapshot)(nil)

type Snapshot struct {
	*pebble.Snapshot
}

func (s *Snapshot) Has(key []byte) (bool, error) {
	val, err := s.Get(key)
	if err != nil {
		return false, err
	}
	return val != nil, nil
}

func (s *Snapshot) Get(key []byte) (val []byte, err error) {
	val, closer, err := s.Snapshot.Get(key)
	if closer != nil {
		closer.Close()
	}
	return val, err
}

func (s *Snapshot) Release() {
	s.Close()
}

func (s *Snapshot) NewIterator(start, end []byte) (engine.Iterator, error) {
	iter, err := s.Snapshot.NewIter(&pebble.IterOptions{
		LowerBound: start,
		UpperBound: end,
	})
	if err != nil {
		return nil, err
	}
	return NewIterator(iter), nil
}
