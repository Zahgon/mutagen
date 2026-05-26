package filesystem

import (
	"golang.org/x/sys/windows"
)

// queryHandleMetadata performs a metadata query using a Windows file handle. It
// must be passed the base name of the path used to open the handle. It supports
// files, directories, and symbolic links. It behavior designed to match that of
// the os.File.Stat method on Windows, specifically the unexported
// newFileStatFromGetFileInformationByHandle function and the various methods of
// os.fileStat. The behavior of this function, in particular its classification
// of file types in modes, must match that of the Go standard library in order
// for this package to function correctly (because the Windows implementation of
// this package partially relies on the standard os package and certain
// invariants will break if this classification behavior differs).
func queryHandleMetadata(name string, handle windows.Handle) (*Metadata, error) {
	_ = "STUB: not implemented"
	// Query the file type to ensure that it's an on-disk type (i.e. a file,
	// directory, or symbolic link).
	return nil, nil
}

// Perform a general file metadata query.

// If the handle refers to a reparse point, then determine whether or not
// it's a symbolic link. When dealing with reparse points, we need to
// perform an attribute query. Unfortunately this query isn't supported on
// some (or perhaps any) non-NTFS filesystems, so we can't use it as our
// only query (even though it returns the same FileAttributes field as the
// general query above). In cases where this query returns an invalid
// parameter error, we assume that we're on a non-NTFS filesystem, in which
// case symbolic links aren't supported in any case. See golang/go#29214 for
// more information.

// Determine whether or not we're dealing with a symbolic link. This
// logic follows that in os.fileStat.isSymlink.
//
// TODO: Update this definition once golang/go#42184 is resolved.

// Compute the mode. Note that the logic here needs to match that of
// os.fileStat.Mode.

// Compute the size.

// Compute the modification time.

// Success.
