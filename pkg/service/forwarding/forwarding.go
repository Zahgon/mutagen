package forwarding

// ensureValid verifies that a CreationSpecification is valid.
func (s *CreationSpecification) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil creation specification is not valid.
	return nil
}

// Verify that the source URL is valid and is a forwarding URL.

// Verify that the destination URL is valid and is a forwarding URL.

// Verify that the configuration is valid.

// Verify that the source-specific configuration is valid.

// Verify that the destination-specific configuration is valid.

// Verify that the name is valid.

// Verify that labels are valid.

// There's no need to validate the Paused field - either value is valid.

// Success.

// ensureValid verifies that a CreateRequest is valid.
func (r *CreateRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil create request is not valid.
	return nil
}

// Ensure that a prompter has been specified.

// Ensure that the creation specification is valid.

// Success.

// EnsureValid verifies that a CreateResponse is valid.
func (r *CreateResponse) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil create response is not valid.
	return nil
}

// Ensure that the session identifier is non-empty.

// Success.

// ensureValid verifies that a ListRequest is valid.
func (r *ListRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil list request is not valid.
	return nil
}

// Ensure that the session selection is valid.

// There's no need to validate the state index - any value is valid.

// Success.

// EnsureValid verifies that a ListResponse is valid.
func (r *ListResponse) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil list response is not valid.
	return nil
}

// Ensure that all states are valid.

// Success.

// ensureValid verifies that a PauseRequest is valid.
func (r *PauseRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil pause request is not valid.
	return nil
}

// Ensure that a prompter has been specified.

// Ensure that the session selection is valid.

// Success.

// EnsureValid verifies that a PauseResponse is valid.
func (r *PauseResponse) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil pause response is not valid.
	return nil
}

// Success.

// ensureValid verifies that a ResumeRequest is valid.
func (r *ResumeRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil resume request is not valid.
	return nil
}

// Ensure that a prompter has been specified.

// Ensure that the session selection is valid.

// Success.

// EnsureValid verifies that a ResumeResponse is valid.
func (r *ResumeResponse) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil resume response is not valid.
	return nil
}

// Success.

// ensureValid verifies that a TerminateRequest is valid.
func (r *TerminateRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil terminate request is not valid.
	return nil
}

// Ensure that a prompter has been specified.

// Ensure that the session selection is valid.

// Success.

// EnsureValid verifies that a TerminateResponse is valid.
func (r *TerminateResponse) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil terminate response is not valid.
	return nil
}

// Success.
