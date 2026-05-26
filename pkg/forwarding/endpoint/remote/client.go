package remote

import (
	"io"
	"net"

	"github.com/mutagen-io/mutagen/pkg/forwarding"
	"github.com/mutagen-io/mutagen/pkg/logging"
	"github.com/mutagen-io/mutagen/pkg/multiplexing"
)

// client is a client for a remote forwarding.Endpoint and implements
// forwarding.Endpoint itself.
type client struct {
	// logger is the underlying logger.
	logger *logging.Logger
	// transportErrors is the transport error channel.
	transportErrors <-chan error
	// multiplexer is the underlying multiplexer.
	multiplexer *multiplexing.Multiplexer
	// listener indicates whether or not the remote endpoint is operating as a
	// listener.
	listener bool
}

// NewEndpoint creates a new remote forwarding.Endpoint operating over the
// specified stream with the specified metadata. If this function fails, then
// the provided stream will be closed. Once the endpoint has been established,
// the underlying stream is owned by the endpoint and will be closed when the
// endpoint is shut down. The provided stream must unblock read and write
// operations when closed.
func NewEndpoint(
	logger *logging.Logger,
	stream io.ReadWriteCloser,
	version forwarding.Version,
	configuration *forwarding.Configuration,
	protocol string,
	address string,
	source bool,
) (forwarding.Endpoint, error) {
	_ = "STUB: not implemented"
	// Adapt the stream to serve as a multiplexer carrier. This will also give
	// us the buffering functionality we'll need for initialization.
	return *new(forwarding.Endpoint), nil
}

// Defer closure of the carrier in the event that initialization isn't
// successful. Otherwise, we'll rely on closure of the multiplexer to close
// the carrier.

// Create and send the initialization request.

// Receive the initialization response, ensure that it's valid, and check
// for initialization errors.

// Mark initialization as successful.

// Multiplex the carrier.

// Create a channel to monitor for transport errors and a Goroutine to
// populate it.

// Success.

// TransportErrors implements forwarding.Endpoint.TransportErrors.
func (c *client) TransportErrors() <-chan error { _ = "STUB: not implemented"; return nil }

// Open implements forwarding.Endpoint.Open.
func (c *client) Open() (net.Conn, error) { _ = "STUB: not implemented"; return *new(net.Conn), nil }

// Shutdown implements forwarding.Endpoint.Shutdown.
func (c *client) Shutdown() error { _ = "STUB: not implemented"; return nil }
