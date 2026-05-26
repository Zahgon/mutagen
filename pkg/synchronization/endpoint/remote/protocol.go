package remote

const (
	// controlStreamBufferSize is the buffer size to use for control stream
	// buffering. It should be ideally large enough to fill the kernel buffer
	// for whatever stream is being used as a transport, which in our case is
	// typically an OS pipe.
	controlStreamCompressedBufferSize = 64 * 1024
	// controlStreamCompressorBufferSize is the buffer size to use for
	// compressor input and decompressor output.
	controlStreamUncompressedBufferSize = 64 * 1024
)

// ensureValid ensures that the InitializeSynchronizationRequest's invariants
// are respected.
func (r *InitializeSynchronizationRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil initialize request is not valid.
	return nil
}

// Ensure that the session identifier is non-empty.

// Ensure that the session version is supported.

// Ensure that the configuration is valid.

// Ensure that the root path is non-empty.

// There's no need to validate Alpha - either value is correct.

// Success.

// ensureValid ensures that the InitializeSynchronizationResponse's invariants
// are respected.
func (r *InitializeSynchronizationResponse) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil initialize response is not valid.
	return nil
}

// Success.

// ensureValid ensures that the PollRequest's invariants are respected.
func (r *PollRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil poll request is not valid.
	return nil
}

// Success.

// ensureValid ensures that the PollCompletionRequest's invariants are
// respected.
func (r *PollCompletionRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil poll completion request is not valid.
	return nil
}

// Success.

// ensureValid ensures that the PollResponse's invariants are respected.
func (r *PollResponse) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil poll response is not valid.
	return nil
}

// Success.

// ensureValid ensures that the ScanRequest's invariants are respected.
func (r *ScanRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil scan request is not valid.
	return nil
}

// Ensure that the baseline snapshot signature is valid.

// Full is correct regardless of value, so no validation is required.

// Success.

// ensureValid ensures that the ScanCompletionRequest's invariants are
// respected.
func (r *ScanCompletionRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil scan completion request is not valid.
	return nil
}

// Success.

// ensureValid ensures that the ScanResponse's invariants are respected.
func (r *ScanResponse) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil scan response is not valid.
	return nil
}

// Ensure that each snapshot delta operation is valid.

// If an error is set, then make sure that no snapshot delta was provided.

// Success.

// ensureValid ensures that the StageRequest's invariants are respected.
func (r *StageRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil stage request is not valid.
	return nil
}

// Ensure that there are a non-zero number of paths. This isn't an invariant
// that we really *need* to enforce, as our logic is capable of handling it,
// but it's a useful check to make sure that the client is avoiding
// transmission in these cases.

// NOTE: We could perform an additional check that the specified paths are
// unique, but this isn't quite so cheap, and it won't break anything if
// they're not.

// HACK: We don't verify that the paths are valid (and we'd have a hard time
// doing so in any sense other than syntactically) because we use the
// filesystem.Opener infrastructure to properly traverse the synchronization
// root. It would also be expensive to verify the correctness of these paths
// and it would be of little benefit. I'd class this as a hack because it's
// sort of a layering violation (the message shouldn't know about the code
// that uses it), but I'm willing to live with it because the message is so
// tightly coupled to the endpoint implementation anyway.

// Ensure that the number of digests matches the number of paths.

// NOTE: We could perform an additional check that the specified digests are
// valid, but this isn't really necessary, and we'd have to handle varying
// digest lengths.

// Success.

// ensureValid ensures that StageResponse's invariants are respected.
func (r *StageResponse) ensureValid(paths []string) error {
	_ = "STUB: not implemented"
	// A nil stage response is not valid.
	return nil
}

// Verify that path and signature counts are sane. We have to take into
// account the shorthand that's used when all paths are requested.

// Verify that all signatures are valid.

// Verify that paths and signatures are not present if there's an error.

// Success.

// ensureValid ensures that SupplyRequest's invariants are respected.
func (r *SupplyRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil supply request is not valid.
	return nil
}

// Ensure that the number of paths matches the number of signatures.

// Ensure that all signatures are valid.

// Success.

// ensureValid ensures that TransitionRequest's invariants are respected.
func (r *TransitionRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil transition request is not valid.
	return nil
}

// Ensure that each change is valid. Each should contain only synchronizable
// content.

// Success.

// ensureValid ensures that the TransitionCompletionRequest's invariants are
// respected.
func (r *TransitionCompletionRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil transition completion request is not valid.
	return nil
}

// Success.

// ensureValid ensures that TransitionResponse's invariants are respected.
func (r *TransitionResponse) ensureValid(expectedCount int) error {
	_ = "STUB: not implemented"
	// A nil transition response is not valid.
	return nil
}

// Ensure that the number of results matches the number expected.

// Validate that each result is a valid archive. Each should contain only
// synchronizable content.

// Validate that each problem is a valid problem specification.

// Success.

// ensureValid ensures that EndpointRequest's invariants are respected.
func (r *EndpointRequest) ensureValid() error {
	_ = "STUB: not implemented"
	// A nil endpoint request is not valid.
	return nil
}

// Ensure that exactly one field is set.

// Success.
