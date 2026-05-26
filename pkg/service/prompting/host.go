package prompting

import (
	"context"

	"github.com/mutagen-io/mutagen/pkg/prompting"
)

// Host is a utility function for hosting a prompter via the Prompting service's
// Host method. Although the Host method can be used directly, it requires
// complex interaction and most callers will simply want to host a prompter.
// Prompting is hosted in a background Goroutine. The identifier for the
// prompter is returned, as well as an error channel that will be populated with
// the first error to occur during prompting. Hosting will be terminated when
// either an error occurs or the provided context is cancelled. The error
// channel will be closed after hosting has terminated. If an error occurs
// during hosting setup, then it will be returned and hosting will not commence.
func Host(
	ctx context.Context, client PromptingClient,
	prompter prompting.Prompter, allowPrompts bool,
) (string, <-chan error, error) {
	_ = "STUB: not implemented"
	// Create a subcontext that we can use to perform cancellation in case of a
	// client-side messaging or prompting error.
	return "", nil, nil
}

// Initiate hosting.

// Send the initialization request.

// Receive the initialization response, validate it, and extract the
// prompter identifier.

// Create an error monitoring channel.

// Start hosting in a background Goroutine.

// Defer closure of the errors channel.

// Defer cancellation of the context to ensure context resource cleanup
// and server-side cancellation in the event of a client-side error.

// Loop and handle requests indefinitely.

// Success.
