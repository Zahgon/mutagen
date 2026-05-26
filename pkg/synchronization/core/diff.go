package core

// differ provides recursive diffing infrastructure.
type differ struct {
	// changes is the list of changes being tracked by the diff.
	changes []*Change
}

// diff is the recursive diff entry point.
func (d *differ) diff(path string, base, target *Entry) {
	_ = "STUB: not implemented"
	// If the nodes at this path aren't equal, then do a complete replacement.
	return
}

// Extract contents.

// Compute the prefix to add to content names to compute their paths.

// The nodes were equal at this path, so check their contents.

// diff performs a diff operation between a base and target entry (treating both
// as rooted at the specified path) and generates a list of changes that, if
// applied to base, would transform it into target.
func diff(path string, base, target *Entry) []*Change {
	_ = "STUB: not implemented"
	// Create the differ.
	return nil
}

// Populate changes.

// Done.

// Diff performs a diff operation between a base and target entry and generates
// a list of changes that, if applied to base, would transform it into target.
func Diff(base, target *Entry) []*Change { _ = "STUB: not implemented"; return nil }
