//go:build !(darwin && cgo) && !(linux && mutagensspl && mutagenfanotify) && !windows

package watching

// RecursiveWatchingSupported indicates whether or not the current platform
// supports native recursive watching.
const RecursiveWatchingSupported = false

// NewRecursiveWatcher creates a new recursive watcher on platforms that support
// native recursive watching. This platform does not support recursive watching
// and this function will panic if called.
func NewRecursiveWatcher(_ string) (RecursiveWatcher, error) {
	_ = "STUB: not implemented"
	return *new(RecursiveWatcher), nil
}
