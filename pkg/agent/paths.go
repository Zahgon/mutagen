package agent

const (
	// BaseName is the base name for agent executables (sans any
	// platform-specific suffix like ".exe").
	BaseName = "mutagen-agent"
)

// installPath computes and creates the parent directories of the path where the
// current executable should be installed if it is an agent binary with the
// current Mutagen version.
func installPath() (string, error) {
	_ = "STUB: not implemented"
	// Compute (and create) the path to the agent parent directory.
	return "", nil
}

// Compute the target executable name.

// Compute the installation path.
