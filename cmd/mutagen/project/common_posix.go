//go:build !windows

package project

// runInShell runs the specified command using the system shell. On POSIX
// systems, this is /bin/sh.
func runInShell(command string) error {
	_ = "STUB: not implemented"
	// Set up the process.
	return nil
}

// Run the process and wait for its completion.
