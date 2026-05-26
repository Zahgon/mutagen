package platform

// ExecutableName computes the name for an executable for a given base name on a
// specified operating system.
func ExecutableName(base, goos string) string {
	_ = "STUB: not implemented"
	// If we're on Windows, append ".exe".
	return ""
}

// Otherwise return the base name unmodified.
