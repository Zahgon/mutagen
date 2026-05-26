package local

import (
	"net"
	"syscall"
)

// listenWindowsNamedPipe attempts to create a named pipe listener on Windows.
func listenWindowsNamedPipe(address string) (net.Listener, error) {
	_ = "STUB: not implemented"
	// Compute the SID of the user.
	return *new(net.Listener), nil
}

// Create the security descriptor for the pipe. This is constructed using
// the Security Descriptor Definition Language (SDDL) (the Discretionary
// Access Control List (DACL) format), where the value in parentheses is an
// Access Control Entry (ACE) string. The P flag in the DACL prevents
// inherited permissions. The ACE string in this case grants "Generic All"
// (GA) permissions to its associated SID. More information can be found
// here:
//  SDDL: https://msdn.microsoft.com/en-us/library/windows/desktop/aa379570(v=vs.85).aspx
//  ACEs: https://msdn.microsoft.com/en-us/library/windows/desktop/aa374928(v=vs.85).aspx
//  SIDs: https://msdn.microsoft.com/en-us/library/windows/desktop/aa379602(v=vs.85).aspx

// Set up the named pipe configuration. We set the named pipe to message
// mode because the go-winio named pipe connection implementation supports
// CloseWrite for message mode pipes by using a zero-length message to
// encode write closure. This is a signaling mechanism unique to go-winio,
// but it has no adverse effect on any other named pipe implementations,
// which simply won't attach any meaning to a zero-length message. The
// resulting connection will still behave like a stream-based connection,
// the only difference is that the name pipe will ensure each write's bytes
// are grouped into a single message. Additionally, message mode pipes can
// still be read as byte streams (essentially just ignoring message
// boundaries), so there's no loss of compatibility.
// https://docs.microsoft.com/en-us/windows/win32/ipc/named-pipe-type-read-and-wait-modes

// Attempt to create the listener.

const (
	// WSAEADDRINUSE is the Winsock API error equivalent of POSIX's EADDRINUSE.
	// https://docs.microsoft.com/en-us/windows/win32/winsock/windows-sockets-error-codes-2
	WSAEADDRINUSE syscall.Errno = 10048
)

// isConflictingSocket returns whether or not a Unix domain socket listening
// error is due to a conflicting socket.
func isConflictingSocket(err error) bool { _ = "STUB: not implemented"; return false }
