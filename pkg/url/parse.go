package url

// Parse parses a raw URL string into a URL type. It accepts information about
// the URL kind (e.g. synchronization vs. forwarding) and position (i.e. the URL
// is considered an alpha/source URL if first is true and a beta/destination URL
// otherwise).
func Parse(raw string, kind Kind, first bool) (*URL, error) {
	_ = "STUB: not implemented"
	// Ensure that the kind is supported.
	return nil, nil
}

// Don't allow empty raw URLs.

// Dispatch URL parsing based on type. We have to be careful about the
// ordering here because URLs may be classified as multiple types (e.g. a
// Docker URL would also be classified as an SCP-style SSH URL), but we only
// want them to be parsed according to the better and more specific match.
// If we don't match anything, we assume the URL is a local path.
