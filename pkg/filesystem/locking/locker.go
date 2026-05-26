package locking

import (
	"os"
)

// Locker provides file locking facilities.
type Locker struct {
	// file is the underlying file object that's locked.
	file *os.File
	// held indicates whether or not the lock is currently held.
	held bool
}

// NewLocker attempts to create a lock with the file at the specified path,
// creating the file if necessary. The lock is returned in an unlocked state.
func NewLocker(path string, permissions os.FileMode) (*Locker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Held returns whether or not the lock is currently held.
func (l *Locker) Held() bool {
	_ = "STUB: not implemented"

	// Read implements io.Reader.Read on the underlying file, but errors if the lock
	// is not currently held.
	return false
}

func (l *Locker) Read(buffer []byte) (int, error) {
	_ = "STUB: not implemented"
	// Verify that the lock is held.
	return 0, nil
}

// Perform the read.

// Write implements io.Writer.Write on the underlying file, but errors if the
// lock is not currently held.
func (l *Locker) Write(buffer []byte) (int, error) {
	_ = "STUB: not implemented"
	// Verify that the lock is held.
	return 0, nil
}

// Perform the write.

// Truncate implements file truncation for the underlying file, but errors if
// the lock is not currently held.
func (l *Locker) Truncate(size int64) error {
	_ = "STUB: not implemented"
	// Verify that the lock is held.
	return nil
}

// Perform the truncation.

// Close closes the file underlying the locker. This will release any lock held
// on the file and disable future locking. On POSIX platforms, this also
// releases other locks held on the same file.
func (l *Locker) Close() error { _ = "STUB: not implemented"; return nil }
