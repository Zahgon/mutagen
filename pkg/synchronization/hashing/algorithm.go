package hashing

import (
	"hash"
)

// IsDefault indicates whether or not the algorithm is
// Algorithm_AlgorithmDefault.
func (a Algorithm) IsDefault() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (a Algorithm) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (a *Algorithm) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a hashing algorithm.

// Success.

// AlgorithmSupportStatus encodes support status for a hashing algorithm.
type AlgorithmSupportStatus uint8

const (
	// AlgorithmSupportStatusUnsupported indicates that an algorithm is
	// completely unsupported.
	AlgorithmSupportStatusUnsupported AlgorithmSupportStatus = iota
	// AlgorithmSupportStatusRequiresLicense indicates that an algorithm is
	// supported but requires a (currently absent) Mutagen Pro license.
	AlgorithmSupportStatusRequiresLicense
	// AlgorithmSupportStatusSupported indicates that an algorithm is fully
	// supported, either due to being supported universally in Mutagen or due to
	// the presence of a Mutagen Pro license.
	AlgorithmSupportStatusSupported
)

// SupportStatus returns the support status for a particular algorithm.
func (a Algorithm) SupportStatus() AlgorithmSupportStatus {
	_ = "STUB: not implemented"
	return *new(AlgorithmSupportStatus)
}

// Description returns a human-readable description of a hashing algorithm.
func (a Algorithm) Description() string { _ = "STUB: not implemented"; return "" }

// Factory returns a constructor for the hashing algorithm. If invoked on a
// default or invalid Algorithm value, this method will panic.
func (a Algorithm) Factory() func() hash.Hash { _ = "STUB: not implemented"; return nil }
