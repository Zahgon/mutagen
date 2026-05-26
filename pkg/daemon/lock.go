package daemon

import (
	"github.com/mutagen-io/mutagen/pkg/filesystem/locking"
)

// Lock represents the global daemon lock. It is held by a single daemon
// instance at a time.
type Lock struct {
	// locker is the underlying file locker.
	locker *locking.Locker
}

// AcquireLock attempts to acquire the global daemon lock.
func AcquireLock() (*Lock, error) {
	_ = "STUB: not implemented"
	// Compute the lock path.
	return nil, nil
}

// Create the locker and attempt to acquire the lock.

// Create the lock.

// Release releases the daemon lock.
func (l *Lock) Release() error {
	_ = "STUB: not implemented"
	// Release the lock.
	return nil
}

// Close the locker.

// Success.
