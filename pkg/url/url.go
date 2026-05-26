package url

// Supported returns whether or not a URL kind is supported.
func (k Kind) Supported() bool { _ = "STUB: not implemented"; return false }

// MarshalText implements encoding.TextMarshaler.MarshalText.
func (p Protocol) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements encoding.TextUnmarshaler.UnmarshalText.
func (p *Protocol) UnmarshalText(textBytes []byte) error {
	_ = "STUB: not implemented"
	// Convert the bytes to a string.
	return nil
}

// Convert to a protocol.

// Success.

// EnsureValid ensures that URL's invariants are respected.
func (u *URL) EnsureValid() error {
	_ = "STUB: not implemented"
	// Ensure that the URL is non-nil.
	return nil
}

// Ensure that the kind is supported.

// Validate the User, Host, Port, and Environment components based on
// protocol.

// In the case of Docker, we intentionally avoid validating environment
// variables since the values used could change over time. Since we
// default to empty values for unspecified environment variables, this
// works out fine, at least so long as Docker continues to treat empty
// environment variables the same as unspecified ones.

// Validate the path component depending on the URL kind.

// Ensure the path is non-empty.

// If this is a local URL, then ensure that the path is absolute.
//
// HACK: The Mutagen Extension for Docker Desktop needs to avoid this
// check because of its internal faux-local URLs. In particular, Windows
// absolute paths will appear as non-absolute in the extension's
// Linux-based backend container. Technically we only need to avoid this
// check for alpha URLs, but since we control all of the extension's
// URLs, and since we don't know which endpoint this URL is targeting,
// we just disable it entirely when running in the extension.

// If this is a Docker URL, we can actually do a bit of additional
// validation.

// Parse the forwarding endpoint URL to ensure that it's valid.

// If this is a local URL and represents a Unix domain socket endpoint,
// then ensure that the socket path is absolute.

// TODO: It would be nice to perform some sort of validation on Windows
// named pipe addresses, but there's not much we can do because the
// allowed formats vary between source and destination endpoints (so
// we'd have to weave that information through this function). The only
// difference is that the ServerName component (see the link below) must
// be "." for source endpoints but can also name a remote server in the
// case of destination endpoints. But that's not really the biggest
// issue. The problem is that the name specification is kind of vague.
// It says that the PipeName component (again, see the link below) "can
// include any character other than a backslash, including numbers and
// special characters", but it doesn't mention whitespace characters
// (for example a newline character), which, as far as I'm aware, are
// not allowed. It also limits the "entire pipe name string" to 256
// characters, but it's not clear if this refers to the PipeName
// component or the entire address. Finding an appropriate matcher for
// possible server names is also an uphill battle. This might be
// specified in the UNC specification. In the end though, we're probably
// just better off letting the OS decide what to accept and simply
// returning its errors directly. For further reading, see:
// https://docs.microsoft.com/en-us/windows/win32/ipc/pipe-names

// Success.

// Equal returns whether or not the URL is equivalent to another. The result of
// this method is only valid if both URLs are valid.
func (u *URL) Equal(other *URL) bool {
	_ = "STUB: not implemented"
	// Ensure that both are non-nil.
	return false
}

// Perform an equivalence check.
