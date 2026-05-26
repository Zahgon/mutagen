package sync

import (
	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd/mutagen/common/templating"

	"github.com/mutagen-io/mutagen/pkg/synchronization"
)

// computeMonitorStatusLine constructs a monitoring status line for a
// synchronization session.
func computeMonitorStatusLine(state *synchronization.State) string {
	_ = "STUB: not implemented"
	// Build the status line.
	return ""
}

// Add a conflict flag if there are conflicts.

// Add a problems flag if there are problems.

// Add an error flag if there is one present.

// Handle the formatting based on status. If we're in a staging mode,
// then extract the relevant progress information. Despite not having a
// built-in mechanism for knowing the total expected size of a staging
// operation, we do know the number of files that the staging operation
// is performing, so if that's equal to the number of files on the
// source endpoint, then we know that we can use the total file size on
// the source endpoint as an estimate for the total staging size.

// Print staging progress, if available.

// Done.

// monitorMain is the entry point for the monitor command.
func monitorMain(_ *cobra.Command, arguments []string) error {
	_ = "STUB: not implemented"
	// Create the session selection specification that will select our initial
	// batch of sessions.
	return nil
}

// Load the formatting template (if any has been specified).

// Determine the listing mode.

// Connect to the daemon and defer closure of the connection.

// Create a session service client.

// Create the list request that we'll use.

// If no template has been specified, then create a status line printer with
// bold text and defer a line break operation.

// Track the last update time.

// Track whether or not we've identified an individual session in the
// non-templated case.

// Loop and print monitoring information indefinitely.

// Regulate the update frequency (and tame CPU usage in both the monitor
// command and the daemon) by enforcing a minimum update cycle interval.

// Perform a list operation.

// Update the state tracking index.

// If a template has been specified, then use that to format output with
// public model types. No validation is necessary here since we don't
// require any specific number of sessions.

// No template has been specified, but our command line monitoring
// interface only supports dynamic status displays for a single session
// at a time, so we choose the newest session identified by the initial
// criteria and update our selection to target it specifically.

// Select the most recently created session matching the
// selection criteria (which are ordered by creation date).

// Update the selection criteria to target only that session.

// Print session information.

// Record that we've identified our target session.

// Compute the status line.

// Print the status line.

// monitorCommand is the monitor command.
var monitorCommand = &cobra.Command{
	Use:          "monitor [<session>...]",
	Short:        "Display streaming session status information",
	RunE:         monitorMain,
	SilenceUsage: true,
}

// monitorConfiguration stores configuration for the monitor command.
var monitorConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// long indicates whether or not to use long-format monitoring.
	long bool
	// labelSelector encodes a label selector to be used in identifying which
	// sessions should be paused.
	labelSelector string
	// TemplateFlags store custom templating behavior.
	templating.TemplateFlags
}

func init() {
	// Grab a handle for the command line flags.
	flags := monitorCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&monitorConfiguration.help, "help", "h", false, "Show help information")

	// Wire up monitor flags.
	flags.BoolVarP(&monitorConfiguration.long, "long", "l", false, "Show detailed session information")
	flags.StringVar(&monitorConfiguration.labelSelector, "label-selector", "", "Monitor the most recently created session matching the specified label selector")

	// Wire up templating flags.
	monitorConfiguration.TemplateFlags.Register(flags)
}
