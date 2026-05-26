package core

import (
	"github.com/mutagen-io/mutagen/pkg/filesystem"
)

// anyExecutableBitSet returns true if any executable bit is set on the file,
// false otherwise.
func anyExecutableBitSet(mode filesystem.Mode) bool { _ = "STUB: not implemented"; return false }

// EnsureDefaultFileModeValid validates that a user-provided default file mode
// is valid in the context of a given permissions mode.
func EnsureDefaultFileModeValid(permissionsMode PermissionsMode, mode filesystem.Mode) error {
	_ = "STUB: not implemented"
	// Verify that the mode is non-zero. This should never be the case, because
	// we treat a zero-value mode as unspecified.
	return nil
}

// Verify that only permission bits are set.

// Verify that no executability bits are set since they're controlled by
// Mutagen when propagating executability.

// Success.

// EnsureDefaultDirectoryModeValid validates that a user-provided default
// directory mode is valid in the context of a given permissions mode.
func EnsureDefaultDirectoryModeValid(_ PermissionsMode, mode filesystem.Mode) error {
	_ = "STUB: not implemented"
	// Verify that the mode is non-zero. This should never be the case, because
	// we treat a zero-value mode as unspecified.
	return nil
}

// Verify that only permission bits are set.

// Success.

// markExecutableForReaders sets the executable bit for the mode for any case
// where the corresponding read bit is set. It's worth noting that we implement
// this function as three separate checks (rather than a mask with a bit shift
// that's or'ed into the result) because we don't want to assume anything about
// the layout of permission bits.
func markExecutableForReaders(mode filesystem.Mode) filesystem.Mode {
	_ = "STUB: not implemented"
	// Set the user executable bit if necessary.
	return *new(filesystem.Mode)
}

// Set the group executable bit if necessary.

// Set the others executable bit if necessary.

// Done.
