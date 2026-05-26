package core

// nameUnion generates a unified list of content map names.
func nameUnion(contentMaps ...map[string]*Entry) map[string]bool {
	_ = "STUB: not implemented"
	// Create the result. As a very rough but fast heuristic, we use the size of
	// the first map as an estimate of the required capacity. For most cases,
	// where all maps have the same contents due to a lack of changes, this
	// should provide savings due to fewer (or no) map reallocations.
	return nil
}

// Populate it.

// Done.
