package core

// IsDefault indicates whether or not the synchronization mode is
// SynchronizationMode_SynchronizationModeDefault.
func (m SynchronizationMode) IsDefault() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (m SynchronizationMode) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (m *SynchronizationMode) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a synchronization mode.

// Success.

// Supported indicates whether or not a particular synchronization mode is a
// valid, non-default value.
func (m SynchronizationMode) Supported() bool { _ = "STUB: not implemented"; return false }

// Description returns a human-readable description of a synchronization mode.
func (m SynchronizationMode) Description() string { _ = "STUB: not implemented"; return "" }
