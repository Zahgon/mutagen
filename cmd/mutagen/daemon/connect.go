package daemon

import (
	"time"

	"google.golang.org/grpc"
)

const (
	// dialTimeout is the timeout to use when attempting to connect to the
	// daemon IPC endpoint.
	dialTimeout = 500 * time.Millisecond
	// autostartWaitInterval is the wait period between reconnect attempts after
	// autostarting the daemon.
	autostartWaitInterval = 100 * time.Millisecond
	// autostartRetryCount is the number of times to try reconnecting after
	// autostarting the daemon.
	autostartRetryCount = 10
)

// Connect creates a new daemon client connection and optionally verifies that
// the daemon version matches the current process' version.
func Connect(autostart, enforceVersionMatch bool) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	// Compute the path to the daemon IPC endpoint.
	return nil, nil
}

// Check if autostart has been disabled programmatically or by an
// environment variable.

// Create a status line printer and defer a clear.

// Perform dialing in a loop until failure or success.

// Create a context to timeout the dial.

// Attempt to dial.

// Cancel the dialing context. If the dialing operation has already
// succeeded, this has no effect, but it is necessary to clean up the
// Goroutine that backs the context.

// Check for errors.

// Handle failure due to timeouts.

// If autostart is enabled, and we have attempts remaining, then
// try autostarting, waiting, and retrying.

// Otherwise just fail due to the timeout.

// If we failed for any other reason, then bail.

// Print a notice if we started the daemon.

// We've successfully dialed, so break out of the dialing loop.

// If requested, verify that the daemon version matches the current process'
// version.

// Success.
