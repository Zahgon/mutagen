package core

// Apply applies a series of changes to a base entry. It ignores the Old value
// for changes and only fails if the path to a change can't be resolved.
func Apply(base *Entry, changes []*Change) (*Entry, error) {
	_ = "STUB: not implemented"
	// If there are no changes, then we can just return the base unmodified.
	return nil, nil
}

// If there's only a single change and it's a root replacement, then we can
// just return the new entry.

// Create a deep copy of the base entry for mutation.

// Apply changes.

// Handle the special case of a root replacement. This typically won't
// occur mid-change-list, so we don't optimize for this case here in the
// same way that we do above.

// Crawl down the tree until we reach the parent of the target location.

// Depending on the new value, either set or remove the entry. If we're
// setting a new entry, then we need to create a mutable copy of it in
// case any subsequent changes affect it.

// Done.
