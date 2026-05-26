//go:build !windows

package filesystem

// file is the readable file implementation used on POSIX systems. We avoid
// using os.File because its construction and operation can be expensive, its
// internals are complex, and it doesn't add any benefit for regular on-disk
// files (since polling and asynchronous I/O aren't currently supported).
type file int

// Read implements io.Reader.Read.
func (f file) Read(buffer []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Seek implements io.Seeker.Seek.
func (f file) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Close implements io.Closer.Close.
func (f file) Close() error { _ = "STUB: not implemented"; return nil }
