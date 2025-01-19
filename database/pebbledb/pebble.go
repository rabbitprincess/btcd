// Copyright (c) 2015-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package pebbledb

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/btcsuite/btcd/database"
	"github.com/btcsuite/btcd/wire"
	"github.com/cockroachdb/pebble"
)

// openDB opens the database at the provided path.  database.ErrDbDoesNotExist
// is returned if the database doesn't exist and the create flag is not set.
func openDB(dbPath string, network wire.BitcoinNet, create bool) (database.DB, error) {
	// Error if the database doesn't exist and the create flag is not set.
	metadataDbPath := filepath.Join(dbPath, metadataDbName)
	dbExists := fileExists(metadataDbPath)
	if !create && !dbExists {
		str := fmt.Sprintf("database %q does not exist", metadataDbPath)
		return nil, makeDbErr(database.ErrDbDoesNotExist, str, nil)
	}

	// Ensure the full path to the database exists.
	if !dbExists {
		// The error can be ignored here since the call to
		// pebble.Open will fail if the directory couldn't be
		// created.
		_ = os.MkdirAll(dbPath, 0700)
	}

	// Open the metadata database (will create it if needed).
	opts := &pebble.Options{
		ErrorIfExists: create, // Fail if the database exists and create is true
	}
	ldb, err := pebble.Open(metadataDbPath, opts)
	if err != nil {
		return nil, convertErr(err.Error(), err)
	}

	// Create the block store which includes scanning the existing flat
	// block files to find what the current write cursor position is
	// according to the data that is actually on disk.  Also create the
	// database cache which wraps the underlying PebbleDB database to provide
	// write caching.
	store, err := newBlockStore(dbPath, network)
	if err != nil {
		return nil, convertErr(err.Error(), err)
	}
	cache := newDbCache(ldb, store, defaultCacheSize, defaultFlushSecs)
	pdb := &db{store: store, cache: cache}

	// Perform any reconciliation needed between the block and metadata as
	// well as database initialization, if needed.
	return reconcileDB(pdb, create)
}

// initDB creates the initial buckets and values used by the package. This is
// mainly in a separate function for testing purposes.
func initDB(db *pebble.DB) error {
	// The starting block file write cursor location is file num 0, offset 0.
	batch := db.NewBatch()
	defer batch.Close()

	// Insert the starting block file write cursor location.
	err := batch.Set(bucketizedKey(metadataBucketID, writeLocKeyName),
		serializeWriteRow(0, 0), pebble.NoSync)
	if err != nil {
		return fmt.Errorf("failed to set writeLocKeyName: %w", err)
	}

	// Create block index bucket and set the current bucket id.
	//
	// NOTE: Since buckets are virtualized through the use of prefixes,
	// there is no need to store the bucket index data for the metadata
	// bucket in the database. However, the first bucket ID to use does
	// need to account for it to ensure there are no key collisions.
	err = batch.Set(bucketIndexKey(metadataBucketID, blockIdxBucketName),
		blockIdxBucketID[:], pebble.NoSync)
	if err != nil {
		return fmt.Errorf("failed to set blockIdxBucketName: %w", err)
	}

	err = batch.Set(curBucketIDKeyName, blockIdxBucketID[:], pebble.NoSync)
	if err != nil {
		return fmt.Errorf("failed to set curBucketIDKeyName: %w", err)
	}

	// Apply the batch write.
	if err := batch.Commit(pebble.Sync); err != nil {
		str := fmt.Sprintf("failed to initialize metadata database: %v", err)
		return convertErr(str, err)
	}

	return nil
}
