package core

// synchronizable returns true if the entry kind is synchronizable and false if
// the entry kind is unsynchronizable.
func (k EntryKind) synchronizable() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (k EntryKind) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (k *EntryKind) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a forwarding status.

// Success.

// EnsureValid ensures that Entry's invariants are respected. If synchronizable
// is true, then unsynchronizable content will be considered invalid.
func (e *Entry) EnsureValid(synchronizable bool) error {
	_ = "STUB: not implemented"
	// A nil entry represents an absence of content and is therefore valid.
	return nil
}

// Otherwise validate based on kind.

// Ensure that no invalid fields are set.

// Validate contents. Nil entries are not considered valid for contents.

// Ensure that no invalid fields are set.

// Ensure that the digest is non-empty.

// Ensure that no invalid fields are set.

// Ensure that the target is non-empty. We avoid any further validation
// because there's none that we can reasonably perform.

// Verify that unsynchronizable content is allowed.

// Ensure that no invalid fields are set.

// Verify that unsynchronizable content is allowed.

// Ensure that no invalid fields are set.

// Ensure that the problem is non-empty.

// Verify that unsynchronizable content is allowed.

// Ensure that no invalid fields are set.

// Validate contents. Nil entries are not considered valid for contents.

// Success.

// entryVisitor is a callback type used for Entry.walk.
type entryVisitor func(path string, entry *Entry)

// walk performs a depth-first traversal of the entry, invoking the specified
// visitor on each element in the entry hierarchy. The path argument specifies
// the path at which the root entry should be treated as residing. If reverse is
// false, then each entry will be visited before its contents (i.e. a normal
// depth-first traversal), otherwise it will be visited after its contents (i.e.
// a reverse depth-first traversal).
func (e *Entry) walk(path string, visitor entryVisitor, reverse bool) {
	_ = "STUB: not implemented"
	// If this is a normal walk, then visit the entry before its contents.
	return
}

// If this entry is non-nil, then visit any child entries. We don't bother
// checking if the entry is a directory since this is an internal method and
// the caller is responsible for enforcing entry invariants (meaning that
// only directories will have child entries).

// Compute the prefix to add to content names to compute their paths.

// Process the child entries.

// If this is a reverse walk, then visit the entry after its contents.

// Count returns the total number of entries within the entry hierarchy rooted
// at the entry, excluding nil and unsynchronizable entries.
func (e *Entry) Count() uint64 {
	_ = "STUB: not implemented"
	// Nil entries represent an empty hierarchy.
	return 0
}

// Unsynchronizable entries can be excluded from the count because they
// don't represent content that can or will be synchronized.

// Count ourselves.

// Count any child entries. We don't bother checking if the entry is a
// directory since the caller is responsible for enforcing entry invariants
// (meaning that only directories will have child entries).

// TODO: At the moment, we don't worry about overflow here. The
// reason is that, in order to overflow uint64, we'd need a minimum
// of 2**64 entries in the hierarchy. Even assuming that each entry
// consumed only one byte of memory (and they consume at least an
// order of magnitude more than that), we'd have to be on a system
// with (at least) ~18.5 exabytes of memory. Additionally, Protocol
// Buffers messages have even lower size limits that would prevent
// such an Entry from being sent over the network. But we should
// still fix this at some point.

// Done.

// entryEqualWildcardProblemMatch controls whether or not wildcard problem
// matching is enabled for Entry.Equal. Ideally this would be a constant so that
// the compiler could optimize away the unused branch in Entry.Equal, but
// there's no "test" build tag that we can use to redefine constants for tests
// only. The Go developers seem adamant that no such flag should be added. We
// could define one manually, but modern CPUs will chew through this additional
// check quickly enough anyway, so it's not worth the trouble.
var entryEqualWildcardProblemMatch bool

