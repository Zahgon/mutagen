package ssh

import (
	"context"
	"io"

	"github.com/mutagen-io/mutagen/pkg/forwarding"
	"github.com/mutagen-io/mutagen/pkg/logging"
	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
)

// protocolHandler implements the forwarding.ProtocolHandler interface for
// connecting to remote endpoints over SSH. It uses the agent infrastructure
// over an SSH transport.
type protocolHandler struct{}

// dialResult provides asynchronous agent dialing results.
type dialResult struct {
	// stream is the stream returned by agent dialing.
	stream io.ReadWriteCloser
	// error is the error returned by agent dialing.
	error error
}

// Connect connects to an SSH endpoint.
func (p *protocolHandler) Connect(
	ctx context.Context,
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

// Create an SSH agent transport.

// Create a channel to deliver the dialing result.

// Perform dialing in a background Goroutine so that we can monitor for
// cancellation.

// Perform the dialing operation.

// Transmit the result or, if cancelled, close the stream.

// Wait for dialing results or cancellation.

// Create the endpoint.

func init() {
	// Register the SSH protocol handler with the forwarding package.
	forwarding.ProtocolHandlers[urlpkg.Protocol_SSH] = &protocolHandler{}
}
