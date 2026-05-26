//go:build !windows

package local

import (
	"net"
)

// listenWindowsNamedPipe returns an "unsupported" error on POSIX systems.
func listenWindowsNamedPipe(_ string) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

// isConflictingSocket returns whether or not a Unix domain socket listening
// error is due to a conflicting socket.
func isConflictingSocket(err error) bool {
	_ = "STUB: not implemented"
	// On POSIX systems, both of these errors are possible depending on the
	// nature of the conflicting on-disk object and (if it's a socket) whether
	// or not a listener is currently bound to it.
	return false
}
