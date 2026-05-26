package core

// EnsureValid ensures that Cache's invariants are respected.
func (c *Cache) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil cache is considered valid (though obviously that requires using
	// the GetEntries accessor).
	return nil
}

// Technically we could validate each path, but that's error prone,
// expensive, and not really needed for memory safety. Also note that an
// empty path is valid when the synchronization root is a file.

// Nil cache entries are invalid.

// Success.

// Equal determines whether or not another cache is equal to this one. It is
// designed specifically for tests, though it is exported so that it can be used
// by scan_bench.
func (c *Cache) Equal(other *Cache) bool {
	_ = "STUB: not implemented"
	// Verify non-nilness. We don't consider nil caches valid, so we don't
	// consider them equal.
	return false
}

// Handle equivalence fast paths.

// Check lengths.

// Check contents.

// Extract corresponding content.

// Watch for nil values as a sanity check.

// Verify equivalence

// Success.

// byteLookupMap is the interface implemented by all byteLookupMap types.
type byteLookupMap interface {
	// length returns the length of the map.
	length() int
	// insert adds a key-value pair to the map.
	insert(k []byte, v string)
	// find looks for a key in the map, returning the associated value
	// (defaulting to an empty string if the key was not present) and whether or
	// not the key was found.
	find(k []byte) (string, bool)
}

// ReverseLookupMap provides facilities for doing reverse lookups to avoid
// expensive staging operations in the case of renames and copies.
type ReverseLookupMap struct {
	// lookupMap is the underlying map.
	lookupMap byteLookupMap
}

// Length returns the number of entries in the map.
func (m *ReverseLookupMap) Length() int { _ = "STUB: not implemented"; return 0 }

// Lookup attempts a lookup in the map.
func (m *ReverseLookupMap) Lookup(digest []byte) (string, bool) {
	_ = "STUB: not implemented"
	return "", false

	// GenerateReverseLookupMap creates a reverse lookup map from a cache.
}

func (c *Cache) GenerateReverseLookupMap() (*ReverseLookupMap, error) {
	_ = "STUB: not implemented"
	// Create a placeholder for the map that we're going to initialize.
	return nil, nil
}

// Track the digest size and ensure it's consistent.

// Loop over entries.

// Compute and validate the digest size and allocate the map.

// Insert the entry.

// If there are no entries, then we'll still need a lookup map.

// Success.
