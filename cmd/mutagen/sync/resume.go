package sync

import (
	"github.com/spf13/cobra"

	"google.golang.org/grpc"

	"github.com/mutagen-io/mutagen/pkg/selection"
)

// ResumeWithSelection is an orchestration convenience method that performs a
// resume operation using the provided daemon connection and session selection.
func ResumeWithSelection(
	daemonConnection *grpc.ClientConn,
	selection *selection.Selection,
) error {
	_ = "STUB: not implemented"
	// Initiate command line prompting.
	return nil
}

// Perform the resume operation, cancel prompting, and handle errors.

// Success.

// resumeMain is the entry point for the resume command.
func resumeMain(_ *cobra.Command, arguments []string) error {
	_ = "STUB: not implemented"
	// Create session selection specification.
	return nil
}

// Connect to the daemon and defer closure of the connection.

// Perform the resume operation.

// resumeCommand is the resume command.
var resumeCommand = &cobra.Command{
	Use:          "resume [<session>...]",
	Short:        "Resume a paused or disconnected synchronization session",
	RunE:         resumeMain,
	SilenceUsage: true,
}

// resumeConfiguration stores configuration for the resume command.
var resumeConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// all indicates whether or not all sessions should be resumed.
	all bool
	// labelSelector encodes a label selector to be used in identifying which
	// sessions should be paused.
	labelSelector string
}

func init() {
	// Grab a handle for the command line flags.
	flags := resumeCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&resumeConfiguration.help, "help", "h", false, "Show help information")

	// Wire up resume flags.
	flags.BoolVarP(&resumeConfiguration.all, "all", "a", false, "Resume all sessions")
	flags.StringVar(&resumeConfiguration.labelSelector, "label-selector", "", "Resume sessions matching the specified label selector")
}
