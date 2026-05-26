package daemon

import (
	"golang.org/x/sys/windows/registry"
)

// RegistrationSupported indicates whether or not daemon registration is
// supported on this platform.
const RegistrationSupported = true

const (
	// rootKey is the registry root for daemon registration.
	rootKey = registry.CURRENT_USER
	// runPath is the path to the registry entries for automatic startup.
	runPath = "Software\\Microsoft\\Windows\\CurrentVersion\\Run"
	// runKeyName is the key used to register Mutagen for automatic startup.
	runKeyName = "Mutagen"
)

// Register performs automatic daemon startup registration.
func Register() error {
	_ = "STUB: not implemented"
	// Attempt to open the relevant registry path and ensure it's cleaned up
	// when we're done.
	return nil
}

// Compute the path to the current executable.

// Compute the command to start the Mutagen daemon.

// Attempt to register the daemon.

// Success.

// Unregister performs automatic daemon startup de-registration.
func Unregister() error {
	_ = "STUB: not implemented"
	// Attempt to open the relevant registry path and ensure it's cleaned up
	// when we're done.
	return nil
}

// Attempt to deregister the daemon.

// Success.

// RegisteredStart potentially handles daemon start operations if the daemon is
// registered for automatic start with the system. It returns false if the start
// operation was not handled and should be handled by the normal start command.
func RegisteredStart() (bool, error) {
	_ = "STUB: not implemented"

	// RegisteredStop potentially handles stop start operations if the daemon is
	// registered for automatic start with the system. It returns false if the stop
	// operation was not handled and should be handled by the normal stop command.
	return false, nil
}

func RegisteredStop() (bool, error) { _ = "STUB: not implemented"; return false, nil }
