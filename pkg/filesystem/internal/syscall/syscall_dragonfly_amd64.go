package syscall

// Symlinkat is a Go entry point for the symlinkat system call.
func Symlinkat(target string, directory int, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// _zero is a zero-value that can be used when a valid pointer is needed to 0
// bytes.
var _zero uintptr

// Readlinkat is a Go entry point for the readlinkat system call.
func Readlinkat(directory int, path string, buffer []byte) (int, error) {
	_ = "STUB: not implemented"
	// Extract a raw pointer to the path bytes.
	return 0, nil
}

// Extract a raw pointer to the buffer bytes.

// Perform the system call.

// Success.
