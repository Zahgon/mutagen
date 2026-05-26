package main

import (
	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd"

	"github.com/mutagen-io/mutagen/pkg/agent"
)

// forwarderMain is the entry point for the forwarder command.
func forwarderMain(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Create a channel to track termination signals. We do this before creating
	// and starting other infrastructure so that we can ensure things terminate
	// smoothly, not mid-initialization.
	return nil
}

// Set up a logger on the standard error stream.

// Create a stream using standard input/output.

// Perform an agent handshake.

// Perform a version handshake.

// Serve a forwarder on standard input/output and monitor for its
// termination.

// Wait for termination from a signal or the forwarder.

// forwarderCommand is the forwarder command.
var forwarderCommand = &cobra.Command{
	Use:          agent.CommandForwarder,
	Short:        "Run the agent in forwarder mode",
	Args:         cmd.DisallowArguments,
	RunE:         forwarderMain,
	SilenceUsage: true,
}

// forwarderConfiguration stores configuration for the forwarder command.
var forwarderConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// logLevel indicates the log level to use.
	logLevel string
}

func init() {
	// Grab a handle for the command line flags.
	flags := forwarderCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&forwarderConfiguration.help, "help", "h", false, "Show help information")

	// Wire up logging flags.
	flags.StringVar(&forwarderConfiguration.logLevel, agent.FlagLogLevel, "", "Set the log level")
}
