package forward

import (
	"github.com/spf13/cobra"

	"google.golang.org/grpc"

	"github.com/mutagen-io/mutagen/cmd/mutagen/common/templating"

	"github.com/mutagen-io/mutagen/pkg/selection"
)

// ListWithSelection is an orchestration convenience method that performs a list
// operation using the provided daemon connection and session selection and then
// prints status information.
func ListWithSelection(
	daemonConnection *grpc.ClientConn,
	selection *selection.Selection,
	long bool,
) error {
	_ = "STUB: not implemented"
	// Load the formatting template (if any has been specified).
	return nil
}

// Determine the listing mode.

// Perform the list operation.

// If a template was specified, then use that to format output with public
// model types, otherwise use custom formatting code.

// Success.

// listMain is the entry point for the list command.
func listMain(_ *cobra.Command, arguments []string) error {
	_ = "STUB: not implemented"
	// Create session selection specification.
	return nil
}

// Connect to the daemon and defer closure of the connection.

// Perform the list operation and print status information.

// listCommand is the list command.
var listCommand = &cobra.Command{
	Use:          "list [<session>...]",
	Short:        "List existing forwarding sessions and their statuses",
	RunE:         listMain,
	SilenceUsage: true,
}

// listConfiguration stores configuration for the list command.
var listConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// long indicates whether or not to use long-format listing.
	long bool
	// labelSelector encodes a label selector to be used in identifying which
	// sessions should be paused.
	labelSelector string
	// TemplateFlags store custom templating behavior.
	templating.TemplateFlags
}

func init() {
	// Grab a handle for the command line flags.
	flags := listCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&listConfiguration.help, "help", "h", false, "Show help information")

	// Wire up list flags.
	flags.BoolVarP(&listConfiguration.long, "long", "l", false, "Show detailed session information")
	flags.StringVar(&listConfiguration.labelSelector, "label-selector", "", "List sessions matching the specified label selector")

	// Wire up templating flags.
	listConfiguration.TemplateFlags.Register(flags)
}
