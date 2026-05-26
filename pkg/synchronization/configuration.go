package synchronization

// EnsureValid ensures that Configuration's invariants are respected. The
// validation of the configuration depends on whether or not it is
// endpoint-specific.
func (c *Configuration) EnsureValid(endpointSpecific bool) error {
	_ = "STUB: not implemented"
	// A nil configuration is not considered valid.
	return nil
}

// Validate the synchronization mode.

// Verify that the hashing algorithm is unspecified or supported.

// The maximum entry count doesn't need to be validated - any of its values
// are technically valid regardless of the source.

// The maximum staging file size doesn't need to be validated - any of its
// values are technically valid regardless of the source.

// Verify that the probe mode is unspecified or supported.

// Verify that the scan mode is unspecified or supported.

// Verify that the staging mode is unspecified or supported.

// Verify that the symbolic link mode is unspecified or supported.

// Verify that the watch mode is unspecified or supported.

// The watch polling interval doesn't need to be validated - any of its
// values are technically valid regardless of the source.

// Verify that the ignore syntax is unspecified or supported.

// Verify that default ignores are unset for endpoint-specific
// configurations. This field is deprecated, but existing sessions may have
// it set, in which case we'll just prepend it to the nominal list of
// ignores when running the session. We don't bother rejecting its presence
// based on source. This is the only meaningful validation that we can do
// here because we don't yet know the ignore syntax being used (even if it's
// not a default value, it could be overridden in a different configuration
// that will be merged on top of this one). These ignores will eventually be
// validated at endpoint initialization time, but there's no convenient way
// to do it earlier in the session creation or loading process.

// Verify that ignores are unset for endpoint-specific configurations. This
// is the only meaningful validation that we can do here because we don't
// yet know the ignore syntax being used (even if it's not a default value,
// it could be overridden in a different configuration that will be merged
// on top of this one). These ignores will eventually be validated at
// endpoint initialization time, but there's no convenient way to do it
// earlier in the session creation or loading process.

// Verify that the VCS ignore mode is unspecified or supported.

// Verify that the permissions mode is unspecified or supported. Also
// determine the effective permissions mode for validating file and
// directory modes.

// HACK: We don't have a reference to the session version in this
// method, so we compute the default permissions mode by using the
// default session version for the current version of Mutagen. For
// more information on the reasoning behind this, see the note in
// Version.DefaultPermissionsMode.

// Verify that the default file mode is valid for the effective permissions
// mode.

// Verify that the default directory mode is valid for the effective
// permissions mode.

// Verify the default owner specification.

// Verify the default group specification.

// Verify that the compression algorithm is unspecified or supported.

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

// Merge the synchronization mode.

// Merge the hashing algorithm.

// Merge the maximum entry count.

// Merge the maximum staging file size.

// Merge the probing mode.

// Merge the scanning mode.

// Merge the staging mode.

// Merge the symbolic link mode.

// Merge the watching mode.

// Merge the polling interval.

// Merge the ignore syntax.

// Merge default ignores. In theory, at most one of these should be
// non-empty, but we'll still implement it as if they both might have
// content.

// Merge ignores.

// Merge the VCS ignore mode.

// Merge the permissions mode.

// Merge the default file mode.

// Merge the default directory mode.

// Merge the default owner.

// Merge the default group.

// Merge the compression algorithm.

// Done.
