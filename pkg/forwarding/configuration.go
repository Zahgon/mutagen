package forwarding

// EnsureValid ensures that Configuration's invariants are respected. The
// validation of the configuration depends on whether or not it is
// endpoint-specific.
func (c *Configuration) EnsureValid(endpointSpecific bool) error {
	_ = "STUB: not implemented"
	// A nil configuration is not considered valid.
	return nil
}

// Verify that the socket overwrite mode is unspecified or supported.

// Verify the socket owner specification.

// Verify the socket group specification.

// We don't verify the socket permission mode because there's not really any
// way to know if it's a sane value.

// Success.

// Equal returns whether or not the configuration is equivalent to another. The
// result of this method is only valid if both configurations are valid.
func (c *Configuration) Equal(other *Configuration) bool {
	_ = "STUB: not implemented"
	// Ensure that both are non-nil.
	return false
}

// Perform an equivalence check.

// MergeConfigurations merges two configurations of differing priorities. Both
// configurations must be non-nil.
func MergeConfigurations(lower, higher *Configuration) *Configuration {
	_ = "STUB: not implemented"
	// Create the resulting configuration.
	return nil
}

// Merge the socket overwrite mode.

// Merge the socket owner.

// Merge the socket group.

// Merge the socket permission mode.

// Done.
