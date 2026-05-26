package forwarding

// IsDefault indicates whether or not the socket overwrite mode is
// SocketOverwriteMode_SocketOverwriteModeDefault.
func (m SocketOverwriteMode) IsDefault() bool { _ = "STUB: not implemented"; return false }

// AttemptOverwrite indicates whether or not the socket overwrite mode is
// SocketOverwriteMode_SocketOverwriteModeOverwrite.
func (m SocketOverwriteMode) AttemptOverwrite() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (m SocketOverwriteMode) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (m *SocketOverwriteMode) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a socket overwrite mode.

// Success.

// Supported indicates whether or not a particular socket overwrite mode is a
// valid, non-default value.
func (m SocketOverwriteMode) Supported() bool { _ = "STUB: not implemented"; return false }

// Description returns a human-readable description of a socket overwrite mode.
func (m SocketOverwriteMode) Description() string { _ = "STUB: not implemented"; return "" }
