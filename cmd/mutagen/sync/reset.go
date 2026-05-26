package sync

import (
	"github.com/spf13/cobra"

	"google.golang.org/grpc"

	"github.com/mutagen-io/mutagen/pkg/selection"
)

// ResetWithSelection is an orchestration convenience method that performs a
// reset operation using the provided daemon connection and session selection.
func ResetWithSelection(
	daemonConnection *grpc.ClientConn,
	selection *selection.Selection,
) error {
	_ = "STUB: not implemented"
	// Initiate command line prompting.
	return nil
}

// Perform the reset operation, cancel prompting, and handle errors.

// Success.

// resetMain is the entry point for the reset command.
func resetMain(_ *cobra.Command, arguments []string) error {
	_ = "STUB: not implemented"
	// Create session selection specification.
	return nil
}

// Connect to the daemon and defer closure of the connection.

// Perform the reset operation.

// resetCommand is the reset command.
var resetCommand = &cobra.Command{
	Use:          "reset [<session>...]",
	Short:        "Reset synchronization session history",
	RunE:         resetMain,
	SilenceUsage: true,
}

// resetConfiguration stores configuration for the reset command.
var resetConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// all indicates whether or not all sessions should be reset.
	all bool
	// labelSelector encodes a label selector to be used in identifying which
	// sessions should be paused.
	labelSelector string
}

func init() {
	// Grab a handle for the command line flags.
	flags := resetCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&resetConfiguration.help, "help", "h", false, "Show help information")

	// Wire up reset flags.
	flags.BoolVarP(&resetConfiguration.all, "all", "a", false, "Reset all sessions")
	flags.StringVar(&resetConfiguration.labelSelector, "label-selector", "", "Reset sessions matching the specified label selector")
}
