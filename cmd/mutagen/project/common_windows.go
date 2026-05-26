package project

// runInShell runs the specified command using the system shell. On Windows
// systems, this is %ComSpec% (with a fallback to a fully qualified cmd.exe if
// %ComSpec% is not an absolute path (which includes cases where it's empty)).
func runInShell(command string) error {
	_ = "STUB: not implemented"
	// Determine the shell to use.
	return nil
}

// Set up the process.

// Run the process and wait for its completion.
