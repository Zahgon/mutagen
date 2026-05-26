package local

import (
	"context"
	"net"

	"github.com/mutagen-io/mutagen/pkg/forwarding"
	"github.com/mutagen-io/mutagen/pkg/logging"
)

// dialerEndpoint implements forwarding.Endpoint for dialer endpoints.
type dialerEndpoint struct {
	// logger is the underlying logger.
	logger *logging.Logger
	// dialingCtx limits the duration of dialing operations.
	dialingCtx context.Context
	// dialingCancel cancels the dialing context.
	dialingCancel context.CancelFunc
	// dialer is the dialer used for TCP and Unix domain socket dialing.
	dialer *net.Dialer
	// protocol is the protocol to use for dialing.
	protocol string
	// address is the address to use for dialing.
	address string
}

// NewDialerEndpoint creates a new forwarding.Endpoint that acts as a dialer.
func NewDialerEndpoint(
	logger *logging.Logger,
	version forwarding.Version,
	configuration *forwarding.Configuration,
	protocol string,
	address string,
) (forwarding.Endpoint, error) {
	_ = "STUB: not implemented"
	// Create a cancellable context that we can use to regulate connections.
	return *new(forwarding.Endpoint), nil
}

// Create the dialer (unless we're targeting a Windows named pipe).

// Create the endpoint.

// TransportErrors implements forwarding.Endpoint.TransportErrors.
func (e *dialerEndpoint) TransportErrors() <-chan error {
	_ = "STUB: not implemented"

	// Open implements forwarding.Endpoint.Open.
	return nil
}

func (e *dialerEndpoint) Open() (net.Conn, error) {
	_ = "STUB: not implemented"
	// If we're dealing with a Windows named pipe target, then perform dialing
	// using the platform-specific dialing function.
	return *new(net.Conn), nil
}

// For all other protocols (i.e. TCP and Unix domain sockets), use the
// standard dialer.

// Shutdown implements forwarding.Endpoint.Shutdown.
func (e *dialerEndpoint) Shutdown() error {
	_ = "STUB: not implemented"
	// Cancel the dialing context to unblock any dialing operations.
	return nil
}

// Success.
