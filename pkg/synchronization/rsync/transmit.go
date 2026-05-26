package rsync

// Transmit performs streaming transmission of files (in rsync deltified form)
// to the specified receiver. It is the responsibility of the caller to ensure
// that the provided signatures are valid by invoking their EnsureValid method.
// In order for this function to perform efficiently, paths should be passed in
// depth-first traversal order.
func Transmit(root string, paths []string, signatures []*Signature, receiver Receiver) error {
	_ = "STUB: not implemented"
	// Ensure that the transmission request is sane.
	return nil
}

// Create a file opener that we can use to safely open files, and defer its
// closure.

// Create an rsync engine.

// Create a transmission object that we can re-use to avoid allocating.

// Handle the requested files.

// Open the file and extract its size. Failure here is non-terminal, but
// we need to inform the receiver. If sending the message fails, that is
// a terminal error.

// Create an operation transmitter for deltification and track reception
// errors. We can safely set transmitError on each call because as soon
// as it's returned non-nil, the transmit function won't be called
// again.

// Perform deltification.

// Close the file.

// Handle any transmission errors. These are terminal.

// Inform the client the operation stream for this file is complete. Any
// internal (non-transmission) errors are non-terminal but should be
// reported to the receiver.

// Ensure that the receiver is finalized.

// Success.
