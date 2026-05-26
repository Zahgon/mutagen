package filesystem

// LibexecPath computes the expected libexec path assuming a Filesystem
// Hierarchy Standard layout with the current executable located in the bin
// directory. It will return an error if the executable does not exist within
// the "bin" directory of such a layout, but it does not verify that the libexec
// directory exists.
func LibexecPath() (string, error) {
	_ = "STUB: not implemented"
	// Compute the path to the current executable.
	return "", nil
}

// If the executable path is a symbolic link, then perform resolution.
// Unfortunately there's no way to do this in a completely race-free
// fashion, but we're dealing with system prefixes here so it shouldn't be a
// problem.

// Check that the executable resides within a bin directory.

// Compute the expected libexec path.
