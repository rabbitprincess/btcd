package leveldb

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/btcsuite/btcd/database/engine"
	"github.com/stretchr/testify/require"
)

func TestSuiteLevelDB(t *testing.T) {
	dbPath := filepath.Join(os.TempDir(), "leveldb-testsuite")
	defer func() {
		require.NoError(t, os.RemoveAll(dbPath))
	}()

	engine.TestSuiteEngine(t, func() engine.Engine {
		leveldb, err := NewDB(dbPath, true)
		require.NoErrorf(t, err, "failed to create leveldb")
		return leveldb
	})
}
