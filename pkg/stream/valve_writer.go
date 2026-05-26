package stream

import (
	"io"
	"sync"
)

// ValveWriter is an io.Writer that wraps another io.Writer and forwards writes
// to it until the ValveWriter's internal valve is shut, after which writes will
// continue to succeed but not actually be written to the underlying writer.
type ValveWriter struct {
	// writerLock serializes access to the underlying writer.
	writerLock sync.Mutex
	// writer is the underlying writer.
	writer io.Writer
}

// NewValveWriter creates a new ValveWriter instance using the specified writer.
// The writer may be nil, in which case the writer will start pre-shut.
func NewValveWriter(writer io.Writer) *ValveWriter { _ = "STUB: not implemented"; return nil }

// Write implements io.Writer.Write.
func (w *ValveWriter) Write(buffer []byte) (int, error) {
	_ = "STUB: not implemented"
	// Lock the writer and defer its release.
	return 0, nil
}

// If there's no writer, then just pretend that we wrote all of the data.

// Otherwise write to the underlying writer.

// Shut closes the valve and prevents future writes to the underlying writer. It
// is safe to call Shut concurrently with Write, but doing so will not preempt
// or unblock pending calls to Write. Calling Shut will release the reference to
// the underlying writer.
func (w *ValveWriter) Shut() {
	_ = "STUB: not implemented"
	// Lock the writer and defer its release.
	return
}

// Nil out the writer to stop any future writes to it.
