package staging

import (
	"hash"
	"io"

	"github.com/mutagen-io/mutagen/pkg/synchronization/endpoint/local/staging/store"
)

// Stager is an implementation of local.stager that uses a content-addressable
// store to stage files.
type Stager struct {
	// store is the stager's underlying store.
	store *store.Store
}

// NewStager creates a new stager.
func NewStager(root string, hideRoot bool, maximumFileSize uint64, hasherFactory func() hash.Hash) *Stager {
	_ = "STUB: not implemented"
	return nil
}

// Initialize implements local.stager.Initialize.
func (s *Stager) Initialize() error { _ = "STUB: not implemented"; return nil }

// Contains implements local.stager.Contains.
func (s *Stager) Contains(path string, digest []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Sink implements rsync.Sinker.Sink.
func (s *Stager) Sink(path string) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

// Provide implements core.Provider.Provide.
func (s *Stager) Provide(path string, digest []byte) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Finalize implements local.stager.Finalize.
func (s *Stager) Finalize() error { _ = "STUB: not implemented"; return nil }

// Sink implements io.WriterCloser for Stager's Sink method.
type Sink struct {
	// path is the path associated with the sink.
	path string
	// storage is the underlying file storage.
	storage *store.Storage
}

// Writer implements io.Writer.Write.
func (s *Sink) Write(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close implements io.Closer.Close.
func (s *Sink) Close() error { _ = "STUB: not implemented"; return nil }
