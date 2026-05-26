package netpipe

import (
	"context"

	"github.com/mutagen-io/mutagen/pkg/forwarding"
	"github.com/mutagen-io/mutagen/pkg/logging"
	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
)

// forwardingProtocolHandler implements the forwarding.ProtocolHandler interface
// for connecting to "remote" endpoints that actually exist in memory via an
// in-memory pipe.
type forwardingProtocolHandler struct{}

// Dial starts an endpoint server in a background Goroutine and creates an
// endpoint client connected to the server via an in-memory connection.
func (h *forwardingProtocolHandler) Connect(
	_ context.Context,
	logger *logging.Logger,
	url *urlpkg.URL,
	prompter string,
	session string,
	version forwarding.Version,
	configuration *forwarding.Configuration,
	source bool,
) (forwarding.Endpoint, error) {
	_ = "STUB: not implemented"
	// Verify that the URL is of the correct kind and protocol.
	return *new(forwarding.Endpoint), nil
}

// Parse the target specification from the URL's Path component.

// Create an in-memory network connection.

// Server the endpoint in a background Goroutine. This will terminate once
// the client connection is closed.

// Create a client for this endpoint.

// Success.

func init() {
	// Register the netpipe protocol handler with the forwarding package.
	forwarding.ProtocolHandlers[Protocol_Netpipe] = &forwardingProtocolHandler{}
}
