package ssh

import (
	"context"
	"os/exec"
)

// CompressionFlag returns a flag that can be passed to scp or ssh to enable
// compression. Note that while SSH does have a CompressionLevel configuration
// option, this only applies to SSHv1. SSHv2 defaults to a DEFLATE level of 6,
// which is what we want anyway.
func CompressionFlag() string {
	_ = "STUB: not implemented"

	// ConnectTimeoutFlag returns a flag that can be passed to scp or ssh to limit
	// connection time. The provided timeout is in seconds. The timeout must be
	// greater than 0, otherwise this function will panic.
	return ""
}

func ConnectTimeoutFlag(timeout uint64) string {
	_ = "STUB: not implemented"
	// Validate the timeout.
	return ""
}

// Format the flag.

// ServerAliveFlags returns a set of flags that can be passed to scp or ssh to
// enable use of server alive messages. The provided interval is in seconds.
// Both the interval and count must be greater than 0, otherwise this function
// will panic.
func ServerAliveFlags(interval, countMax int) []string {
	_ = "STUB: not implemented"
	// Validate the interval and count.
	return nil
}

// Format the flags.

// sshCommandPath returns the full path to use for invoking ssh. It will use the
// MUTAGEN_SSH_PATH environment variable if provided, otherwise falling back to
// a platform-specific implementation.
func sshCommandPath() (string, error) {
	_ = "STUB: not implemented"
	// If MUTAGEN_SSH_PATH is specified, then use it to perform the lookup.
	return "", nil
}

// Otherwise fall back to the platform-specific implementation.

// SSHCommand prepares (but does not start) an SSH command with the specified
// arguments and scoped to lifetime of the provided context.
func SSHCommand(ctx context.Context, args ...string) (*exec.Cmd, error) {
	_ = "STUB: not implemented"
	// Identify the command name or path.
	return nil, nil
}

// Create the command.

// scpCommandPath returns the full path to use for invoking scp. It will use the
// MUTAGEN_SSH_PATH environment variable if provided, otherwise falling back to
// a platform-specific implementation.
func scpCommandPath() (string, error) {
	_ = "STUB: not implemented"
	// If MUTAGEN_SSH_PATH is specified, then use it to perform the lookup.
	return "", nil
}

// Otherwise fall back to the platform-specific implementation.

// SCPCommand prepares (but does not start) an SCP command with the specified
// arguments and scoped to lifetime of the provided context.
func SCPCommand(ctx context.Context, args ...string) (*exec.Cmd, error) {
	_ = "STUB: not implemented"
	// Identify the command name or path.
	return nil, nil
}

// Create the command.
