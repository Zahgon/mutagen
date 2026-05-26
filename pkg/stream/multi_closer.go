package stream

import (
	"io"
)

// multiCloser is the io.Closer implementation underlying NewMultiCloser.
type multiCloser struct {
	// closers are the underlying closers.
	closers []io.Closer
}

// NewMultiCloser creates a single closer that closes multiple underlying
// closers. The closers are closed in the order specified, and thus higher
// layers should be specified before lower. All closers will be closed, but only
// the first error encountered will be returned.
func NewMultiCloser(closers ...io.Closer) io.Closer {
	_ = "STUB: not implemented"
	return *new(io.Closer)
}

// Close implements io.Closer.Close.
func (c *multiCloser) Close() error { _ = "STUB: not implemented"; return nil }
