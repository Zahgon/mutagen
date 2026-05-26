//go:build !windows

package filesystem

// MarkHidden ensures that a path is hidden.
func MarkHidden(path string) error {
	_ = "STUB: not implemented"
	// POSIX platforms don't have the notion of a hidden attribute, they only
	// hide dot-prefixed paths, so ensure that the path begins with a dot.
	return nil
}

// Success.
