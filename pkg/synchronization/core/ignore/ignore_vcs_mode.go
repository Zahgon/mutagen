package ignore

// IsDefault indicates whether or not the VCS ignore mode is
// IgnoreVCSMode_IgnoreVCSModeDefault.
func (m IgnoreVCSMode) IsDefault() bool { _ = "STUB: not implemented"; return false }

// MarshalJSON implements encoding/json.Marshaler.MarshalJSON.
func (m IgnoreVCSMode) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (m *IgnoreVCSMode) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a VCS mode.

// Success.

// UnmarshalJSON implements encoding/json.Unmarshaler.UnmarshalJSON.
func (m *IgnoreVCSMode) UnmarshalJSON(textBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Supported indicates whether or not a particular VCS ignore mode is a valid,
// non-default value.
func (m IgnoreVCSMode) Supported() bool { _ = "STUB: not implemented"; return false }

// Description returns a human-readable description of a VCS ignore mode.
func (m IgnoreVCSMode) Description() string { _ = "STUB: not implemented"; return "" }
