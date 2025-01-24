package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuiteEngine(t *testing.T, new func() Engine) {
	t.Run("TransactionSnapshot", func(t *testing.T) {
		engine := new()
		defer engine.Close()

		// Create new transaction
		tx, err := engine.Transaction()
		require.NoErrorf(t, err, "failed to create transaction")

		// Put some data into the transaction
		key := []byte("key1")
		value := []byte("value1")
		err = tx.Put(key, value)
		require.NoErrorf(t, err, "failed to put data into transaction")

		// Create a snapshot and find empty data
		snapshot, err := engine.Snapshot()
		require.NoErrorf(t, err, "failed to create snapshot")

		has, err := snapshot.Has(key)
		require.NoErrorf(t, err, "failed to check if key exists in snapshot")
		require.Falsef(t, has, "expected key to not exist in snapshot")

		gotValue, err := snapshot.Get(key)
		require.Errorf(t, err, "expected to get error when getting value from snapshot")
		require.Nil(t, gotValue, "expected to get nil value from snapshot")

		snapshot.Release()

		// Commit the transaction
		err = tx.Commit()
		require.NoErrorf(t, err, "failed to commit transaction")

		// Create a snapshot and verify the data
		snapshot, err = engine.Snapshot()
		require.NoErrorf(t, err, "failed to create snapshot")

		gotValue, err = snapshot.Get(key)
		require.NoErrorf(t, err, "failed to get value from snapshot")
		require.Equalf(t, gotValue, value, "snapshot value mismatch")
		snapshot.Release()
	})

	t.Run("TransactionIterator", func(t *testing.T) {

	})

	t.Run("DbClose", func(t *testing.T) {
		// engine := new()
		// err := engine.Close()
		// require.NoErrorf(t, err, "failed to close engine")

		// // Ensure that the engine is closed
		// err = engine.Close()
		// require.Errorf(t, err, "expected to get error when closing closed engine")

	})

}
