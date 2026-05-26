package core

// propagateExecutabilityRecursive propagates executability recursively.
func propagateExecutabilityRecursive(ancestor, source, target *Entry) {
	_ = "STUB: not implemented"
	// If there is no location from which executability information can be
	// propagated or no location to which executability information can be
	// propagated, then we can discontinue recursion along this path.
	return
}

// Handle based on target kind.

// If this is a directory, then grab the contents of the target, source,
// and ancestor.

// If both the source and ancestor are empty, then we can terminate
// recursion here, because there won't be anything to propagate at lower
// levels. The same is true of target is empty, but that's implicitly
// checked below since that's what we loop over.

// Loop over the target contents and recursively propagate
// executability.

// If this is a file, then we use a series of heuristics to perform the
// correct propagation.

// If the source is also a file with the same contents as the target,
// then we can assume that they should have the same executability
// setting. Since the target isn't capable of modifying executability
// settings, we know that we're not overwriting or ignoring such a
// modification from the target side. This handles the most frequent
// case that occurs for files during synchronization cycles: both sides
// being unmodified. It also handles cases of both-modified-same
// behavior, and in that sense this particular check is a heuristic,
// because it assumes that the creator on the non-preserving side would
// want the file to have the same executability setting as on the
// preserving side (a reasonable assumption since the files were created
// at the same time with the same contents). Fortunately, the assumption
// behind this heuristic is irrelevant, since there is no executability
// setting that will be propagated to and visible on the non-preserving
// side. In that sense, the assumption here is only used to make the
// three-way merge proceed smoothly. Finally, this check handles the
// case where a session is being started between preserving and
// non-preserving endpoints where both endpoints have a copy of the
// files already (which is actually also just a special case of
// both-modified-same behavior without an ancestor). In this case, we
// also assume that the non-preserving side should have the same
// executability bits as the preserving side for files at the same path
// with the same contents, and again this assumption has no visible
// effects - it is only to make the merging proceed smoothly.

// If the source and target differ, then we look to the ancestor. If the
// ancestor is also a file with the same contents as the target, then we
// propagate executability from the ancestor. This check handles cases
// where the preserving side has been modified but the non-preserving
// side is unmodified. This is safe even in the case where the
// preserving side has changed its executability setting, because in
// that case the non-preserving side will appear completely unmodified
// after this propagation and be overwritten.

// If the target contents differ from both the ancestor and the source,
// then there has been a modification to the contents of the file on the
// target side. In this case, we check if the source is also a file that
// is unmodified from the ancestor. If that's the case, then we
// propagate executability settings from the source. This is necessary
// to handle cases where (e.g.) an executable file has been modified on
// the non-preserving side (e.g. editing a shell script on Windows with
// an executable bit on the other side of the connection that you want
// to preserve). This is *definitely* a heuristic, because it assumes
// that you'd want to preserve an executability bit even when completely
// replacing the contents of a file. This might not make sense in some
// cases, e.g. where you completely replace a shell script with an image
// file on the non-preserving side. However, this is unlikely to occur
// in practice, and even if it does, it's a one-time operation to then
// strip the executability bit on the preserving side. In contrast, if
// we didn't have this heuristic, we'd end up having to re-mark a file
// as executable on the preserving side *every* time we edited it on the
// non-preserving side.

// We intentially avoid propagating executability bits in the case that
// both the preserving and non-preserving side have modified file
// contents. There is no heuristic that consistently makes sense for
// even a small fraction of such cases.

// PropagateExecutability propagates file executability from the ancestor and
// source to the target in a recursive fashion. Executability information is
// only propagated if entry paths, types, and contents match, with source taking
// precedent over ancestor.
func PropagateExecutability(ancestor, source, target *Entry) *Entry {
	_ = "STUB: not implemented"
	// Create a copy of the snapshot that we can mutate.
	return nil
}

// Perform propagation.

// Done.
