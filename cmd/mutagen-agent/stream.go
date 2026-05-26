package main

import (
	"io"
)

// stdioStream is an io.ReadWriteCloser that uses standard input/output.
type stdioStream struct {
	io.Reader
	io.Writer
}

// newStdioStream creates a new stream that uses standard input/output.
func newStdioStream() io.ReadWriteCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadWriteCloser)
}

// Close implements io.Closer.Close.
func (s *stdioStream) Close() error {
	_ = "STUB: not implemented"
	// HACK: We can't really close standard input/output, because on many
	// platforms these can't be unblocked on reads and writes, and they'll
	// actually block the call to Close. In the case of the agent, where we're
	// just running an endpoint (which will be on the path to termination by the
	// time this method is invoked), this is fine.
	return nil
}
