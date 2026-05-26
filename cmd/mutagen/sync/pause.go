package sync

import (
	"github.com/spf13/cobra"

	"google.golang.org/grpc"

	"github.com/mutagen-io/mutagen/pkg/selection"
)

// PauseWithSelection is an orchestration convenience method that performs a
// pause operation using the provided service client and session selection.
func PauseWithSelection(
	daemonConnection *grpc.ClientConn,
	selection *selection.Selection,
) error {
	_ = "STUB: not implemented"
	// Initiate command line messaging.
	return nil
}

// Perform the pause operation, cancel prompting, and handle errors.

// Success.

// pauseMain is the entry point for the pause command.
func pauseMain(_ *cobra.Command, arguments []string) error {
	_ = "STUB: not implemented"
	// Create session selection specification.
	return nil
}

// Connect to the daemon and defer closure of the connection.

// Perform the pause operation.

// pauseCommand is the pause command.
var pauseCommand = &cobra.Command{
	Use:          "pause [<session>...]",
	Short:        "Pause a synchronization session",
	RunE:         pauseMain,
	SilenceUsage: true,
}

// pauseConfiguration stores configuration for the pause command.
var pauseConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// all indicates whether or not all sessions should be paused.
	all bool
	// labelSelector encodes a label selector to be used in identifying which
	// sessions should be paused.
	labelSelector string
}

func init() {
	// Grab a handle for the command line flags.
	flags := pauseCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&pauseConfiguration.help, "help", "h", false, "Show help information")

	// Wire up pause flags.
	flags.BoolVarP(&pauseConfiguration.all, "all", "a", false, "Pause all sessions")
	flags.StringVar(&pauseConfiguration.labelSelector, "label-selector", "", "Pause sessions matching the specified label selector")
}
