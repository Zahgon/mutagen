//go:build !windows

package ssh

// sshCommandPathForPlatform searches for the ssh command in the user's path.
func sshCommandPathForPlatform() (string, error) {
	_ = "STUB: not implemented"
	return "",

		// scpCommandPathForPlatform searches for the scp command in the user's path.
		nil
}

func scpCommandPathForPlatform() (string, error) { _ = "STUB: not implemented"; return "", nil }
