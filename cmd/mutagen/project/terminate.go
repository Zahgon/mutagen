package project

import (
	"github.com/spf13/cobra"

	"github.com/mutagen-io/mutagen/cmd"
)

// terminateMain is the entry point for the terminate command.
func terminateMain(_ *cobra.Command, _ []string) error {
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

// Load the configuration file.

// Perform pre-termination commands.

// Connect to the daemon and defer closure of the connection.

// Compute the selection that we're going to use to terminate sessions.

// Terminate forwarding sessions.

// Terminate synchronization sessions.

// Perform post-termination commands.

// Schedule the project lock for removal.

// Success.

// terminateCommand is the terminate command.
var terminateCommand = &cobra.Command{
	Use:          "terminate",
	Short:        "Terminate project sessions",
	Args:         cmd.DisallowArguments,
	RunE:         terminateMain,
	SilenceUsage: true,
}

// terminateConfiguration stores configuration for the terminate command.
var terminateConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// projectFile is the path to the project file, if non-default.
	projectFile string
}

func init() {
	// Grab a handle for the command line flags.
	flags := terminateCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&terminateConfiguration.help, "help", "h", false, "Show help information")

	// Wire up project file flags.
	flags.StringVarP(&terminateConfiguration.projectFile, "project-file", "f", "", "Specify project file")
}
