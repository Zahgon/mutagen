package core

// IsDefault indicates whether or not the permissions mode is
// PermissionsMode_PermissionsModeDefault.
func (m PermissionsMode) IsDefault() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (m PermissionsMode) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (m *PermissionsMode) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a permissions mode.

// Success.

// Supported indicates whether or not a particular permissions mode is a valid,
// non-default value.
func (m PermissionsMode) Supported() bool { _ = "STUB: not implemented"; return false }

// Description returns a human-readable description of a permissions mode.
func (m PermissionsMode) Description() string { _ = "STUB: not implemented"; return "" }
