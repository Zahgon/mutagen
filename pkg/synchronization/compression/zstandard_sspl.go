//go:build mutagensspl

package compression

import (
	"io"

	"github.com/mutagen-io/mutagen/pkg/stream"
)

// zstandardSupportStatus returns Zstandard compression support status.
func zstandardSupportStatus() AlgorithmSupportStatus {
	_ = "STUB: not implemented"
	return *new(AlgorithmSupportStatus)
}

// compressZstandard implements compression for Zstandard streams.
func compressZstandard(compressed io.Writer) stream.WriteFlushCloser {
	_ = "STUB: not implemented"
	return *new(stream.WriteFlushCloser)
}

// decompressZstandard implements decompression for Zstandard streams.
func decompressZstandard(compressed io.Reader) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}
