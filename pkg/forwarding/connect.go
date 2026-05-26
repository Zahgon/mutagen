package forwarding

import (
	"context"

	"github.com/mutagen-io/mutagen/pkg/logging"
	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
)

// ProtocolHandler defines the interface that protocol handlers must support in
// order to connect to endpoints.
type ProtocolHandler interface {
	// Connect connects to an endpoint using the connection parameters in the
	// provided URL and the specified prompter (if any). It then initializes the
	// endpoint using the specified parameters.
	Connect(
		ctx context.Context,
		logger *logging.Logger,
		url *urlpkg.URL,
		prompter string,
		session string,
		version Version,
		configuration *Configuration,
		source bool,
	) (Endpoint, error)
}

// ProtocolHandlers is a map of registered protocol handlers. It should only be
// modified during init() operations.
var ProtocolHandlers = map[urlpkg.Protocol]ProtocolHandler{}

// connect attempts to establish a connection to an endpoint.
func connect(
	ctx context.Context,
	logger *logging.Logger,
	url *urlpkg.URL,
	prompter string,
	session string,
	version Version,
	configuration *Configuration,
	source bool,
) (Endpoint, error) {
	_ = "STUB: not implemented"
	// Local the appropriate protocol handler.
	return *new(Endpoint), nil
}

// Dispatch the dialing.

// Success.
