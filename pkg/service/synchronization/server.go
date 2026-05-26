package synchronization

import (
	"context"

	"github.com/mutagen-io/mutagen/pkg/synchronization"
)

// Server provides an implementation of the Synchronization service.
type Server struct {
	// UnimplementedSynchronizationServer is the required base implementation.
	UnimplementedSynchronizationServer
	// manager is the underlying session manager.
	manager *synchronization.Manager
}

// NewServer creates a new session server.
func NewServer(manager *synchronization.Manager) *Server { _ = "STUB: not implemented"; return nil }

// Create creates a new session.
func (s *Server) Create(ctx context.Context, request *CreateRequest) (*CreateResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform creation.

// Success.

// List queries session status.
func (s *Server) List(ctx context.Context, request *ListRequest) (*ListResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform listing.

// Success.

// Flush flushes sessions.
func (s *Server) Flush(ctx context.Context, request *FlushRequest) (*FlushResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform flushing.

// Success.

// Pause pauses sessions.
func (s *Server) Pause(ctx context.Context, request *PauseRequest) (*PauseResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform pausing.

// Success.

// Resume resumes sessions.
func (s *Server) Resume(ctx context.Context, request *ResumeRequest) (*ResumeResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform resuming.

// Success.

// Reset resets sessions.
func (s *Server) Reset(ctx context.Context, request *ResetRequest) (*ResetResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform resuming.

// Success.

// Terminate terminates sessions.
func (s *Server) Terminate(ctx context.Context, request *TerminateRequest) (*TerminateResponse, error) {
	_ = "STUB: not implemented"
	// Validate the request.
	return nil, nil
}

// Perform termination.

// Success.
