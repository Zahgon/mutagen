package watching

import (
	"context"
	"sync"

	"github.com/mutagen-io/mutagen/pkg/container/lru"
	"github.com/mutagen-io/mutagen/pkg/filesystem/watching/internal/third_party/notify"
)

const (
	// NonRecursiveWatchingSupported indicates whether or not the current
	// platform supports native non-recursive watching.
	NonRecursiveWatchingSupported = true

	// inotifyChannelCapacity is the capacity to use for the internal inotify
	// events channel.
	inotifyChannelCapacity = 50
	// inotifyDefaultMaximumWatches is the default maximum number of inotify
	// watches that will be allowed to exist per-watcher.
	inotifyDefaultMaximumWatches = 50
)

// nonRecursiveWatcher implements NonRecursiveWatcher using inotify, with paths
// evicted on an LRU-basis.
type nonRecursiveWatcher struct {
	// watch is the underlying inotify-based watcher.
	watch notify.Watcher
	// evictor performs LRU-based watch eviction.
	evictor *lru.Cache[string, int]
	// events is the event delivery channel.
	events chan string
	// errors is the error delivery channel.
	errors chan error
	// cancel is the run loop cancellation function.
	cancel context.CancelFunc
	// done is the run loop completion signaling mechanism.
	done sync.WaitGroup
}

// NewNonRecursiveWatcher creates a new inotify-based non-recursive watcher.
func NewNonRecursiveWatcher() (NonRecursiveWatcher, error) {
	_ = "STUB: not implemented"
	// Create the raw event channel.
	return *new(NonRecursiveWatcher), nil
}

// Create a context to regulate the watcher's run loop.

// Create the watcher. The LRU evictor ensures that we don't
// exceed the maximum number of inotify watches by unwatching the
// least recently used paths when the cache overflows.

// Track run loop termination.

// Start the run loop.

// Success.

// run implements the event processing run loop for nonRecursiveWatcher.
func (w *nonRecursiveWatcher) run(ctx context.Context, rawEvents <-chan notify.EventInfo) error {
	_ = "STUB: not implemented"
	// Loop indefinitely, polling for cancellation and events.
	return nil
}

// Ensure that the event channel wasn't closed.

// Transmit the path.

// Watch implements NonRecursiveWatcher.Watch.
func (w *nonRecursiveWatcher) Watch(path string) {
	_ = "STUB: not implemented"
	// Attempt to evict the path if already watched, that way we can establish a
	// clean watch and make the path the most-recently-added record. If the path
	// isn't currently watched, then this is a no-op.
	return
}

// Start the watch. If it fails due to a non-existence error, then we can
// just avoid adding it. If it fails for any other reason, then report the
// error via the errors channel.

// Unwatch implements NonRecursiveWatcher.Unwatch.
func (w *nonRecursiveWatcher) Unwatch(path string) {
	_ = "STUB: not implemented"
	// Remove the watch via eviction. This is a no-op if the path isn't watched.
	return
}

// Events implements NonRecursiveWatcher.Events.
func (w *nonRecursiveWatcher) Events() <-chan string {
	_ = "STUB: not implemented"

	// Errors implements NonRecursiveWatcher.Errors.
	return nil
}

func (w *nonRecursiveWatcher) Errors() <-chan error {
	_ = "STUB: not implemented"

	// Terminate implements NonRecursiveWatcher.Terminate.
	return nil
}

func (w *nonRecursiveWatcher) Terminate() error {
	_ = "STUB: not implemented"
	// Signal termination.
	return nil
}

// Wait for the run loop to exit.

// Terminate the underlying watcher.
