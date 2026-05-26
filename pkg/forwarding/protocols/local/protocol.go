package local

import (
	"context"

	"github.com/mutagen-io/mutagen/pkg/forwarding"
	"github.com/mutagen-io/mutagen/pkg/logging"
	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
)

// protocolHandler implements the forwarding.ProtocolHandler interface for
// connecting to local forwarding endpoints.
type protocolHandler struct{}

// Connect implements forwarding.ProtocolHandler.Connect.
func (p *protocolHandler) Connect(
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

// Handle creation based on mode.

func init() {
	// Register the local protocol handler with the forwarding package.
	forwarding.ProtocolHandlers[urlpkg.Protocol_Local] = &protocolHandler{}
}
