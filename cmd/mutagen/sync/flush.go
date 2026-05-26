package sync

import (
	"github.com/spf13/cobra"

	"google.golang.org/grpc"

	"github.com/mutagen-io/mutagen/pkg/selection"
)

// FlushWithSelection is an orchestration convenience method that performs a
// flush operation using the provided daemon connection and session selection.
func FlushWithSelection(
	daemonConnection *grpc.ClientConn,
	selection *selection.Selection,
	skipWait bool,
) error {
	_ = "STUB: not implemented"
	// Initiate command line messaging.
	return nil
}

// Perform the flush operation, cancel prompting, and handle errors.

// Success.

// flushMain is the entry point for the flush command.
func flushMain(_ *cobra.Command, arguments []string) error {
	_ = "STUB: not implemented"
	// Create session selection specification.
	return nil
}

// Connect to the daemon and defer closure of the connection.

// Perform the flush operation.

// flushCommand is the flush command.
var flushCommand = &cobra.Command{
	Use:          "flush [<session>...]",
	Short:        "Force a synchronization cycle",
	RunE:         flushMain,
	SilenceUsage: true,
}

// flushConfiguration stores configuration for the flush command.
var flushConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// all indicates whether or not all sessions should be flushed.
	all bool
	// labelSelector encodes a label selector to be used in identifying which
	// sessions should be paused.
	labelSelector string
	// skipWait indicates whether or not the flush operation should block until
	// a synchronization cycle completes for each sesion requested.
	skipWait bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := flushCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&flushConfiguration.help, "help", "h", false, "Show help information")

	// Wire up flush flags.
	flags.BoolVarP(&flushConfiguration.all, "all", "a", false, "Flush all sessions")
	flags.StringVar(&flushConfiguration.labelSelector, "label-selector", "", "Flush sessions matching the specified label selector")
	flags.BoolVar(&flushConfiguration.skipWait, "skip-wait", false, "Avoid waiting for the resulting synchronization cycle(s) to complete")
}
