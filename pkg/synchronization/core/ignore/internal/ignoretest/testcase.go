package ignoretest

import (
	"testing"

	"github.com/mutagen-io/mutagen/pkg/synchronization/core/ignore"
)

// ignoreStatusDescription returns a human-readable ignore status description.
func ignoreStatusDescription(status ignore.IgnoreStatus) string {
	_ = "STUB: not implemented"
	return ""
}

// TestValue encodes a test operation in an TestCase.
type TestValue struct {
	// Path is the path to test.
	Path string
	// Directory indicates whether or not the path is a directory.
	Directory bool
	// ExpectedStatus is the expected ignore status.
	ExpectedStatus ignore.IgnoreStatus
	// ExpectedContinueTraversal is the expected traversal continuation status.
	ExpectedContinueTraversal bool
}

// TestCase encodes a sequence of test values for a specified set of ignore
// patterns.
type TestCase struct {
	// PatternValidator is the pattern validation callback.
	PatternValidator func(string) error
	// Constructor is the ignorer constructor callback.
	Constructor func([]string) (ignore.Ignorer, error)
	// Ignores are the ignore patterns.
	Ignores []string
	// Tests are the ignore tests to run.
	Tests []TestValue
}

// Run invokes the test with the specified test runner.
func (c *TestCase) Run(t *testing.T) {
	_ = "STUB: not implemented"
	// Mark this runner as a helper.
	return
}

// Ensure that all patterns are valid.

// Create the ignorer.

// Verify test values.
