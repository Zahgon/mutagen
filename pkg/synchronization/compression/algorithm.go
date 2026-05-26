package compression

import (
	"io"

	"github.com/mutagen-io/mutagen/pkg/stream"
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

// Convert to a compression algorithm.

// Success.

// AlgorithmSupportStatus encodes support status for a compression algorithm.
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

// Description returns a human-readable description of a compression algorithm.
func (a Algorithm) Description() string { _ = "STUB: not implemented"; return "" }

// Compress creates a compressor that writes compressed output to the specified
// stream using the compression algorithm. If invoked on a default or invalid
// Algorithm value, this method will panic. The Flush and Close methods on the
// resulting compressor only operate on the compressor - they have no effect on
// the compressed stream itself. The compressor should be flushed and/or closed
// before the underlying stream.
func (a Algorithm) Compress(compressed io.Writer) stream.WriteFlushCloser {
	_ = "STUB: not implemented"
	return *new(stream.WriteFlushCloser)
}

// Decompress creates a decompressor that reads compressed input from the
// specified stream using the compression algorithm. If invoked on a default or
// invalid Algorithm value, this method will panic. The Close method on the
// resulting decompressor releases decompression resources - it has no effect on
// the compressed stream itself. The decompressor should be closed after the
// underlying stream.
func (a Algorithm) Decompress(compressed io.Reader) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}
