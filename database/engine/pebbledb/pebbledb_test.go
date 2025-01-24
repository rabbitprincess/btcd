package pebbledb

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/btcsuite/btcd/database/engine"
	"github.com/stretchr/testify/require"
)

func TestSuitePebbleDB(t *testing.T) {
	dbPath := filepath.Join(os.TempDir(), "pebbledb-testsuite")
	defer func() {
		require.NoError(t, os.RemoveAll(dbPath))
	}()

	engine.TestSuiteEngine(t, func() engine.Engine {
		leveldb, err := NewDB(dbPath, true, 0, 0)
		require.NoErrorf(t, err, "failed to create leveldb")
		return leveldb
	})
}
