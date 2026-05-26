package daemon

const (
	// lockName is the name of the daemon lock. It resides within the daemon
	// subdirectory of the Mutagen directory.
	lockName = "daemon.lock"
	// endpointName is the name of the daemon IPC endpoint. It resides within
	// the daemon subdirectory of the Mutagen directory.
	endpointName = "daemon.sock"
)

// subpath computes a subpath of the daemon subdirectory, creating the daemon
// subdirectory in the process.
func subpath(name string) (string, error) {
	_ = "STUB: not implemented"
	// Compute the daemon root directory path and ensure it exists.
	return "", nil
}

// Compute the combined path.

// lockPath computes the path to the daemon lock, creating any intermediate
// directories as necessary.
func lockPath() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// EndpointPath computes the path to the daemon IPC endpoint, creating any
		// intermediate directories as necessary.
		nil
}

func EndpointPath() (string, error) { _ = "STUB: not implemented"; return "", nil }
