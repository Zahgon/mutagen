package project

import (
	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd"
)

// flushMain is the entry point for the flush command.
func flushMain(_ *cobra.Command, _ []string) error {
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

// Read the project identifier from the lock file. If the lock file is
// empty, then we can assume that we created it when we created the lock and
// just remove it.

// Ensure that the project identifier is valid.

// Connect to the daemon and defer closure of the connection.

// Compute the selection that we're going to use to flush sessions.

// Flush synchronization sessions.

// Success.

// flushCommand is the flush command.
var flushCommand = &cobra.Command{
	Use:          "flush",
	Short:        "Flush project synchronization sessions",
	Args:         cmd.DisallowArguments,
	RunE:         flushMain,
	SilenceUsage: true,
}

// flushConfiguration stores configuration for the flush command.
var flushConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// projectFile is the path to the project file, if non-default.
	projectFile string
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

	// Wire up project file flags.
	flags.StringVarP(&flushConfiguration.projectFile, "project-file", "f", "", "Specify project file")

	// Wire up flush flags.
	flags.BoolVar(&flushConfiguration.skipWait, "skip-wait", false, "Avoid waiting for the resulting synchronization cycle(s) to complete")
}
