package synchronization

// IsDefault indicates whether or not the staging mode is
// StageMode_StageModeDefault.
func (m StageMode) IsDefault() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (m StageMode) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (m *StageMode) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a staging mode.

// Success.

// Supported indicates whether or not a particular staging mode is a valid,
// non-default value.
func (m StageMode) Supported() bool { _ = "STUB: not implemented"; return false }

// Description returns a human-readable description of a staging mode.
func (m StageMode) Description() string { _ = "STUB: not implemented"; return "" }
