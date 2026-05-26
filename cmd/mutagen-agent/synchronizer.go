package main

import (
	"context"
	"time"

	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd"

	"github.com/mutagen-io/mutagen/pkg/agent"
	"github.com/mutagen-io/mutagen/pkg/logging"
)

const (
	// housekeepingInterval is the interval at which housekeeping will be
	// invoked by the agent.
	housekeepingInterval = 24 * time.Hour
)

// housekeepRegularly is the entry point for the housekeeping Goroutine.
func housekeepRegularly(ctx context.Context, logger *logging.Logger) {
	_ = "STUB: not implemented"
	// Perform an initial housekeeping operation since the ticker won't fire
	// straight away.
	return
}

// Create a ticker to regulate housekeeping and defer its shutdown.

// Loop and wait for the ticker or cancellation.

// synchronizerMain is the entry point for the synchronizer command.
func synchronizerMain(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Create a channel to track termination signals. We do this before creating
	// and starting other infrastructure so that we can ensure things terminate
	// smoothly, not mid-initialization.
	return nil
}

// Set up a logger on the standard error stream.

// Set up regular housekeeping and defer its shutdown.

// Create a stream using standard input/output.

// Perform an agent handshake.

// Perform a version handshake.

// Serve a synchronizer on standard input/output and monitor for its
// termination.

// Wait for termination from a signal or the synchronizer.

// synchronizerCommand is the synchronizer command.
var synchronizerCommand = &cobra.Command{
	Use:          agent.CommandSynchronizer,
	Short:        "Run the agent in synchronizer mode",
	Args:         cmd.DisallowArguments,
	RunE:         synchronizerMain,
	SilenceUsage: true,
}

// synchronizerConfiguration stores configuration for the synchronizer command.
var synchronizerConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// logLevel indicates the log level to use.
	logLevel string
}

func init() {
	// Grab a handle for the command line flags.
	flags := synchronizerCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&synchronizerConfiguration.help, "help", "h", false, "Show help information")

	// Wire up logging flags.
	flags.StringVar(&synchronizerConfiguration.logLevel, agent.FlagLogLevel, "", "Set the log level")
}
