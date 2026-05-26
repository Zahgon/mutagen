package synchronization

import (
	"github.com/mutagen-io/mutagen/pkg/synchronization/core"
)

// oneEndpointEmptiedRoot determines whether or not one endpoint (but not both)
// transitioned from a directory root with a non-trivial amount of content to a
// directory root without any content.
func oneEndpointEmptiedRoot(ancestor, alpha, beta *core.Entry) bool {
	_ = "STUB: not implemented"
	// Check that all three entries are directories. If not, then this check
	// doesn't apply.
	return false
}

// Check whether or not the ancestor has a non-trivial amount of content.
// We define a non-trivial amount of content as two or more entries that are
// immediate children of the synchronization root. If that's not the case,
// then this check doesn't apply.

// Check if alpha deleted all content within the root.

// Check if beta deleted all content within the root.

// Determine whether one (and only one) endpoint emptied root content.

// containsRootDeletion determines whether or not any of the specified changes
// is a root deletion change.
func containsRootDeletion(changes []*core.Change) bool {
	_ = "STUB: not implemented"
	// Look for root deletions.
	return false
}

// Done.

// containsRootTypeChange determines whether or not any of the specified changes
// is a root type change.
func containsRootTypeChange(changes []*core.Change) bool {
	_ = "STUB: not implemented"
	// Look for root type changes.
	return false
}

// Done.

// filteredPathsAreSubset checks whether or not a slice of filtered paths is a
// subset of a larger slice of unfiltered paths. The paths in the filtered slice
// must share the same relative ordering as in the original slice.
func filteredPathsAreSubset(filteredPaths, originalPaths []string) bool {
	_ = "STUB: not implemented"
	// Loop over the list of filtered paths.
	return false
}

// Track whether or not we find a match for this path in what remains of
// the original path list.

// Loop over what remains of the original paths.

// Check if we found a match.

// Success.
