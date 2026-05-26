package sidecar

import (
	"sync"
)

// environmentIsSidecar is the cached result of the sidecar environment check.
var environmentIsSidecar bool

// checkEnvironmentOnce gates access to environmentIsSidecar.
var checkEnvironmentOnce sync.Once

// EnvironmentIsSidecar returns true if the current operating environment is a
// Mutagen sidecar container.
func EnvironmentIsSidecar() bool { _ = "STUB: not implemented"; return false }
