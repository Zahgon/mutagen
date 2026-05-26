package forwarding

const (
	// sessionsDirectoryName is the name of the session storage directory within
	// the forwarding data directory.
	sessionsDirectoryName = "sessions"
)

// pathForSession computes the path to the serialized session for the given
// session identifier. An empty session identifier will return the sessions
// directory path.
func pathForSession(sessionIdentifier string) (string, error) {
	_ = "STUB: not implemented"
	// Compute/create the sessions directory.
	return "", nil
}

// Success.
