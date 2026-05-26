//go:build darwin || linux

package format

import (
	"golang.org/x/sys/unix"

	"github.com/mutagen-io/mutagen/pkg/filesystem"
)

// statfsRetryingOnEINTR is a wrapper around the statfs system call that retries
// on EINTR errors and returns on the first successful call or non-EINTR error.
func statfsRetryingOnEINTR(path string, metadata *unix.Statfs_t) error {
	_ = "STUB: not implemented"
	return nil
}

// fstatfsRetryingOnEINTR is a wrapper around the fstatfs system call that
// retries on EINTR errors and returns on the first successful call or non-EINTR
// error.
func fstatfsRetryingOnEINTR(fd int, metadata *unix.Statfs_t) error {
	_ = "STUB: not implemented"
	return nil
}

// QueryByPath queries the filesystem format for the specified path.
func QueryByPath(path string) (Format, error) {
	_ = "STUB: not implemented"
	// Perform a filesystem metadata query on the path.
	return *new(Format), nil
}

// Classify the filesystem.

// Query queries the filesystem format for the specified directory.
func Query(directory *filesystem.Directory) (Format, error) {
	_ = "STUB: not implemented"
	// Perform a filesystem metadata query on the directory.
	return *new(Format), nil
}

// Classify the filesystem.
