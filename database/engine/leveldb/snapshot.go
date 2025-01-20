package leveldb

import (
	"github.com/btcsuite/btcd/database/engine"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

func NewSnapshot(snapshot *leveldb.Snapshot) *Snapshot {
	return &Snapshot{Snapshot: snapshot}
}

var _ engine.Snapshot = (*Snapshot)(nil)

type Snapshot struct {
	*leveldb.Snapshot
}

func (s *Snapshot) Has(key []byte) (bool, error) {
	return s.Snapshot.Has(key, nil)
}

func (s *Snapshot) Get(key []byte) (val []byte, err error) {
	return s.Snapshot.Get(key, nil)
}

func (s *Snapshot) Release() {
	s.Snapshot.Release()
}

func (s *Snapshot) NewIterator(start, end []byte) (engine.Iterator, error) {
	return s.Snapshot.NewIterator(&util.Range{
		Start: start,
		Limit: end,
	}, nil), nil
}
