package forwarding

import (
	"github.com/mutagen-io/mutagen/pkg/forwarding"
)

// Session represents a forwarding session.
type Session struct {
	// Identifier is the unique session identifier.
	Identifier string `json:"identifier"`
	// Version is the session version.
	Version forwarding.Version `json:"version"`
	// CreationTime is the session creation timestamp.
	CreationTime string `json:"creationTime"`
	// CreatingVersion is the version of Mutagen that created the session.
	CreatingVersion string `json:"creatingVersion"`
	// Source stores the source endpoint's configuration and state.
	Source Endpoint `json:"source"`
	// Destination stores the destination endpoint's configuration and state.
	Destination Endpoint `json:"destination"`
	// Configuration is the session configuration.
	Configuration
	// Name is the session name.
	Name string `json:"name,omitempty"`
	// Label are the session labels.
	Labels map[string]string `json:"labels,omitempty"`
	// Paused indicates whether or not the session is paused.
	Paused bool `json:"paused"`
	// Status is the session status.
	Status forwarding.Status `json:"status"`
	// SessionState stores state fields relevant to running sessions. It is
	// non-nil if and only if the session is unpaused.
	*SessionState
}

// SessionState encodes fields relevant to unpaused sessions.
type SessionState struct {
	// LastError is the last forwarding error to occur.
	LastError string `json:"lastError,omitempty"`
	// OpenConnections is the number of connections currently open and being
	// forwarded.
	OpenConnections uint64 `json:"openConnections"`
	// TotalConnections is the number of total connections that have been opened
	// and forwarded (including those that are currently open).
	TotalConnections uint64 `json:"totalConnections"`
	// TotalOutboundData is the total amount of data (in bytes) that has been
	// transmitted from source to destination across all forwarded connections.
	TotalOutboundData uint64 `json:"totalOutboundData"`
	// TotalInboundData is the total amount of data (in bytes) that has been
	// transmitted from destination to source across all forwarded connections.
	TotalInboundData uint64 `json:"totalInboundData"`
}

// loadFromInternal sets a session to match an internal Protocol Buffers session
// state representation. The session state must be valid.
func (s *Session) loadFromInternal(state *forwarding.State) {
	_ = "STUB: not implemented"
	// Propagate basic information.
	return
}

// Propagate endpoint information.

// Propagate configuration information.

// Propagate state information if the session isn't paused.

// ExportSessions converts a slice of internal session state representations to
// a slice of public session representations. It is guaranteed to return a
// non-nil value, even in the case of an empty slice.
func ExportSessions(states []*forwarding.State) []Session {
	_ = "STUB: not implemented"
	// Create the resulting slice.
	return nil
}

// Propagate session information

// Done.
