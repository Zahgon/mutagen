package behavior

import (
	"github.com/mutagen-io/mutagen/pkg/filesystem"
)

const (
	// composedFileNamePrefix is the prefix used for temporary files created by
	// the Unicode decomposition test. It is in NFC form.
	composedFileNamePrefix = filesystem.TemporaryNamePrefix + "unicode-test-\xc3\xa9ntry"
	// decomposedFileNamePrefix is the NFD equivalent of composedFileNamePrefix.
	decomposedFileNamePrefix = filesystem.TemporaryNamePrefix + "unicode-test-\x65\xcc\x81ntry"
)

// DecomposesUnicodeByPath determines whether or not the filesystem on which the
// directory at the specified path resides decomposes Unicode filenames. The
// second value returned by this function indicates whether or not probe files
// were used in determining behavior.
func DecomposesUnicodeByPath(path string, probeMode ProbeMode) (bool, bool, error) {
	_ = "STUB: not implemented"
	// Check the filesystem probing mode and see if we can return an assumption.
	return false, false, nil
}

// Check if we have a fast test that will work.

// Create and close a temporary file using the composed filename.

// Grab the file's name. This is calculated from the parameters passed to
// TempFile, not by reading from the OS, so it will still be in a composed
// form. Also calculate a decomposed variant.

// Defer removal of the file. Since we don't know whether the filesystem is
// also normalization-insensitive, we try both compositions.

// Grab the contents of the path.

// Loop through contents and see if we find a match for the decomposed file
// name. It doesn't even need to be our file, though it probably will be.

// If we didn't find any match, something's fishy.

// DecomposesUnicode determines whether or not the specified directory (and its
// underlying filesystem) decomposes Unicode filenames. The second value
// returned by this function indicates whether or not probe files were used in
// determining behavior.
func DecomposesUnicode(directory *filesystem.Directory, probeMode ProbeMode) (bool, bool, error) {
	_ = "STUB: not implemented"
	// Check the filesystem probing mode and see if we can return an assumption.
	return false, false, nil
}

// Check if we have a fast test that will work.

// Create and close a temporary file using the composed filename.

// The name returned from CreateTemporaryFile is calculated from the
// provided pattern, so it will still be in a composed form. Compute the
// decomposed variant.

// Defer removal of the file. Since we don't know whether the filesystem is
// also normalization-insensitive, we try both compositions.

// HACK: If we're on Linux, then re-open the directory after creating the
// temporary file (and defer closure of the re-opened copy). This is
// necessary to work around an issue with osxfs where a directory descriptor
// can't be used to list contents created after the descriptor was opened
// (due either to aggressive caching or some sort of implementation bug).
// See issue #73 for more details. Ideally we'd restrict this workaround to
// osxfs, but we can't actually detect osxfs specifically because the statfs
// type field just indicates that it's a FUSE filesystem. Even if we wanted
// to restrict this behavior to just FUSE filesystems, the statfs call is
// going to be about the same cost (if not more expensive) than the re-open
// call, so it's best to just do this in all cases on Linux. This isn't such
// a big deal since this function is only called once per scan, and we may
// hit a fast path above anyway.

// Grab the content names in the directory.

// Loop through the names and see if we find a match for either the composed
// or decomposed name.

// If we didn't find any match, something's fishy.
