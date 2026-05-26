package filesystem

import (
	"os"
)

const (
	// atomicWriteTemporaryNamePrefix is the file name prefix to use for
	// intermediate temporary files used in atomic writes.
	atomicWriteTemporaryNamePrefix = TemporaryNamePrefix + "atomic-write"
)

// WriteFileAtomic writes a file to disk in an atomic fashion by using an
// intermediate temporary file that is swapped in place using a rename
// operation.
func WriteFileAtomic(path string, data []byte, permissions os.FileMode) error {
	_ = "STUB: not implemented"
	// Create a temporary file. The os package already uses secure permissions
	// for creating temporary files, so we don't need to change them.
	return nil
}

// Write data.

// Close out the file.

// Set the file's permissions.

// Rename the file.

// Success.
