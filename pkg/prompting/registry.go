package prompting

import (
	"sync"
)

// registryLock is the lock on the global prompter registry.
var registryLock sync.RWMutex

// registry is the global prompter registry.
var registry = make(map[string]chan Prompter)

// RegisterPrompter registers a prompter with the global registry. It
// automatically generates a unique identifier for the prompter.
func RegisterPrompter(prompter Prompter) (string, error) {
	_ = "STUB: not implemented"
	// Generate a unique identifier for this prompter.
	return "", nil
}

// Perform registration.

// Success.

// RegisterPrompterWithIdentifier registers a prompter with the global registry
// using the specified identifier.
func RegisterPrompterWithIdentifier(identifier string, prompter Prompter) error {
	_ = "STUB: not implemented"
	// Enforce that the identifier is non-empty.
	return nil
}

// Create and populate a "holder" (channel) for passing the prompter around.

// Lock the registry for writing and defer its release.

// Check for identifier collisions. This won't be a problem with our
// internally generated identifiers, but since this method accepts arbitrary
// identifiers, we want to be sure to avoid collisions.

// Register the holder.

// Success.

// UnregisterPrompter unregisters a prompter from the global registry. If the
// prompter is not registered, this method panics. If a prompter is unregistered
// with prompts pending for it, they will be cancelled.
func UnregisterPrompter(identifier string) {
	_ = "STUB: not implemented"
	// Lock the registry for writing, grab the holder, and remove it from the
	// registry. If it isn't currently registered, this must be a logic error.
	return
}

// Get the prompter back and close the holder to let anyone else who has it
// know that they won't be getting the prompter from it.

// Message invokes the Message method on a prompter in the global registry. If
// the prompter identifier provided is an empty string, this method is a no-op
// and returns a nil error.
func Message(identifier, message string) error {
	_ = "STUB: not implemented"
	// If the prompter identifier is empty, don't do anything.
	return nil
}

// Grab the holder for the specified prompter. We only need a read lock on
// the registry for this purpose.

// Acquire the prompter.

// Perform messaging.

// Return the prompter to the holder.

// Handle errors.

// Success.

// Prompt invokes the Prompt method on a prompter in the global registry.
func Prompt(identifier, prompt string) (string, error) {
	_ = "STUB: not implemented"
	// Grab the holder for the specified prompter. We only need a read lock on
	// the registry for this purpose.
	return "", nil
}

// Acquire the prompter.

// Perform prompting.

// Return the prompter to the holder.

// Handle errors.

// Success.
