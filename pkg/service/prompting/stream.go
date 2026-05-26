package prompting

// streamPrompter implements Prompter on top of a Prompting_HostServer stream.
type streamPrompter struct {
	// allowPrompts indicates whether or not the client allows prompts.
	allowPrompts bool
	// stream is the underlying Prompting_HostServer stream.
	stream Prompting_HostServer
	// errored indicates whether or not the stream has encountered an error.
	errored bool
}

// sendReceive performs a send/receive cycle by sending a HostResponse and
// receiving a HostRequest.
func (p *streamPrompter) sendReceive(response *HostResponse) (*HostRequest, error) {
	_ = "STUB: not implemented"
	// Send the request.
	return nil, nil
}

// Determine the expected request mode.

// Receive the response.

// Message implements the Message method of Prompter.
func (p *streamPrompter) Message(message string) error {
	_ = "STUB: not implemented"
	// Check if a previous transmission error has occurred.
	return nil
}

// Otherwise perform the messaging operation.

// Success.

// Prompt implements the Prompt method of Prompter.
func (p *streamPrompter) Prompt(prompt string) (string, error) {
	_ = "STUB: not implemented"
	// Check if a previous transmission error has occurred.
	return "", nil
}

// Check whether or not prompts are supported by this client.

// Perform the exchange.

// Success.
