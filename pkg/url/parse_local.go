package url

// parseLocal parses a local URL. It simply assumes the URL refers to a local
// path or forwarding endpoint specification.
func parseLocal(raw string, kind Kind) (*URL, error) {
	_ = "STUB: not implemented"
	// If this is a synchronization URL, then ensure that its path is
	// normalized.
	return nil, nil
}

// If this is a forwarding URL, then parse it to ensure that it's valid. If
// it's a Unix domain socket endpoint, then ensure that the socket path is
// normalized.

// Perform parsing.

// Normalize and reformat the endpoint URL if necessary.

// Create the URL.
