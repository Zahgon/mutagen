package filesystem

import (
	"io"
)

// Open opens a filesystem path for traversal and/or other operations. It will
// return either a Directory or an io.ReadSeekCloser object (as an io.Closer for
// convenient closing access without casting), along with Metadata that can be
// used to determine the type of object being returned. Unless requested, this
// function does not allow the leaf component of path to be a symbolic link
// (though intermediate components of the path can be symbolic links and will be
// resolved in the resolution of the path), and an error will be returned if
// this is the case. However, if allowSymbolicLinkLeaf is true, then this
// function will allow resolution of a path leaf component that's a symbolic
// link. In this case, the referenced object must still be a directory or
// regular file, and the returned object will still be either a Directory or an
// io.ReadSeekCloser.
func Open(path string, allowSymbolicLinkLeaf bool) (io.Closer, *Metadata, error) {
	_ = "STUB: not implemented"
	// Verify that the provided path is absolute. This is a requirement on
	// Windows, where all of our operations are path-based.
	return *new(io.Closer), nil, nil
}

// Fix long paths.

// Convert the path to UTF-16.

// Open the path in a manner that is suitable for reading, doesn't allow for
// other threads or processes to delete or rename the file while open,
// avoids symbolic link traversal (at the path leaf), and has suitable
// semantics for both files and directories.

// Query handle metadata.

// Verify that we're not dealing with a symbolic link. If we are allowing
// symbolic links, then they should have been resolved by CreateFile.

// Handle os.File creation based on type.

// Success.
