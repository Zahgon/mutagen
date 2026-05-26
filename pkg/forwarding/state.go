package forwarding

// Description returns a human-readable description of the session status.
func (s Status) Description() string { _ = "STUB: not implemented"; return "" }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (s Status) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (s *Status) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a forwarding status.

// Success.

// ensureValid ensures that EndpointState's invariants are respected.
func (s *EndpointState) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil endpoint state is not valid.
	return nil
}

// We could perform additional validation based on the session status and
// the endpoint connectivity, but it would be prohibitively complex, and all
// we're really concerned about here is memory safety and other structural
// invariants.

// Success.

// EnsureValid ensures that State's invariants are respected.
func (s *State) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil state is not valid.
	return nil
}

// We could perform additional validation based on the session status, but
// it would be prohibitively complex, and all we're really concerned about
// here is memory safety and other structural invariants.

// Ensure the session is valid.

// Ensure that the connection counts are sane.

// Ensure that endpoint states are valid.

// Success.
