package stream

import (
	"io"
)

// cutoffWriter is the io.Writer implementation that underlies NewCutoffWriter.
type cutoffWriter struct {
	// writer is the underlying writer.
	writer io.Writer
	// cutoff is the number of bytes remaining.
	cutoff uint
}

// NewCutoffWriter creates a new io.Writer that wraps and forwards writes to the
// specified writer until a specified maximum number of bytes has been written,
// at which point the writing of any further bytes will succeed automatically,
// but not actually write to the underlying writer.
func NewCutoffWriter(writer io.Writer, cutoff uint) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

// Write implements io.Writer.Write.
func (w *cutoffWriter) Write(buffer []byte) (int, error) {
	_ = "STUB: not implemented"
	// If we've already hit the cutoff, then there's no action to take, we just
	// pretend like the full write succeeded.
	return 0, nil
}

// If the buffer length will fit within the cutoff, then just perform the
// write, update the cutoff, and return the result.

// Otherwise, perform a truncated write, update the cutoff, and return the
// count based on whether or not the truncated write failed. If the
// truncated write succeeded, then the cutoff is now zero and we pretend
// that the rest of the bytes were also written successfully.
