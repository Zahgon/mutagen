package url

// dockerURLPrefix is the lowercase version of the Docker URL prefix.
const dockerURLPrefix = "docker://"

// DockerEnvironmentVariables is a list of Docker environment variables that
// should be locked in to Docker URLs at parse time.
var DockerEnvironmentVariables = []string{
	"DOCKER_HOST",
	"DOCKER_TLS", // This flag is not documented, but is supported by Docker.
	"DOCKER_TLS_VERIFY",
	"DOCKER_CERT_PATH",
	"DOCKER_CONTEXT",
	"DOCKER_CONFIG", // This flag is not documented, but is supported by Docker.
	"DOCKER_API_VERSION",
}

// dockerParameterNames is a list of supported Docker command line parameters.
var dockerParameterNames = []string{
	"config",
	"context",
	"host",
	"tls",
	"tlscacert",
	"tlscert",
	"tlskey",
	"tlsverify",
}

// isDockerURL checks whether or not a URL is a Docker URL. It requires the
// presence of a Docker protocol prefix.
func isDockerURL(raw string) bool { _ = "STUB: not implemented"; return false }

// parseDocker parses a Docker URL.
func parseDocker(raw string, kind Kind, first bool) (*URL, error) {
	_ = "STUB: not implemented"
	// Strip off the prefix.
	return nil, nil
}

// Determine the character that splits the container name from the path or
// forwarding endpoint component.

// Parse off the username. If we hit a '/', then we've reached the end of a
// container specification and there was no username. Similarly, if we hit
// the end of the string without seeing an '@', then there's also no
// username specified. Ideally we'd want also to break on any character that
// isn't allowed in a username, but that isn't well-defined, even for POSIX
// (it's effectively determined by a configurable regular expression -
// NAME_REGEX).

// Split what remains into the container and the path (or forwarding
// endpoint, depending on the URL kind). Ideally we'd want to be a bit more
// stringent here about what characters we accept in container names,
// potentially breaking early with an error if we see a "disallowed"
// character, but we're better off just allowing Docker to reject container
// names that it doesn't like.

// Perform path processing based on URL kind.

// If the path starts with "/~", then we assume that it's supposed to be
// a home-directory-relative path and remove the slash. At this point we
// already know that the path starts with "/" since we retained that as
// part of the path in the split operation above.

// If the path is of the form "/" + Windows path, then assume it's
// supposed to be a Windows path. This is a heuristic, but a reasonable
// one. We do this on all systems (not just on Windows as with SSH URLs)
// because users can connect to Windows containers from non-Windows
// systems. At this point we already know that the path starts with "/"
// since we retained that as part of the path in the split operation
// above.

// For forwarding paths, we need to trim the split character at the
// beginning.

// Parse the forwarding endpoint URL to ensure that it's valid.

// Store any Docker environment variables that we need to preserve. We only
// store variables that are actually present, because Docker behavior will
// vary depending on whether a variable is unset vs. set but empty.

// Success.
