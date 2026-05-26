package core

// IsDefault indicates whether or not the symbolic link mode is
// SymbolicLinkMode_SymbolicLinkModeDefault.
func (m SymbolicLinkMode) IsDefault() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (m SymbolicLinkMode) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (m *SymbolicLinkMode) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a symbolic link mode.

// Success.

// Supported indicates whether or not a particular symbolic link mode is a
// valid, non-default value.
func (m SymbolicLinkMode) Supported() bool { _ = "STUB: not implemented"; return false }

// Description returns a human-readable description of a symbolic link mode.
func (m SymbolicLinkMode) Description() string { _ = "STUB: not implemented"; return "" }
