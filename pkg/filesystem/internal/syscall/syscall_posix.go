//go:build linux || darwin || freebsd || openbsd || netbsd

package syscall

// Symlinkat is a Go entry point for the symlinkat system call.
func Symlinkat(target string, directory int, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// Readlinkat is a Go entry point for the readlinkat system call.
func Readlinkat(directory int, path string, buffer []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
