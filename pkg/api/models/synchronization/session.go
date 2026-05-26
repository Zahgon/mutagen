package synchronization

import (
	"github.com/mutagen-io/mutagen/pkg/synchronization"
)

// Session represents a synchronization session.
type Session struct {
	// Identifier is the unique session identifier.
	Identifier string `json:"identifier"`
	// Version is the session version.
	Version synchronization.Version `json:"version"`
	// CreationTime is the session creation timestamp.
	CreationTime string `json:"creationTime"`
	// CreatingVersion is the version of Mutagen that created the session.
	CreatingVersion string `json:"creatingVersion"`
	// Alpha stores the alpha endpoint's configuration and state.
	Alpha Endpoint `json:"alpha"`
	// Beta stores the beta endpoint's configuration and state.
	Beta Endpoint `json:"beta"`
	// Configuration is the session configuration.
	Configuration
	// Name is the session name.
	Name string `json:"name,omitempty"`
	// Label are the session labels.
	Labels map[string]string `json:"labels,omitempty"`
	// Paused indicates whether or not the session is paused.
	Paused bool `json:"paused"`
	// Status is the session status.
	Status synchronization.Status `json:"status"`
	// SessionState stores state fields relevant to running sessions. It is
	// non-nil if and only if the session is unpaused.
	*SessionState
}

// SessionState encodes fields relevant to unpaused sessions.
type SessionState struct {
	// LastError is the last synchronization error to occur.
	LastError string `json:"lastError,omitempty"`
	// SuccessfulCycles is the number of successful synchronization cycles to
	// occur since successfully connecting to the endpoints.
	SuccessfulCycles uint64 `json:"successfulCycles,omitempty"`
	// Conflicts are the conflicts that identified during reconciliation. This
	// list may be a truncated version of the full list if too many conflicts
	// are encountered to report via the API.
	Conflicts []Conflict `json:"conflicts,omitempty"`
	// ExcludedConflicts is the number of conflicts that have been excluded from
	// Conflicts due to truncation. This value can only be non-zero if conflicts
	// is non-empty.
	ExcludedConflicts uint64 `json:"excludedConflicts,omitempty"`
}

// loadFromInternal sets a session to match an internal Protocol Buffers session
// state representation. The session state must be valid.
func (s *Session) loadFromInternal(state *synchronization.State) {
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
func ExportSessions(states []*synchronization.State) []Session {
	_ = "STUB: not implemented"
	// Create the resulting slice.
	return nil
}

// Propagate session information

// Done.
