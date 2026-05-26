package filesystem

const (
	// ModePermissionsMask is a bit mask that isolates portable permission bits.
	ModePermissionsMask = Mode(0777)

	// ModePermissionUserRead is the user readable bit.
	ModePermissionUserRead = Mode(0400)
	// ModePermissionUserWrite is the user writable bit.
	ModePermissionUserWrite = Mode(0200)
	// ModePermissionUserExecute is the user executable bit.
	ModePermissionUserExecute = Mode(0100)
	// ModePermissionGroupRead is the group readable bit.
	ModePermissionGroupRead = Mode(0040)
	// ModePermissionGroupWrite is the group writable bit.
	ModePermissionGroupWrite = Mode(0020)
	// ModePermissionGroupExecute is the group executable bit.
	ModePermissionGroupExecute = Mode(0010)
	// ModePermissionOthersRead is the others readable bit.
	ModePermissionOthersRead = Mode(0004)
	// ModePermissionOthersWrite is the others writable bit.
	ModePermissionOthersWrite = Mode(0002)
	// ModePermissionOthersExecute is the others executable bit.
	ModePermissionOthersExecute = Mode(0001)
)

// parseMode parses a user-specified octal string and verifies that it is
// limited to the bits specified in mask. It allows, but does not require, the
// string to begin with a 0 (or several 0s). The provided string must not be
// empty.
func parseMode(value string, mask Mode) (Mode, error) {
	_ = "STUB: not implemented"
	return *new(Mode), nil
}

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (m Mode) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText. It requires
// that the specified mode bits lie within ModePermissionsMask, otherwise an
// error is returned. If an error is returned, the mode is unmodified.
func (m *Mode) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Perform parsing. We only allow the mode itself to be modified if parsing
// is successful.

// Success.
