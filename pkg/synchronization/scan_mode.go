package synchronization

// IsDefault indicates whether or not the scan mode is ScanMode_ScanModeDefault.
func (m ScanMode) IsDefault() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (m ScanMode) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (m *ScanMode) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a scan mode.

// Success.

// Supported indicates whether or not a particular scan mode is a valid,
// non-default value.
func (m ScanMode) Supported() bool { _ = "STUB: not implemented"; return false }

// Description returns a human-readable description of a scan mode.
func (m ScanMode) Description() string { _ = "STUB: not implemented"; return "" }
