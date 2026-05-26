package ignore

// IsDefault indicates whether or not the ignore syntax is Syntax_SyntaxDefault.
func (s Syntax) IsDefault() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (s Syntax) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (s *Syntax) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to an ignore syntax.

// Success.

// Supported indicates whether or not a particular ignore syntax is a valid,
// non-default value.
func (s Syntax) Supported() bool { _ = "STUB: not implemented"; return false }

// Description returns a human-readable description of an ignore syntax.
func (s Syntax) Description() string { _ = "STUB: not implemented"; return "" }
