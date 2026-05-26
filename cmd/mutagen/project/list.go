package project

import (
	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd"
)

// listMain is the entry point for the list command.
func listMain(_ *cobra.Command, _ []string) error {
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

// Compute the selection that we're going to use to list sessions.

// List forwarding sessions.

// Print an empty line.

// List synchronization sessions.

// Success.

// listCommand is the list command.
var listCommand = &cobra.Command{
	Use:          "list",
	Short:        "List project sessions",
	Args:         cmd.DisallowArguments,
	RunE:         listMain,
	SilenceUsage: true,
}

// listConfiguration stores configuration for the list command.
var listConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// projectFile is the path to the project file, if non-default.
	projectFile string
	// long indicates whether or not to use long-format listing.
	long bool
}

func init() {
	// Grab a handle for the command line flags.
	flags := listCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&listConfiguration.help, "help", "h", false, "Show help information")

	// Wire up project file flags.
	flags.StringVarP(&listConfiguration.projectFile, "project-file", "f", "", "Specify project file")

	// Wire up list flags.
	flags.BoolVarP(&listConfiguration.long, "long", "l", false, "Show detailed session information")
}
