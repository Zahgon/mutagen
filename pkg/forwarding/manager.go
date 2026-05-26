package forwarding

import (
	"context"

	"github.com/mutagen-io/mutagen/pkg/logging"
	"github.com/mutagen-io/mutagen/pkg/selection"
	"github.com/mutagen-io/mutagen/pkg/state"
	"github.com/mutagen-io/mutagen/pkg/url"
)

// Manager provides forwarding session management facilities. Its methods are
// safe for concurrent usage, so it can be easily exported via an RPC interface.
type Manager struct {
	// logger is the underlying logger.
	logger *logging.Logger
	// tracker tracks changes to session states.
	tracker *state.Tracker
	// sessionLock locks the sessions registry.
	sessionsLock *state.TrackingLock
	// sessions maps sessions to their respective controllers.
	sessions map[string]*controller
}

// NewManager creates a new Manager instance.
func NewManager(logger *logging.Logger) (*Manager, error) {
	_ = "STUB: not implemented"
	// Create a tracker and corresponding lock to watch for state changes.
	return nil, nil
}

// Create the session registry.

// Load existing sessions.

// Success.

// allControllers creates a list of all controllers managed by the manager.
func (m *Manager) allControllers() []*controller {
	_ = "STUB: not implemented"
	// Grab the registry lock and defer its release.
	return nil
}

// Generate a list of all controllers.

// Done.

// findControllersBySpecification generates a list of controllers matching the
// given specifications.
func (m *Manager) findControllersBySpecification(specifications []string) ([]*controller, error) {
	_ = "STUB: not implemented"
	// Grab the registry lock and defer its release.
	return nil, nil
}

// Generate a list of controllers matching the specifications. We allow each
// specification to match multiple controllers, so we store matches in a set
// before converting them to a list. We do require that each specification
// match at least one controller.

// Convert the set to a list.

// Done.

// findControllersByLabelSelector generates a list of controllers using the
// specified label selector.
func (m *Manager) findControllersByLabelSelector(labelSelector string) ([]*controller, error) {
	_ = "STUB: not implemented"
	// Parse the label selector.
	return nil, nil
}

// Grab the registry lock and defer its release.

// Loop over controllers and look for matches.

// Done.

// selectControllers generates a list of controllers using the mechanism
// specified by the provided selection.
func (m *Manager) selectControllers(selection *selection.Selection) ([]*controller, error) {
	_ = "STUB: not implemented"
	// Dispatch selection based on the requested mechanism.
	return nil, nil
}

// TODO: Should we panic here instead?

// Shutdown tells the manager to gracefully halt sessions.
func (m *Manager) Shutdown() {
	_ = "STUB: not implemented"
	// Log the shutdown.
	return
}

// Terminate state tracking to terminate monitoring.

// Grab the registry lock and defer its release.

// Attempt to halt each session so that it can shutdown cleanly. Ignore but
// log any that fail to halt.

// Create tells the manager to create a new session.
func (m *Manager) Create(
	ctx context.Context,
	source, destination *url.URL,
	configuration, configurationSource, configurationDestination *Configuration,
	name string,
	labels map[string]string,
	paused bool,
	prompter string,
) (string, error) {
	_ = "STUB: not implemented"
	// Create a unique session identifier.
	return "", nil
}

// Attempt to create a session.

// Register the controller.

// Done.

// List requests a state snapshot for the specified sessions.
func (m *Manager) List(ctx context.Context, selection *selection.Selection, previousStateIndex uint64) (uint64, []*State, error) {
	_ = "STUB: not implemented"
	// Wait for a state change from the previous index.
	return 0, nil, nil
}

// Extract the controllers for the sessions of interest.

// Create a static snapshot of the state from each controller.

// Sort session states by session creation time.

// Success.

// Pause tells the manager to pause sessions matching the given specifications.
func (m *Manager) Pause(ctx context.Context, selection *selection.Selection, prompter string) error {
	_ = "STUB: not implemented"
	// Extract the controllers for the sessions of interest.
	return nil
}

// Attempt to pause the sessions.

// Success.

// Resume tells the manager to resume sessions matching the given
// specifications.
func (m *Manager) Resume(ctx context.Context, selection *selection.Selection, prompter string) error {
	_ = "STUB: not implemented"
	// Extract the controllers for the sessions of interest.
	return nil
}

// Attempt to resume.

// Success.

// Terminate tells the manager to terminate sessions matching the given
// specifications.
func (m *Manager) Terminate(ctx context.Context, selection *selection.Selection, prompter string) error {
	_ = "STUB: not implemented"
	// Extract the controllers for the sessions of interest.
	return nil
}

// Attempt to terminate the sessions. Since we're terminating them, we're
// responsible for removing them from the session map.

// Success.
