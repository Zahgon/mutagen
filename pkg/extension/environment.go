package extension

import (
	"sync"
)

// environmentIsExtension is the cached result of the extension environment
// check.
var environmentIsExtension bool

// checkEnvironmentOnce gates access to environmentIsExtension.
var checkEnvironmentOnce sync.Once

// EnvironmentIsExtension returns true if the current operating environment is
// the Mutagen Extension for Docker Desktop service container.
func EnvironmentIsExtension() bool { _ = "STUB: not implemented"; return false }
