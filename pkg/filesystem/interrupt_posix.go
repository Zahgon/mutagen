//go:build !windows

package filesystem

import (
	"golang.org/x/sys/unix"
)

// openatRetryingOnEINTR is a wrapper around the openat system call that retries
// on EINTR errors and returns on the first successful call or non-EINTR error.
func openatRetryingOnEINTR(directory int, path string, flags int, mode uint32) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// readRetryingOnEINTR is a wrapper around the read system call that retries on
// EINTR errors and returns on the first successful call or non-EINTR error.
func readRetryingOnEINTR(file int, buffer []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// seekConsideringEINTR is a direct passthrough to the lseek system call that
// doesn't retry on EINTR. It's only defined to highlight the intentional
// absence of seekRetryingOnEINTR. seekRetryingOnEINTR is left unimplemented
// because it would have to handle cases of partially successful seeks (which
// would be complicated in the case of SEEK_CUR or other relative whence values)
// and because POSIX doesn't specify that lseek can return EINTR. The Go
// standard library and runtime also invoke lseek without retrying on EINTR.
func seekConsideringEINTR(file int, offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// closeConsideringEINTR is a direct passthrough to the close system call that
// doesn't retry on EINTR. It's only defined to highlight the intentional
// absence of closeRetryingOnEINTR. closeRetryingOnEINTR is left unimplemented
// because POSIX makes no guarantees about the state of a file descriptor in the
// event of an EINTR error, and thus retrying closure could lead to a race
// condition with file descriptor re-use if the file is, in fact, closed. This
// is the same policy adopted by the Go standard library and runtime.
func closeConsideringEINTR(file int) error { _ = "STUB: not implemented"; return nil }

// mkdiratRetryingOnEINTR is a wrapper around the mkdirat system call that
// retries on EINTR errors and returns on the first successful call or non-EINTR
// error.
func mkdiratRetryingOnEINTR(directory int, path string, mode uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// renameatRetryingOnEINTR is a wrapper around the renameat system call that
// retries on EINTR errors and returns on the first successful call or non-EINTR
// error.
func renameatRetryingOnEINTR(oldDirectory int, oldPath string, newDirectory int, newPath string) error {
	_ = "STUB: not implemented"
	return nil
}

// unlinkatRetryingOnEINTR is a wrapper around the unlinkat system call that
// retries on EINTR errors and returns on the first successful call or non-EINTR
// error.
func unlinkatRetryingOnEINTR(directory int, path string, flags int) error {
	_ = "STUB: not implemented"
	return nil
}

// fstatRetryingOnEINTR is a wrapper around the fstat system call that retries
// on EINTR errors and returns on the first successful call or non-EINTR error.
func fstatRetryingOnEINTR(file int, metadata *unix.Stat_t) error {
	_ = "STUB: not implemented"
	return nil
}

// fchmodRetryingOnEINTR is a wrapper around the fchmod system call that retries
// on EINTR errors and returns on the first successful call or non-EINTR error.
func fchmodRetryingOnEINTR(file int, mode uint32) error { _ = "STUB: not implemented"; return nil }

// fstatatRetryingOnEINTR is a wrapper around the fstatat system call that
// retries on EINTR errors and returns on the first successful call or non-EINTR
// error.
func fstatatRetryingOnEINTR(directory int, path string, metadata *unix.Stat_t, flags int) error {
	_ = "STUB: not implemented"
	return nil
}

// fchmodatRetryingOnEINTR is a wrapper around the fchmodat system call that
// retries on EINTR errors and returns on the first successful call or non-EINTR
// error.
func fchmodatRetryingOnEINTR(directory int, path string, mode uint32, flags int) error {
	_ = "STUB: not implemented"
	return nil
}

// fchownatRetryingOnEINTR is a wrapper around the fchownat system call that
// retries on EINTR errors and returns on the first successful call or non-EINTR
// error.
func fchownatRetryingOnEINTR(directory int, path string, uid int, gid int, flags int) error {
	_ = "STUB: not implemented"
	return nil
}

// symlinkatRetryingOnEINTR is a wrapper around the symlinkat system call that
// retries on EINTR errors and returns on the first successful call or non-EINTR
// error.
func symlinkatRetryingOnEINTR(target string, directory int, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// readlinkatRetryingOnEINTR is a wrapper around the readlinkat system call that
// retries on EINTR errors and returns on the first successful call or non-EINTR
// error.
func readlinkatRetryingOnEINTR(directory int, path string, buffer []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
