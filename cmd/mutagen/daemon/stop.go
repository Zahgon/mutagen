package daemon

import (
	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd"
)

// stopMain is the entry point for the stop command.
func stopMain(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// If the daemon is registered with the system, it may have a different stop
	// mechanism, so see if the system should handle it.
	return nil
}

// Connect to the daemon and defer closure of the connection. We avoid
// version compatibility checks since they would remove the ability to
// terminate an incompatible daemon. This is fine since the daemon service
// portion of the daemon API is stable.

// Create a daemon service client.

// Invoke shutdown. We don't check the response or error, because the daemon
// may terminate before it has a chance to send the response.

// Success.

// stopCommand is the stop command.
var stopCommand = &cobra.Command{
	Use:          "stop",
	Short:        "Stop the Mutagen daemon if it's running",
	Args:         cmd.DisallowArguments,
	RunE:         stopMain,
	SilenceUsage: true,
}

// stopConfiguration stores configuration for the stop command.
var stopConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := stopCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&stopConfiguration.help, "help", "h", false, "Show help information")
}
