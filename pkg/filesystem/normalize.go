package filesystem

// tildeExpand attempts tilde expansion of paths beginning with ~/ or
// ~<username>/. On Windows, it additionally supports ~\ and ~<username>\.
func tildeExpand(path string) (string, error) {
	_ = "STUB: not implemented"
	// Only process relevant paths.
	return "", nil
}

// Find the first character in the path that's considered a path separator
// on the platform. Path seperators are always single-byte, so we can safely
// loop over the path's bytes.

// Divide the path into the "username" portion and the "subpath" portion -
// i.e. those portions coming before and after the separator, respectively.

// Compute the relevant home directory. If the username is empty, then we
// use the current user's home directory, otherwise we need to do a lookup.

// Compute the full path.

// Normalize normalizes a path, expanding home directory tildes, converting it
// to an absolute path, and cleaning the result.
func Normalize(path string) (string, error) {
	_ = "STUB: not implemented"
	// Expand any leading tilde.
	return "", nil
}

// Convert to an absolute path. This will also invoke filepath.Clean.

// Success.
