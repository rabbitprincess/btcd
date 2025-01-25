package pebbledb

import (
	"github.com/btcsuite/btcd/database/engine"
	"github.com/cockroachdb/pebble"
)

func NewSnapshot(snapshot *pebble.Snapshot) engine.Snapshot {
	return &Snapshot{Snapshot: snapshot}
}

type Snapshot struct {
	*pebble.Snapshot
}

func (s *Snapshot) Has(key []byte) (bool, error) {
	val, err := s.Get(key)
	if err == pebble.ErrNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return val != nil, nil
}

func (s *Snapshot) Get(key []byte) (val []byte, err error) {
	ori, closer, err := s.Snapshot.Get(key)
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	val = make([]byte, len(ori))
	copy(val, ori)
	return val, nil
}

func (s *Snapshot) Release() {
	s.Close()
}

func (s *Snapshot) NewIterator(slice *engine.Range) engine.Iterator {
	iter, _ := s.Snapshot.NewIter(&pebble.IterOptions{
		LowerBound: slice.Start,
		UpperBound: slice.Limit,
	})
	iter.SeekLT(slice.Start)
	return NewIterator(iter)
}
