package prompting

// hostRequestMode indicates the mode for a HostRequest.
type hostRequestMode uint8

const (
	// hostRequestModeInitial represents an initial request.
	hostRequestModeInitial hostRequestMode = iota
	// hostRequestModeMessageResponse indicates a response to a message.
	hostRequestModeMessageResponse
	// hostRequestModePromptResponse indicates a response to a prompt.
	hostRequestModePromptResponse
)

// ensureValid verifies that a HostRequest is valid.
func (r *HostRequest) ensureValid(mode hostRequestMode) error {
	_ = "STUB: not implemented"
	// A nil hosting request is not valid.
	return nil
}

// Handle validation based on mode.

// Any setting for prompt allowance is valid.

// Ensure that the response is empty.

// Ensure that prompt allowance hasn't been re-specified.

// If responding to a message, ensure that the response is empty. For
// prompt responses, any value is allowed.

// Success.

// EnsureValid verifies that a HostResponse is valid.
func (r *HostResponse) EnsureValid(first, allowPrompts bool) error {
	_ = "STUB: not implemented"
	// A nil hosting response is not valid.
	return nil
}

// Handle validation based on whether or not this is the first response.

// Ensure that the prompter identifier is specified.

// Ensure that no message type is specified.

// Ensure that no message is provided.

// Ensure that the prompter identifier isn't specified again.

// Ensure that the message type is allowed.

// Any value of the message is considered valid.

// Success.

// ensureValid verifies that a PromptRequest is valid.
func (r *PromptRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil prompt request is not valid.
	return nil
}

// Verify that the prompter identifier is non-empty.

// Any value of the prompt is considered valid.

// Success.

// EnsureValid verifies that a PromptResponse is valid.
func (r *PromptResponse) EnsureValid() error {
	_ = "STUB: not implemented"
	// A nil prompt response is not valid.
	return nil
}

// Any value of the response itself is considered valid.

// Success.
