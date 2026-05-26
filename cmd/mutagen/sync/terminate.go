package sync

import (
	"github.com/spf13/cobra"

	"google.golang.org/grpc"

	"github.com/mutagen-io/mutagen/pkg/selection"
)

// TerminateWithSelection is an orchestration convenience method that performs a
// terminate operation using the provided daemon connection and session
// selection.
func TerminateWithSelection(
	daemonConnection *grpc.ClientConn,
	selection *selection.Selection,
) error {
	_ = "STUB: not implemented"
	// Initiate command line messaging.
	return nil
}

// Perform the terminate operation, cancel prompting, and handle errors.

// Success.

// terminateMain is the entry point for the terminate command.
func terminateMain(_ *cobra.Command, arguments []string) error {
	_ = "STUB: not implemented"
	// Create session selection specification.
	return nil
}

// Connect to the daemon and defer closure of the connection.

// Perform the terminate operation.

// terminateCommand is the terminate command.
var terminateCommand = &cobra.Command{
	Use:          "terminate [<session>...]",
	Short:        "Permanently terminate a synchronization session",
	RunE:         terminateMain,
	SilenceUsage: true,
}

// terminateConfiguration stores configuration for the terminate command.
var terminateConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// all indicates whether or not all sessions should be terminated.
	all bool
	// labelSelector encodes a label selector to be used in identifying which
	// sessions should be paused.
	labelSelector string
}

func init() {
	// Grab a handle for the command line flags.
	flags := terminateCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&terminateConfiguration.help, "help", "h", false, "Show help information")

	// Wire up terminate flags.
	flags.BoolVarP(&terminateConfiguration.all, "all", "a", false, "Terminate all sessions")
	flags.StringVar(&terminateConfiguration.labelSelector, "label-selector", "", "Terminate sessions matching the specified label selector")
}
