package project

import (
	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd"
)

// startMain is the entry point for the start command.
func startMain(_ *cobra.Command, _ []string) error {
	_ = "STUB: not implemented"
	// Compute the name of the configuration file and ensure that our working
	// directory is that in which the file resides. This is required for
	// relative paths (including relative synchronization paths and relative
	// Unix Domain Socket paths) to be resolved relative to the project
	// configuration file.
	return nil
}

// Compute the lock path.

// Track whether or not we should remove the lock file on return.

// Create a locker and defer its closure and potential removal. On Windows
// systems, we have to handle this removal after the file is closed.

// Acquire the project lock and defer its release and potential removal. On
// Windows systems, we can't remove the lock file if it's locked or even
// just opened, so we handle removal for Windows systems after we close the
// lock file (see above). In this case, we truncate the lock file before
// releasing it to ensure that any other process that opens or acquires the
// lock file before we manage to remove it will simply see an empty lock
// file, which it will ignore or attempt to remove.

// Read the full contents of the lock file and ensure that it's empty.

// At this point we know that there was no previous project running, but we
// haven't yet created any resources, so defer removal of the lock file that
// we've created in case we run into any errors loading configuration
// information.

// Create a unique project identifier.

// Write the project identifier to the lock file.

// Load the configuration file.

// Unless disabled, attempt to load configuration from the global
// configuration file and use it as the base for our core session
// configurations.

// Compute the path to the global configuration file.

// Attempt to load and validate the file. We allow it to not exist.

// Extract and validate forwarding defaults.

// Extract and validate synchronization defaults.

// Merge global and default configurations, with defaults taking priority.

// Generate forward session creation specifications.

// Ignore defaults.

// Verify that the name is valid.

// Compute URLs.

// Parse URLs.

// Compute configuration.

// Compute source-specific configuration.

// Compute destination-specific configuration.

// Record the specification.

// Generate synchronization session creation specifications and keep track
// of those that we should flush on creation.

// Ignore defaults.

// Verify that the name is valid.

// Compute URLs.

// Parse URLs.

// Compute configuration.

// Compute alpha-specific configuration.

// Compute beta-specific configuration.

// Record the specification.

// Compute and store flush-on-creation behavior.

// Connect to the daemon and defer closure of the connection.

// At this point, we're going to try to create resources, so we need to
// maintain the lock file in case even some of them are successful.

// Perform pre-creation commands.

// Create forwarding sessions.

// Create synchronization sessions and track those that we should flush.

// Perform session creation.

// Determine whether or not to flush this session.

// Flush synchronization sessions for which flushing has been requested.

// Perform post-creation commands.

// Success.

// startCommand is the start command.
var startCommand = &cobra.Command{
	Use:          "start",
	Short:        "Start project sessions",
	Args:         cmd.DisallowArguments,
	RunE:         startMain,
	SilenceUsage: true,
}

// startConfiguration stores configuration for the start command.
var startConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// projectFile is the path to the project file, if non-default.
	projectFile string
	// paused indicates whether or not to create sessions in a pre-paused state.
	paused bool
	// noGlobalConfiguration specifies whether or not the global configuration
	// file should be ignored.
	noGlobalConfiguration bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := startCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&startConfiguration.help, "help", "h", false, "Show help information")

	// Wire up project file flags.
	flags.StringVarP(&startConfiguration.projectFile, "project-file", "f", "", "Specify project file")

	// Wire up paused flags.
	flags.BoolVarP(&startConfiguration.paused, "paused", "p", false, "Create the session pre-paused")

	// Wire up general configuration flags.
	flags.BoolVar(&startConfiguration.noGlobalConfiguration, "no-global-configuration", false, "Ignore the global configuration file")
}
