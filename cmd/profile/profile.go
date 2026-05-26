package profile

import (
	"os"
)

// Profile manages a CPU and heap profile.
type Profile struct {
	// name is the name of the profile.
	name string
	// cpuProfile is the output file for the CPU profile.
	cpuProfile *os.File
}

// New creates a new profile instance. The profiling begins immediately.
func New(name string) (*Profile, error) {
	_ = "STUB: not implemented"
	// Open the CPU profile output.
	return nil, nil
}

// Start CPU profiling.

// Success.

// Finalize terminates a profile and writes its measurements to disk in the
// current working directory.
func (p *Profile) Finalize() error {
	_ = "STUB: not implemented"
	// Close out the CPU profile.
	return nil
}

// Run a GC cycle to update the heap profile statistics.

// Write a heap profile.

// Success.
