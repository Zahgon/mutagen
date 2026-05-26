package selection

// EnsureNameValid ensures that a name is valid for use as a session name. Empty
// names are treated as valid.
func EnsureNameValid(name string) error {
	_ = "STUB: not implemented"
	// Loop over the string and ensure that its characters are allowed. We allow
	// letters, numbers, and dashses, but we require that the identifier starts
	// with a letter. We disallow underscores to avoid colliding with internal
	// identifiers. If a name contains dashes, then we enforce that it isn't a
	// UUID to avoid collisions with legacy session identifiers.
	return nil
}

// If the name contains a dash, then ensure that it isn't a UUID.

// Disallow "defaults" as a name since it is used as a special key in YAML
// files.

// Success.
