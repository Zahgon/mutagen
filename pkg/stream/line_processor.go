package stream

import (
	"errors"
)

const (
	// defaultLineProcessorMaximumBufferSize is the default maximum buffer size
	// for LineProcessor.
	defaultLineProcessorMaximumBufferSize = 64 * 1024
)

// ErrMaximumBufferSizeExceeded is returned when a write would exceed the
// maximum internal buffer size for a writer.
var ErrMaximumBufferSizeExceeded = errors.New("maximum buffer size exceed")

// trimCarriageReturn trims any single trailing carriage return from the end of
// a byte slice.
func trimCarriageReturn(buffer []byte) []byte { _ = "STUB: not implemented"; return nil }

// LineProcessor is an io.Writer that splits its input stream into lines and
// writes those lines to a callback function. Line splits are performed on any
// instance of '\n' or '\r\n', with the split character(s) removed from the
// callback value.
type LineProcessor struct {
	// Callback is the line processing callback.
	Callback func(string)
	// MaximumBufferSize is the maximum allowed internal buffer size. If writes
	// to the writer exceed this size without incorporating a newline, then an
	// error will be raised. A value of 0 causes the writer to use a reasonable
	// default. A negative value indicates no limit.
	MaximumBufferSize int
	// buffer is any incomplete line fragment left over from a previous write.
	buffer []byte
}

// Write implements io.Writer.Write.
func (p *LineProcessor) Write(data []byte) (int, error) {
	_ = "STUB: not implemented"
	// Ensure that storing the data won't exceed buffer size limits.
	// TODO: We could truncate data here if any capacity remains. A partial
	// fragment could (in theory) contain a newline that would allow the buffer
	// to be cleared out, though it's hard to imagine such an optimization is
	// critical given the relatively large default maximum buffer size and the
	// typical line size of most newline-delimited data.
	return 0, nil
}

// Append the data to our internal buffer.

// Process all lines in the buffer and track the number of processed bytes.

// Find the index of the next newline character.

// Process the line.

// Update the number of bytes that we've processed.

// Update the remaining slice.

// If we managed to process bytes, then truncate our internal buffer.

// Compute the number of leftover bytes.

// If there are leftover bytes, then shift them to the front of the
// buffer.

// Truncate the buffer.

// Done.
