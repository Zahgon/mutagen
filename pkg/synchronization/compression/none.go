package compression

import (
	"io"

	"github.com/mutagen-io/mutagen/pkg/stream"
)

// noneCompressor implements stream.WriteFlushCloser for uncompressed streams.
type noneCompressor struct {
	// Writer is the underlying writer.
	io.Writer
}

// Flush implements stream.Flusher.Flush.
func (c *noneCompressor) Flush() error {
	_ = "STUB: not implemented"

	// Close implements io.Closer.Close.
	return nil
}

func (c *noneCompressor) Close() error {
	_ = "STUB: not implemented"

	// compressNone implements no-op compression for uncompressed streams.
	return nil
}

func compressNone(compressed io.Writer) stream.WriteFlushCloser {
	_ = "STUB: not implemented"
	return *new(stream.WriteFlushCloser)
}

// decompressNone implements no-op decompression for uncompressed streams.
func decompressNone(compressed io.Reader) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}
