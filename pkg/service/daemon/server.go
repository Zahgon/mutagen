package daemon

import (
	"context"
	"time"
)

const (
	// housekeepingInterval is the interval at which housekeeping will be
	// invoked by the daemon.
	housekeepingInterval = 24 * time.Hour
)

// Server provides an implementation of the Daemon service.
type Server struct {
	// UnimplementedDaemonServer is the required base implementation.
	UnimplementedDaemonServer
	// Termination is populated with requests from clients invoking the shutdown
	// method over RPC. It can be ignored by daemon host processes wishing to
	// ignore temination requests originating from clients. The channel is
	// buffered and non-blocking, so it doesn't need to be serviced by the
	// daemon host-process at all - additional incoming shutdown requests will
	// just bounce off once the channel is populated. We do this, instead of
	// closing the channel, because we can't close the channel multiple times.
	Termination chan struct{}
	// workerCtx is the context regulating the server's internal operations.
	workerCtx context.Context
	// shutdown is the context cancellation function for the server's internal
	// operation context.
	shutdown context.CancelFunc
}

// NewServer creates a new daemon server.
func NewServer() *Server {
	_ = "STUB: not implemented"
	// Create a cancellable context for daemon background operations.
	return nil
}

// Create the server.

// Start the housekeeping Goroutine.

// Done.

// housekeep provides regular housekeeping facilities for the daemon.
func (s *Server) housekeep() {
	_ = "STUB: not implemented"
	// Perform an initial housekeeping operation since the ticker won't fire
	// straight away.
	return
}

// Create a ticker to regulate housekeeping and defer its shutdown.

// Loop and wait for the ticker or cancellation.

// Shutdown gracefully shuts down server resources.
func (s *Server) Shutdown() {
	_ = "STUB: not implemented"
	// Cancel all internal operations.
	return
}

// Version provides version information.
func (s *Server) Version(_ context.Context, _ *VersionRequest) (*VersionResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Terminate requests daemon termination.
func (s *Server) Terminate(_ context.Context, _ *TerminateRequest) (*TerminateResponse, error) {
	_ = "STUB: not implemented"
	// Send the termination request in a non-blocking manner.
	return nil, nil
}

// Success.