// Equal performs an equivalence comparison between this entry and another. If
// deep is true, then the comparison is performed recursively, otherwise the
// comparison is only performed between entry properties at the top level and
// content maps are ignored.
func (e *Entry) Equal(other *Entry, deep bool) bool {
	_ = "STUB: not implemented"
	// If the pointers are equal, then the entries are equal, both shallowly and
	// recursively. This includes the case where both pointers are nil, which
	// represents the absence of content. If only one pointer is nil, then they
	// can't possibly be equal.
	return false
}

// Compare all properties except for problem messages.

// Compare problem messages according to whether or not wildcard problem
// matching is enabled. We only enable this for tests, where we can't always
// know the exact problem message ahead of time due to variations between
// different operating systems. Wildcard matching means that if one or both
// of the entries has a problem message of "*", it will be considered a
// match for the other entry's problem message.

// If a deep comparison wasn't requested, then we're done.

// Compare entry contents.

// Done.

// EntryCopyBehavior indicates the type of Copy operation to perform for an
// Entry. All copy types behave the same for scalar entries - they only vary the
// behavior of directory entry copies (including phantom directories).
type EntryCopyBehavior uint8

const (
	// EntryCopyBehaviorDeep indicates that a deep copy of the entry should be
	// created.
	EntryCopyBehaviorDeep EntryCopyBehavior = iota
	// EntryCopyBehaviorDeepPreservingLeaves indicates that a deep copy of the
	// entry should be created, but that all leaf (non-directory) entry types
	// should be copied by value (i.e. by their Entry pointer) to avoid
	// allocation. This copy type can be useful if only directories in the copy
	// are going to be mutated.
	EntryCopyBehaviorDeepPreservingLeaves
	// EntryCopyBehaviorShallow indicates that a shallow copy of the entry
	// should be created.
	EntryCopyBehaviorShallow
	// EntryCopyBehaviorSlim indicates that a "slim" copy of the entry should be
	// created, which is a shallow copy that excludes the content map.
	EntryCopyBehaviorSlim
)

// Copy creates a copy of the entry using the specified copy behavior. In
// general, entries are considered immutable (by convention) and should be
// copied by pointer. However, when creating derived entries (e.g. using Apply),
// a copy operation may be necessary to create a temporarily mutable entry that
// can be modified (until returned). That is the role of this method. Although
// exported for benchmarking, there should generally be no need for code outside
// of this package to use it, except to convert a full entry to a slim entry.
func (e *Entry) Copy(behavior EntryCopyBehavior) *Entry {
	_ = "STUB: not implemented"
	// If the entry is nil, then the copy is nil.
	return nil
}

// Create a slim copy.

// If a slim copy was requested, then we're done.

// If the original entry doesn't have any contents, then return early to
// avoid allocation of the content map.

// Copy the entry contents.

// Done.

// synchronizable returns the subtree of the entry hierarchy consisting of only
// synchronizable content. It is useful for constructing the new value of a
// change when attempting to propagate around unsychronizable content. It will
// return nil if the entry itself is unsynchronizable (which is technically the
// synchronizable subtree of the entry hierarchy in that case).
func (e *Entry) synchronizable() *Entry {
	_ = "STUB: not implemented"
	// If the entry itself is nil, then the resulting subtree is nil.
	return nil
}

// If the entry itself consists of unsynchronizable content, then the
// resulting subtree is nil.

// If the entry (which we know is synchronizable) is not a directory, then
// we can just return the entry itself.

// If the entry (which we know is a directory) doesn't have any contents,
// then we can just return the entry itself.

// Create a slim copy of the entry. We only need to copy fields for
// synchronizable entry types since we know this entry is synchronizable.

// Copy the entry contents. Some may not be synchronizable, in which case we
// exclude them from the resulting map. We don't need to worry about them
// already having been nil since nil entries aren't allowed in content maps.

// Done.

// Problems generates a list of problems from the problematic entries contained
// within the entry hierarchy. The problems are returned in depth-first but
// non-deterministic order. Problem paths are computed assuming the entry
// represents the synchronization root.
func (e *Entry) Problems() []*Problem {
	_ = "STUB: not implemented"
	// Create the result.
	return nil
}

// Perform a walk to record problematic entries.

// Done.
