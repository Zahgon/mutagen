package local

const (
	// alphaName is the name to use for alpha when distinguishing endpoints.
	alphaName = "alpha"
	// betaName is the name to use for beta when distinguishing endpoints.
	betaName = "beta"

	// stagingPrefixLength is the byte length to use for prefix directories when
	// load-balancing staged files.
	stagingPrefixLength = 1
)

// pathForCache computes the path to the serialized cache for the given session
// identifier and endpoint role.
func pathForCache(session string, alpha bool) (string, error) {
	_ = "STUB: not implemented"
	// Compute/create the caches directory.
	return "", nil
}

// Compute the endpoint name.

// Compute the cache name.

// Success.

// pathForMutagenStagingRoot computes the path to the staging root in the
// Mutagen data directory for the given session identifier and endpoint. It
// ensures that staging subdirectory of the Mutagen data directory exists, but
// it does not create the staging root itself.
func pathForMutagenStagingRoot(session string, alpha bool) (string, error) {
	_ = "STUB: not implemented"
	// Compute the path to the staging root parent (the global Mutagen data
	// directory in which staging roots are stored) and ensure that it exists.
	return "", nil
}

// Compute the endpoint name.

// Compute the staging root name.

// Compute the combined path.

// pathForNeighboringStagingRoot computes the path to the staging root which
// neighbors the synchronization root for the given root, session identifier,
// and endpoint. It does not create the directory or any parent directories.
func pathForNeighboringStagingRoot(root, session string, alpha bool) (string, error) {
	_ = "STUB: not implemented"
	// Compute the parent of the staging root.
	return "", nil
}

// Compute the endpoint name.

// Compute the name of the staging directory.

// Compute the path to the staging root.

// pathForInternalStagingRoot computes the path to the staging root which is
// internal to the synchronization root for the given root, session identifier,
// and endpoint. It does not create the directory or any parent directories.
func pathForInternalStagingRoot(root, session string, alpha bool) (string, error) {
	_ = "STUB: not implemented"
	// Compute the endpoint name.
	return "", nil
}

// Compute the name of the staging directory.

// Compute the path to the staging root.
