package synchronization

import (
	"github.com/mutagen-io/mutagen/pkg/synchronization/core"
)

// Conflict represents a filesystem change conflict.
type Conflict struct {
	// Root is the root path for the conflict, relative to the synchronization
	// root.
	Root string `json:"root"`
	// AlphaChanges are the relevant changes on alpha.
	AlphaChanges []Change `json:"alphaChanges"`
	// BetaChanges are the relevant changes on beta.
	BetaChanges []Change `json:"betaChanges"`
}

// loadFromInternal sets a conflict to match an internal Protocol Buffers
// representation. The conflict must be valid.
func (c *Conflict) loadFromInternal(conflict *core.Conflict) {
	_ = "STUB: not implemented"
	// Propagate the conflict root.
	return
}

// Propagate alpha changes.

// Propagate beta changes.

// exportConflicts is a convenience function that calls
// Conflict.loadFromInternal for a slice of conflicts.
func exportConflicts(conflicts []*core.Conflict) []Conflict {
	_ = "STUB: not implemented"
	// If there are no conflicts, then just return a nil slice.
	return nil
}

// Create the resulting slice.

// Done.
