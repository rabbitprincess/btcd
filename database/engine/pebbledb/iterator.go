package pebbledb

import "github.com/cockroachdb/pebble"

func NewIterator(iter *pebble.Iterator) *Iterator {
	return &Iterator{Iterator: iter}
}

type Iterator struct {
	*pebble.Iterator
}

func (i *Iterator) Seek(key []byte) bool {
	return i.Iterator.SeekGE(key)
}

func (i *Iterator) Release() {
	i.Iterator.Close()
}
