//go:build !linux

package watching

const (
	// NonRecursiveWatchingSupported indicates whether or not the current
	// platform supports native non-recursive watching.
	NonRecursiveWatchingSupported = false
)

// NewNonRecursiveWatcher creates a new non-recursive watcher on platforms that
// support native non-recursive watching. This platform does not support
// recursive watching and this function will panic if called.
func NewNonRecursiveWatcher() (NonRecursiveWatcher, error) {
	_ = "STUB: not implemented"
	return *new(NonRecursiveWatcher), nil
}
