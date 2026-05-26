package behavior

import (
	"github.com/mutagen-io/mutagen/pkg/filesystem"
)

const (
	// executabilityProbeFileNamePrefix is the prefix used for temporary files
	// created by the executability preservation test.
	executabilityProbeFileNamePrefix = filesystem.TemporaryNamePrefix + "executability-test"
)

// PreservesExecutabilityByPath determines whether or not the filesystem on
// which the directory at the specified path resides preserves POSIX
// executability bits. It allows for the path leaf to be a symbolic link. The
// second value returned by this function indicates whether or not probe files
// were used in determining behavior.
func PreservesExecutabilityByPath(path string, probeMode ProbeMode) (bool, bool, error) {
	_ = "STUB: not implemented"
	// Check the filesystem probing mode and see if we can return an assumption.
	return false, false, nil
}

// Check if we have a fast test that will work. If we're on Windows, we
// enforce that the fast path was used. There is some code below, namely the
// use of os.File's Chmod method (and possibly the os.File's Stat method,
// which may be racey on Windows), that won't work on Windows (though it
// could possibly be adapted in case we add a force-probe probe mode), which
// is why we require that the fast path succeeds on Windows.

// Create a temporary file.

// Ensure that the file is cleaned up and removed when we're done.

// Mark the file as user-executable. We use the os.File-based Chmod here
// since this code only runs on POSIX systems where this is supported.

// Grab the file statistics and check for executability. We enforce that
// only the user-executable bit is set, because filesystems that don't
// preserve executability on POSIX systems (e.g. FAT32 on older versions of
// Darwin) sometimes mark every file as having every executable bit set,
// which is another type of non-preserving behavior. This behavior is not
// universal (e.g. FAT32 on Linux marks every file as having no executable
// bit set), but this test should catch many of these cases. It is, however,
// susceptible to umask settings (e.g. umask 0022), so it's not foolproof.

// PreservesExecutability determines whether or not the specified directory (and
// its underlying filesystem) preserves POSIX executability bits. The second
// value returned by this function indicates whether or not probe files were
// used in determining behavior.
func PreservesExecutability(directory *filesystem.Directory, probeMode ProbeMode) (bool, bool, error) {
	_ = "STUB: not implemented"
	// Check the filesystem probing mode and see if we can return an assumption.
	return false, false, nil
}

// Check if we have a fast test that will work. If we're on Windows, we
// enforce that the fast path was used. There is some code below, namely the
// use of os.File's Chmod method (and possibly the os.File's Stat method,
// which may be racey on Windows), that won't work on Windows (though it
// could possibly be adapted in case we add a force-probe probe mode), which
// is why we require that the fast path succeeds on Windows.

// Create a temporary file.

// Ensure that the file is cleaned up and removed when we're done.

// HACK: Convert the file to an os.File object for race-free Chmod and Stat
// access. This is an acceptable hack since we live inside the same package
// as the Directory implementation.

// Mark the file as user-executable. We use the os.File-based Chmod here
// since this code only runs on POSIX systems where this is supported.

// Grab the file statistics and check for executability. We enforce that
// only the user-executable bit is set, because filesystems that don't
// preserve executability on POSIX systems (e.g. FAT32 on Darwin) sometimes
// mark every file as having every executable bit set, which is another type
// of non-preserving behavior. This behavior is not universal (e.g. FAT32 on
// Linux marks every file as having no executable bit set), but this test
// should be.
