package agent

import (
	"github.com/mutagen-io/mutagen/pkg/logging"
)

// Install installs the current binary to the appropriate location for an agent
// binary with the current Mutagen version.
func Install() error {
	_ = "STUB: not implemented"
	// Compute the destination.
	return nil
}

// Compute the path to the current executable.

// Relocate the current executable to the installation path.

// Success.

// install attempts to probe an endpoint and install the appropriate agent
// binary over the specified transport.
func install(logger *logging.Logger, transport Transport, prompter string) error {
	_ = "STUB: not implemented"
	// Detect the target platform.
	return nil
}

// Find the appropriate agent binary. Ensure that it's cleaned up when we're
// done with it.

// Copy the agent to the remote. We use a unique identifier for the
// temporary destination. For Windows remotes, we add a ".exe" suffix, which
// will automatically make the file executable on the remote (POSIX systems
// are handled separately below). For POSIX systems, we add a dot prefix to
// hide the executable.

// For cases where we're copying from a Windows system to a POSIX remote,
// invoke "chmod +x" to add executability back to the copied binary. This is
// necessary under the specified circumstances because as soon as the agent
// binary is extracted from the bundle, it will lose its executability bit
// since Windows can't preserve this. This will also be applied to Windows
// POSIX remotes, but a "chmod +x" there will just be a no-op.

// Invoke the remote installation.

// Success.
