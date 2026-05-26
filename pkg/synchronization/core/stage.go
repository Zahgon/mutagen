package core

// stagingPathFinder recursively identifies paths/entries that need be staged in
// order to perform transitioning.
type stagingPathFinder struct {
	// paths is the list of paths for encountered file entries.
	paths []string
	// digests is the list of digests for encountered file entries, with length
	// and contents corresponding to paths.
	digests [][]byte
}

// find recursively searches for file entries that need staging.
func (f *stagingPathFinder) find(path string, entry *Entry) { _ = "STUB: not implemented"; return }

// Compute the prefix to add to content names to compute their paths.

// Process contents.

// TransitionDependencies analyzes a list of transitions and determines the file
// paths (and their corresponding digests) that will need to be provided in
// order to apply the transitions using Transition. It will return these paths
// in depth-first traversal order.
func TransitionDependencies(transitions []*Change) ([]string, [][]byte) {
	_ = "STUB: not implemented"
	// Create a path finder.
	return nil, nil
}

// Have it find paths for all the transitions.

// If this is a file-to-file transition and only the executability bit
// is changing, then we don't need to stage, because Transition will
// just modify the target on disk.

// Otherwise we need to perform a full scan.

// Success.
