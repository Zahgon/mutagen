package environment

// ParseBlock parses an environment variable block of the form
// VAR1=value1[\r]\nVAR2=value2[\r]\n... into a slice of KEY=value strings. It
// opts for performance over extensive format validation.
func ParseBlock(block string) []string {
	_ = "STUB: not implemented"
	// Replace all instances of \r\n with \n.
	return nil
}

// Trim whitespace from around the block.

// Split the block into individual lines.
