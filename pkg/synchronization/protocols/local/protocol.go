package local

import (
	"context"

	"github.com/mutagen-io/mutagen/pkg/logging"
	"github.com/mutagen-io/mutagen/pkg/synchronization"
	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
)

// protocolHandler implements the synchronization.ProtocolHandler interface for
// connecting to local endpoints.
type protocolHandler struct{}

// Connect connects to a local endpoint.
func (h *protocolHandler) Connect(
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

// Create a local endpoint.

// Success.

func init() {
	// Register the local protocol handler with the synchronization package.
	synchronization.ProtocolHandlers[urlpkg.Protocol_Local] = &protocolHandler{}
}
