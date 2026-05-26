//go:build !windows

package ipc

import (
	"context"
	"net"
)

// DialContext attempts to establish an IPC connection, timing out if the
// provided context expires.
func DialContext(ctx context.Context, path string) (net.Conn, error) {
	_ = "STUB: not implemented"
	// Create a zero-valued dialer, which will have the same dialing behavior as
	// the raw dialing functions.
	return *new(net.Conn), nil
}

// Perform dialing.

// NewListener creates a new IPC listener.
func NewListener(path string) (net.Listener, error) {
	_ = "STUB: not implemented"
	// Create the listener.
	return *new(net.Listener), nil
}

// Explicitly set socket permissions. Unfortunately we can't do this
// atomically on socket creation, but we can do it quickly.

// Success.
