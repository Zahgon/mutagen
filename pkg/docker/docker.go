package docker

import (
	"context"
	"os/exec"
)

// CommandPath returns the absolute path specification to use for invoking
// Docker. It will use the MUTAGEN_DOCKER_PATH environment variable if provided,
// otherwise falling back to a platform-specific implementation.
func CommandPath() (string, error) {
	_ = "STUB: not implemented"
	// If MUTAGEN_DOCKER_PATH is specified, then use it to perform the lookup.
	return "", nil
}

// Otherwise fall back to the platform-specific implementation.

// Command prepares (but does not start) a Docker command with the specified
// arguments and scoped to lifetime of the provided context.
func Command(ctx context.Context, args ...string) (*exec.Cmd, error) {
	_ = "STUB: not implemented"
	// Identify the command path.
	return nil, nil
}

// Create the command.
