package pebbledb

import (
	"github.com/btcsuite/btcd/database/engine"
	"github.com/cockroachdb/pebble"
)

func NewIterator(iter *pebble.Iterator) *Iterator {
	return &Iterator{Iterator: iter}
}

var _ engine.Iterator = (*Iterator)(nil)

type Iterator struct {
	*pebble.Iterator
}

func (i *Iterator) Seek(key []byte) bool {
	return i.Iterator.SeekGE(key)
}

func (i *Iterator) Release() {
	i.Iterator.Close()
}
