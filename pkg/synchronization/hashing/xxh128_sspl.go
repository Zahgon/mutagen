//go:build mutagensspl

package hashing

import (
	"hash"
)

// xxh128SupportStatus returns XXH128 hashing support status.
func xxh128SupportStatus() AlgorithmSupportStatus {
	_ = "STUB: not implemented"
	return *new(AlgorithmSupportStatus)
}

// newXXH128Factory creates a new hasher factory for XXH128 hashers.
func newXXH128Factory() func() hash.Hash { _ = "STUB: not implemented"; return nil }
