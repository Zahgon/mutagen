package environment

// ToMap converts an environment variable specification from a slice of
// "KEY=value" strings to a map with equivalent contents. Any entries not
// adhering to the specified format are ignored. Entries are processed in order,
// meaning that the last entry seen for a key will be what populates the map.
func ToMap(environment []string) map[string]string {
	_ = "STUB: not implemented"
	// Allocate result storage.
	return nil
}

// Convert variables.

// Done.

// FromMap converts a map of environment variables into a slice of "KEY=value"
// strings. If the provided environment is nil, then the resulting slice will be
// nil. If the provided environment is non-nil but empty, then the resulting
// slice will be empty. These two properties are critical to usage with the
// os/exec package.
func FromMap(environment map[string]string) []string {
	_ = "STUB: not implemented"
	// If the environment is nil, then return a nil slice.
	return nil
}

// Allocate result storage.

// Convert entries.

// Done.
