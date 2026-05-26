package netpipe

import (
	"context"

	"github.com/mutagen-io/mutagen/pkg/logging"
	"github.com/mutagen-io/mutagen/pkg/synchronization"
	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
)

// synchronizationProtocolHandler implements the synchronization.ProtocolHandler
// interface for connecting to "remote" endpoints that actually exist in memory
// via an in-memory pipe.
type synchronizationProtocolHandler struct{}

// waitingSynchronizationEndpoint wraps and implements synchronization.Endpoint,
// but adds a waiting function that's invoked after invoking Shutdown on the
// underlying endpoint. It is necessary to ensure full endpoint shutdown in
// tests, where open file descriptors or handles can prevent temporary directory
// removal.
type waitingSynchronizationEndpoint struct {
	// Endpoint is the underlying endpoint.
	synchronization.Endpoint
	// wait is an arbitrary waiting function.
	wait func()
}

// Shutdown implements synchronization.Endpoint.Shutdown.
func (w *waitingSynchronizationEndpoint) Shutdown() error {
	_ = "STUB: not implemented"
	// Shutdown on the underlying endpoint.
	return nil
}

// Perform the wait operation.

// Done.

// Dial starts an endpoint server in a background Goroutine and creates an
// endpoint client connected to the server via an in-memory connection.
func (h *synchronizationProtocolHandler) Connect(
	_ context.Context,
	logger *logging.Logger,
	url *urlpkg.URL,
	prompter string,
	session string,
	version synchronization.Version,
	configuration *synchronization.Configuration,
	alpha bool,
) (synchronization.Endpoint, error) {
	_ = "STUB: not implemented"
	// Verify that the URL is of the correct kind and protocol.
	return *new(synchronization.Endpoint), nil
}

// Create an in-memory network connection.

// Serve the endpoint in a background Goroutine. This will terminate once
// the client connection is closed. We monitor for its termination so that
// we can block on it in our endpoint wrapper.

// Create a client for this endpoint.

// Wrap the client so that it blocks on the full shutdown of the remote
// endpoint after closing the connection. This is necessary for testing,
// where we need to ensure that all file descriptors or handles point to
// temporary test directories are closed before attempting to remove those
// directories. This is not necessary for other remote protocols in normal
// usage (because we don't have the same constraints) or in testing (because
// the underlying connection closure waits for agent process termination).

// Success.

func init() {
	// Register the netpipe protocol handler with the synchronization package.
	synchronization.ProtocolHandlers[Protocol_Netpipe] = &synchronizationProtocolHandler{}
}
