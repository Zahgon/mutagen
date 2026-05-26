package selection

// EnsureValid verifies that a Selection is valid.
func (s *Selection) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil selection is not valid.
	return nil
}

// Count the number of selection mechanisms present.

// Enforce that exactly one selection mechanism is present.

// Enforce that specifications are non-empty.

// We avoid validating the label selector, if present, because it doesn't
// pose a risk to parse unvalidated and it would only be possible to
// validate by parsing, so we'll catch any format errors later.

// Success.
