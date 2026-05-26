package stream

// multiFlusher is the Flusher implementation underlying NewMultiFlusher.
type multiFlusher struct {
	// flushers are the underlying flushers.
	flushers []Flusher
}

// NewMultiFlusher creates a single flusher that flushes multiple underlying
// flushers. The flushers are flushed in the order specified, and thus higher
// layers should be specified before lower. If an error occurs, then flushing
// halts and subsequent flushers are not flushed.
func NewMultiFlusher(flushers ...Flusher) Flusher { _ = "STUB: not implemented"; return *new(Flusher) }

// Flush implements Flusher.Flush.
func (f *multiFlusher) Flush() error { _ = "STUB: not implemented"; return nil }
