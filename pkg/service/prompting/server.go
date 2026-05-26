package prompting

import (
	"context"
)

// Server provides an implementation of the Prompting service.
type Server struct {
	// UnimplementedPromptingServer is the required base implementation.
	UnimplementedPromptingServer
}

// NewServer creates a new prompt server.
func NewServer() *Server {
	_ = "STUB: not implemented"

	// Host performs prompt hosting.
	return nil
}

func (s *Server) Host(stream Prompting_HostServer) error {
	_ = "STUB: not implemented"
	// Receive and validate the initial request.
	return nil
}

// Create a unique identifier for the prompter.

// Send the initial response.

// Extract the request context.

// Wrap the stream to create a prompter.

// Register the prompter.

// Wait for the request or connection to be terminated.

// Unregister the promper.

// Success.

// asyncPromptResponse provides a structure for returning prompt results
// asynchronously, allowing prompting to be cancelled.
type asyncPromptResponse struct {
	// response is the response returned by the prompter.
	response string
	// error is the error returned by the prompter.
	error error
}

// Prompt performs prompting against registered prompters.
func (s *Server) Prompt(ctx context.Context, request *PromptRequest) (*PromptResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform prompting from the global registry asynchronously.
// TODO: Should we build cancellation into the Prompter interface itself?

// Wait for a response or cancellation.
