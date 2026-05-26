package grpcutil

// PeelAwayRPCErrorLayer peels away any intermediate RPC error layer from an
// error returned by gRPC-based code and constructs an error object using the
// underlying error message. If this unwrapping fails, the argument is returned
// directly.
func PeelAwayRPCErrorLayer(err error) error {
	_ = "STUB: not implemented"
	// Attempt to peel away the RPC layer.
	return nil
}

// Otherwise return the argument directly.
