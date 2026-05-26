package rsync

// resetToZeroMaintainingCapacity resets a Transmission to its zero value, with
// the exception that it will leave the Operation member allocated if it's
// already set will simply call resetToZeroMaintainingCapacity on the Operation.
// This allows some decoders to re-use the Operation data slice capacity when
// decoding.
func (t *Transmission) resetToZeroMaintainingCapacity() {
	_ = "STUB: not implemented"
	// Reset the expected file size.
	return
}

// Reset the Done parameter.

// Reset the operation to its zero value if non-nil.

// Reset the error parameter.

// EnsureValid ensures that the Transmission's invariants are respected.
func (t *Transmission) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil transmission is not valid.
	return nil
}

// Handle validation based on whether or not the transmission is marked as
// being the end of a file.

// Success.
