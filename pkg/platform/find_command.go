package platform

// FindCommand searches for a command with the specified name within the
// specified list of directories. It's similar to os/exec.LookPath, except that
// it allows one to manually specify paths, and it uses a slightly simpler
// lookup mechanism.
func FindCommand(name string, paths []string) (string, error) {
	_ = "STUB: not implemented"
	// Iterate through the directories.
	return "", nil
}

// Compute the target name.

// Check if the target exists and has the correct type.
// TODO: Should we do more extensive (and platform-specific) testing on
// the resulting metadata? See, e.g., the implementation of
// os/exec.LookPath.

// Failure.
