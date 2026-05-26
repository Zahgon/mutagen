package forwarding

import (
	"context"

	"github.com/mutagen-io/mutagen/pkg/forwarding"
)

// Server provides an implementation of the Forwarding service.
type Server struct {
	// UnimplementedForwardingServer is the required base implementation.
	UnimplementedForwardingServer
	// manager is the underlying session manager.
	manager *forwarding.Manager
}

// NewServer creates a new session server.
func NewServer(manager *forwarding.Manager) *Server { _ = "STUB: not implemented"; return nil }

// Create creates a new session.
func (s *Server) Create(ctx context.Context, request *CreateRequest) (*CreateResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform creation.

// Success.

// List lists existing sessions.
func (s *Server) List(ctx context.Context, request *ListRequest) (*ListResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform listing.

// Success.

// Pause pauses existing sessions.
func (s *Server) Pause(ctx context.Context, request *PauseRequest) (*PauseResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform pausing.

// Success.

// Resume resumes existing sessions.
func (s *Server) Resume(ctx context.Context, request *ResumeRequest) (*ResumeResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform resuming.

// Success.

// Terminate terminates existing sessions.
func (s *Server) Terminate(ctx context.Context, request *TerminateRequest) (*TerminateResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform termination.

// Success.
