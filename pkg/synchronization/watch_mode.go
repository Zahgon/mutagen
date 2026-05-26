package synchronization

// IsDefault indicates whether or not the watch mode is
// WatchMode_WatchModeDefault.
func (m WatchMode) IsDefault() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (m WatchMode) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (m *WatchMode) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a watch mode.

// Success.

// Supported indicates whether or not a particular watch mode is a valid,
// non-default value.
func (m WatchMode) Supported() bool { _ = "STUB: not implemented"; return false }

// Description returns a human-readable description of a watch mode.
func (m WatchMode) Description() string { _ = "STUB: not implemented"; return "" }
