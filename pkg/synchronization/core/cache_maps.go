package core

// emptyByteLookupMap implements byteLookupMap for empty caches.
type emptyByteLookupMap struct{}

// length returns the length of the map.
func (m *emptyByteLookupMap) length() int {
	_ = "STUB: not implemented"

	// insert adds a key-value pair to the map.
	return 0
}

func (m *emptyByteLookupMap) insert(_ []byte, _ string) {
	_ = "STUB: not implemented"

	// find looks for a key in the map, returning the associated value (defaulting
	// to an empty string if the key was not present) and whether or not the key was
	// found.
	return
}

func (m *emptyByteLookupMap) find(_ []byte) (string, bool) {
	_ = "STUB: not implemented"

	// byteLookupMap20 implements byteLookupMap for 20-byte digests.
	return "", false
}

type byteLookupMap20 map[[20]byte]string

// length returns the length of the map.
func (m byteLookupMap20) length() int {
	_ = "STUB: not implemented"

	// insert adds a key-value pair to the map.
	return 0
}

func (m byteLookupMap20) insert(k []byte, v string) { _ = "STUB: not implemented"; return }

// find looks for a key in the map, returning the associated value (defaulting
// to an empty string if the key was not present) and whether or not the key was
// found.
func (m byteLookupMap20) find(k []byte) (string, bool) { _ = "STUB: not implemented"; return "", false }

// byteLookupMap32 implements byteLookupMap for 32-byte digests.
type byteLookupMap32 map[[32]byte]string

// length returns the length of the map.
func (m byteLookupMap32) length() int {
	_ = "STUB: not implemented"

	// insert adds a key-value pair to the map.
	return 0
}

func (m byteLookupMap32) insert(k []byte, v string) { _ = "STUB: not implemented"; return }

// find looks for a key in the map, returning the associated value (defaulting
// to an empty string if the key was not present) and whether or not the key was
// found.
func (m byteLookupMap32) find(k []byte) (string, bool) { _ = "STUB: not implemented"; return "", false }

// byteLookupMap16 implements byteLookupMap for 16-byte digests.
type byteLookupMap16 map[[16]byte]string

// length returns the length of the map.
func (m byteLookupMap16) length() int {
	_ = "STUB: not implemented"

	// insert adds a key-value pair to the map.
	return 0
}

func (m byteLookupMap16) insert(k []byte, v string) { _ = "STUB: not implemented"; return }

// find looks for a key in the map, returning the associated value (defaulting
// to an empty string if the key was not present) and whether or not the key was
// found.
func (m byteLookupMap16) find(k []byte) (string, bool) { _ = "STUB: not implemented"; return "", false }
