package stream

import (
	"io"
	"sync"
)

// concurrentWriter is an io.Writer that serializes calls to Write.
type concurrentWriter struct {
	// lock serializes operations on the writer.
	lock sync.Mutex
	// writer is the underlying writer.
	writer io.Writer
}

// NewConcurrentWriter creates a new writer that serializes operations on the
// underlying writer.
func NewConcurrentWriter(writer io.Writer) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

// Write implements io.Writer.Write.
func (w *concurrentWriter) Write(buffer []byte) (int, error) {
	_ = "STUB: not implemented"
	// Lock the writer and defer its release.
	return 0, nil
}

// Perform the write.
