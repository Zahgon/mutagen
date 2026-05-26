package url

// Format formats a URL into a human-readable (and reparsable) format.
func (u *URL) Format(environmentPrefix string) string { _ = "STUB: not implemented"; return "" }

// formatLocal formats a local URL.
func (u *URL) formatLocal() string {
	_ = "STUB: not implemented"

	// formatSSH formats an SSH URL into an SCP-style URL.
	return ""
}

func (u *URL) formatSSH() string {
	_ = "STUB: not implemented"
	// Create the base result.
	return ""
}

// Add username if present.

// Add port if present.

// Add path.

// Done.

// invalidDockerURLFormat is the value returned by formatDocker when a URL is
// provided that breaks invariants.
const invalidDockerURLFormat = "<invalid-docker-url>"

// formatDocker formats a Docker URL.
func (u *URL) formatDocker(environmentPrefix string) string {
	_ = "STUB: not implemented"
	// Start with the container name.
	return ""
}

// Add username if present.

// Append the path in a manner that depends on the URL kind.

// If this is a home-directory-relative path or a Windows path, then we
// need to prepend a slash.

// Add the scheme.

// Add environment variable information if requested.

// Add parameter information, if requested.

// Done.
