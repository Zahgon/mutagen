package synchronization

import (
	"github.com/mutagen-io/mutagen/pkg/synchronization/core"
)

// Problem represents a synchronization problem.
type Problem struct {
	// Path is the path at which the problem occurred, relative to the
	// synchronization root.
	Path string `json:"path"`
	// Error is a human-readable summary of the problem.
	Error string `json:"error"`
}

// loadFromInternal sets a problem to match an internal Protocol Buffers
// representation. The problem must be valid.
func (p *Problem) loadFromInternal(problem *core.Problem) { _ = "STUB: not implemented"; return }

// exportProblems is a convenience function that calls Problem.loadFromInternal
// for a slice of problems.
func exportProblems(problems []*core.Problem) []Problem {
	_ = "STUB: not implemented"
	// If there are no problems, then just return a nil slice.
	return nil
}

// Create the resulting slice.

// Done.
