//go:build !windows && !darwin

package daemon

// RegistrationSupported indicates whether or not daemon registration is
// supported on this platform.
const RegistrationSupported = false

// Register performs automatic daemon startup registration.
func Register() error { _ = "STUB: not implemented"; return nil }

// Unregister performs automatic daemon startup de-registration.
func Unregister() error { _ = "STUB: not implemented"; return nil }

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
