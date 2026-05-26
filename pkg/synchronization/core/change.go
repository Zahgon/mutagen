package core

// EnsureValid ensures that Change's invariants are respected. If synchronizable
// is true, then unsynchronizable content in either the Old or New field will be
// considered invalid.
func (c *Change) EnsureValid(synchronizable bool) error {
	_ = "STUB: not implemented"
	// A nil change is not valid.
	return nil
}

// Technically we could validate the path, but that's error-prone,
// expensive, and not really needed for memory safety. We also can't enforce
// that the old entry value is not equal to the new entry value because for
// the "synthetic" changes generated in unidirectional synchronization they
// may be identical.

// Validate entries.

// Success.

// slim creates a "slim" copy of the Change object, where both entries are slim
// copies with contents excluded.
func (c *Change) slim() *Change { _ = "STUB: not implemented"; return nil }

// IsRootDeletion indicates whether or not the change represents a root
// deletion.
func (c *Change) IsRootDeletion() bool { _ = "STUB: not implemented"; return false }

// IsRootTypeChange indicates whether or not the change represents a root type
// change.
func (c *Change) IsRootTypeChange() bool { _ = "STUB: not implemented"; return false }
