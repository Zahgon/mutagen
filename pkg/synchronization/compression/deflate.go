package compression

import (
	"io"

	"github.com/mutagen-io/mutagen/pkg/stream"
)

// compressDeflate implements compression for DEFLATE streams.
func compressDeflate(compressed io.Writer) stream.WriteFlushCloser {
	_ = "STUB: not implemented"
	// Create the compressor. We check for errors, but we don't include them as
	// part of the interface because they can only occur with an invalid
	// compressor configuration (which can't occur when we only use defaults).
	return *new(stream.WriteFlushCloser)
}

// Success.

// decompressDeflate implements decompression for DEFLATE streams.
func decompressDeflate(compressed io.Reader) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}
