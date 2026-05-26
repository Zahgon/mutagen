package stream

import (
	"hash"
	"io"
)

// hashedWriter is the io.Writer implementation underlying NewHashedWriter.
type hashedWriter struct {
	// writer is the underlying writer.
	writer io.Writer
	// hasher is the associated hash function.
	hasher hash.Hash
}

// NewHashedWriter creates a new io.Writer that attaches a hash function to an
// existing writer, ensuring that the hash processes all bytes that are
// successfully written to the associated writer.
func NewHashedWriter(writer io.Writer, hasher hash.Hash) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

// Write implements io.Writer.Write.
func (w *hashedWriter) Write(data []byte) (int, error) {
	_ = "STUB: not implemented"
	// Write to the underlying writer.
	return 0, nil
}

// Write the corresponding bytes to the hasher. This write can't fail, so we
// can safely assume that all provided bytes are processed.

// Done.
