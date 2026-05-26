package behavior

// IsDefault indicates whether or not the probe mode is
// ProbeMode_ProbeModeDefault.
func (m ProbeMode) IsDefault() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (m ProbeMode) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (m *ProbeMode) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a probe mode.

// Success.

// Supported indicates whether or not a particular probe mode is a valid,
// non-default value.
func (m ProbeMode) Supported() bool { _ = "STUB: not implemented"; return false }

// Description returns a human-readable description of a probe mode.
func (m ProbeMode) Description() string { _ = "STUB: not implemented"; return "" }
