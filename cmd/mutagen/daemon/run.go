package daemon

import (
	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd"
)

// runMain is the entry point for the run command.
func runMain(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Attempt to acquire the daemon lock and defer its release.
	return nil
}

// Create a channel to track termination signals. We do this before creating
// and starting other infrastructure so that we can ensure things terminate
// smoothly, not mid-initialization.

// Create the root logger.

// Create a forwarding session manager and defer its shutdown.

// Create a synchronization session manager and defer its shutdown.

// Create the gRPC server and defer its termination. We use a hard stop
// rather than a graceful stop so that it doesn't hang on open requests.

// Create the daemon server, defer its shutdown, and register it.

// Create and register the prompt server.

// Create and register the forwarding server.

// Create and register the synchronization server.

// Compute the path to the daemon IPC endpoint.

// Create the daemon listener and defer its closure. Since we hold the
// daemon lock, we preemptively remove any existing socket since it should
// be stale.

// Serve incoming requests and watch for server failure.

// Wait for termination from a signal, the daemon service, or the gRPC
// server. We treat termination via the daemon service as a non-error.

// runCommand is the run command.
var runCommand = &cobra.Command{
	Use:          "run",
	Short:        "Run the Mutagen daemon",
	Args:         cmd.DisallowArguments,
	Hidden:       true,
	RunE:         runMain,
	SilenceUsage: true,
}

// runConfiguration stores configuration for the run command.
var runConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := runCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&runConfiguration.help, "help", "h", false, "Show help information")
}
