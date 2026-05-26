package types

// ByteSize is a uint64 value that supports unmarshalling from both
// human-friendly string representations and numeric representations. It can be
// cast to a uint64 value, where it represents a byte count.
type ByteSize uint64

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (s *ByteSize) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Parse and store the value.

// Success.
