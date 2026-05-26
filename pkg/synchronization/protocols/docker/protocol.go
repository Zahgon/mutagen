package docker

import (
	"context"
	"io"

	"github.com/mutagen-io/mutagen/pkg/logging"
	"github.com/mutagen-io/mutagen/pkg/synchronization"
	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
)

// protocolHandler implements the synchronization.ProtocolHandler interface for
// connecting to remote endpoints inside Docker containers. It uses the agent
// infrastructure over a Docker transport.
type protocolHandler struct{}

// dialResult provides asynchronous agent dialing results.
type dialResult struct {
	// stream is the stream returned by agent dialing.
	stream io.ReadWriteCloser
	// error is the error returned by agent dialing.
	error error
}

// Connect connects to a Docker endpoint.
func (h *protocolHandler) Connect(
	ctx context.Context,
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

// Create a Docker agent transport.

// Create a channel to deliver the dialing result.

// Perform dialing in a background Goroutine so that we can monitor for
// cancellation.

// Perform the dialing operation.

// Transmit the result or, if cancelled, close the stream.

// Wait for dialing results or cancellation.

// Create the endpoint client.

func init() {
	// Register the Docker protocol handler with the synchronization package.
	synchronization.ProtocolHandlers[urlpkg.Protocol_Docker] = &protocolHandler{}
}
