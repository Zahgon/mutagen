package core

// EnsureValid ensures that Problem's invariants are respected.
func (p *Problem) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil problem is not valid.
	return nil
}

// Ensure that an error message has been provided.

// Success.

// CopyProblems creates a copy of a list of problems in a new slice, usually for
// the purpose of modifying the list. The problem objects themselves are not
// copied. It preserves nil vs. non-nil characteristics for empty slices.
func CopyProblems(problems []*Problem) []*Problem {
	_ = "STUB: not implemented"
	// If the slice is nil, then preserve its nilness. For zero-length, non-nil
	// slices, we still allocate on the heap to preserve non-nilness.
	return nil
}

// Make a copy.

// Done.

// sortableProblemList implements sort.Interface for problem lists.
type sortableProblemList []*Problem

// Len implements sort.Interface.Len.
func (l sortableProblemList) Len() int {
	_ = "STUB: not implemented"

	// Less implements sort.Interface.Less.
	return 0
}

func (l sortableProblemList) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap implements sort.Interface.Swap.
func (l sortableProblemList) Swap(i, j int) { _ = "STUB: not implemented"; return }

// SortProblems sorts a list of conflicts based on their problem paths.
func SortProblems(conflicts []*Problem) { _ = "STUB: not implemented"; return }
