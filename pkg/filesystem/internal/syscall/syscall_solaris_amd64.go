package syscall

import (
	"syscall"
)

// syscallFunc is a handle type for Solaris libc functions.
type syscallFunc uintptr

// sysvicall6 is a handle for the Solaris system call implementation in the
// syscall package (which is itself just a thin wrapper around the actual
// implementation in the runtime package). It is wired up via assembly in
// syscall_asm_solaris_amd64.s.
func sysvicall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

//go:cgo_import_dynamic libc_symlinkat symlinkat "libc.so"
//go:cgo_import_dynamic libc_readlinkat readlinkat "libc.so"

//go:linkname procSymlinkat libc_symlinkat
//go:linkname procReadlinkat libc_readlinkat

var (
	// procSymlinkat is a handle for the symlinkat libc function.
	procSymlinkat,
	// procReadlinkat is a handle for the readlinkat libc function.
	procReadlinkat syscallFunc
)

// Symlinkat is a Go entry point for the symlinkat system call.
func Symlinkat(target string, directory int, path string) error {
	_ = "STUB: not implemented"
	// Extract a raw pointer to the target path bytes.
	return nil
}

// Extract a raw pointer to the path bytes.

// Perform the system call.

// Success.

// Readlinkat is a Go entry point for the readlinkat system call.
func Readlinkat(directory int, path string, buffer []byte) (int, error) {
	_ = "STUB: not implemented"
	// Extract a raw pointer to the path bytes.
	return 0, nil
}

// Extract a raw pointer to the buffer bytes.

// Perform the system call.

// Success.
