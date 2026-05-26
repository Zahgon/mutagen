package core

// EnsureValid ensures that Conflict's invariants are respected.
func (c *Conflict) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil conflict is not valid.
	return nil
}

// There's not much validation we can perform on the conflict root path. It
// may be empty and any format validation would be limited.

// Each side's changes must be non-empty and must all be valid. We always
// allow unsynchronizable content in conflict changes because they can be
// the source of the conflict.

// There's technically a bit more validation we could do, but it would be
// expensive and wouldn't be exhaustive in any case. The purpose of this
// function is simply to enforce memory safety invariants and algorithmic
// invariants. Memory safety is fully verified by the checks performed here,
// and conflicts don't enter into the synchronization algorithm (they're
// purely a byproduct of it), so they can't corrupt a synchronization
// session (and they come from a trusted source anyway - the daemon's
// reconciliation algorithm).

// Success.

// Slim returns a copy of the conflict where each Change object has had its root
// entry reduced to a shallow copy (i.e. excluding contents). The conflict will
// still have enough metadata to determine its root path, and it will be
// considered valid.
func (c *Conflict) Slim() *Conflict {
	_ = "STUB: not implemented"
	// Recompute alpha changes.
	return nil
}

// Recompute beta changes.

// Done.

// CopyConflicts creates a copy of a list of conflicts in a new slice, usually
// for the purpose of modifying the list. The conflict objects themselves are
// not copied. It preserves nil vs. non-nil characteristics for empty slices.
func CopyConflicts(conflicts []*Conflict) []*Conflict {
	_ = "STUB: not implemented"
	// If the slice is nil, then preserve its nilness. For zero-length, non-nil
	// slices, we still allocate on the heap to preserve non-nilness.
	return nil
}

// Make a copy.

// Done.

// sortableConflictList implements sort.Interface for conflict lists.
type sortableConflictList []*Conflict

// Len implements sort.Interface.Len.
func (l sortableConflictList) Len() int {
	_ = "STUB: not implemented"

	// Less implements sort.Interface.Less.
	return 0
}

func (l sortableConflictList) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface.Swap.
func (l sortableConflictList) Swap(i, j int) { _ = "STUB: not implemented"; return }

// SortConflicts sorts a list of conflicts based on their root conflict paths.
func SortConflicts(conflicts []*Conflict) { _ = "STUB: not implemented"; return }
