package remote

import (
	"io"

	"github.com/mutagen-io/mutagen/pkg/forwarding"
	"github.com/mutagen-io/mutagen/pkg/logging"
)

// initializeEndpoint initializes the underlying endpoint based on the provided
// initialization request.
func initializeEndpoint(logger *logging.Logger, request *InitializeForwardingRequest) (forwarding.Endpoint, error) {
	_ = "STUB: not implemented"
	// If this is a Unix domain socket endpoint, perform normalization on the
	// socket path.
	return *new(forwarding.Endpoint), nil
}

// Create the underlying endpoint based on the initialization parameters.

// ServeEndpoint creates and serves a remote endpoint on the specified stream.
// It enforces that the provided stream is closed by the time this function
// returns, regardless of failure. The provided stream must unblock read and
// write operations when closed.
func ServeEndpoint(logger *logging.Logger, stream io.ReadWriteCloser) error {
	_ = "STUB: not implemented"
	// Adapt the connection to serve as a multiplexer carrier. This will also
	// give us the buffering functionality we'll need for initialization.
	return nil
}

// Defer closure of the carrier in the event that initialization isn't
// successful. Otherwise, we'll rely on closure of the multiplexer to close
// the carrier.

// Receive the initialization request, ensure that it's valid, and perform
// initialization.

// Send the initialization response, indicating any initialization error
// that occurred.

// If initialization failed, then bail.

// Multiplex the carrier and defer closure of the multiplexer.

// Start a Goroutine couple the lifetime of the underlying endpoint to the
// lifetime of the multiplexer. This will cause the underlying endpoint to
// be shut down if this function returns or if the multiplexer shuts down
// due to remote closure or an internal error. This is particularly
// important for preempting local accept operations.

// Receive and forward connections indefinitely.

// Receive the next incoming connection. If this fails, then we should
// terminate serving because either the local listener has failed or the
// multiplexer has failed.

// Open the corresponding outgoing connection. If the multiplexer fails,
// then we should terminate serving. If local dialing fails, then we can
// just close the incoming connection to indicate dialing failure.

// Perform forwarding.
