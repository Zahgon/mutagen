package forward

import (
	"github.com/spf13/cobra"

	"google.golang.org/grpc"

	"github.com/mutagen-io/mutagen/pkg/forwarding"
	forwardingsvc "github.com/mutagen-io/mutagen/pkg/service/forwarding"
)

// loadAndValidateGlobalSynchronizationConfiguration loads a YAML-based global
// configuration, extracts the forwarding component, converts it to a Protocol
// Buffers session configuration, and validates it.
func loadAndValidateGlobalForwardingConfiguration(path string) (*forwarding.Configuration, error) {
	_ = "STUB: not implemented"
	// Load the YAML configuration.
	return nil, nil
}

// Convert the YAML configuration to a Protocol Buffers representation and
// validate it.

// Success.

// CreateWithSpecification is an orchestration convenience method that performs
// a create operation using the provided daemon connection and session
// specification.
func CreateWithSpecification(
	daemonConnection *grpc.ClientConn,
	specification *forwardingsvc.CreationSpecification,
) (string, error) {
	_ = "STUB: not implemented"
	// Initiate command line prompting.
	return "", nil
}

// Perform the create operation, cancel prompting, and handle errors.

// Success.

// createMain is the entry point for the create command.
func createMain(_ *cobra.Command, arguments []string) error {
	_ = "STUB: not implemented"
	// Validate, extract, and parse URLs.
	return nil
}

// Validate the name.

// Parse, validate, and record labels.

// Create a default session configuration that will form the basis of our
// cumulative configuration.

// Unless disabled, attempt to load configuration from the global
// configuration file and merge it into our cumulative configuration.

// Compute the path to the global configuration file.

// Attempt to load the file. We allow it to not exist.

// If additional default configuration files have been specified, then load
// them and merge them into the cumulative configuration.

// Validate and convert socket overwrite mode specifications.

// Validate socket owner specifications.

// Validate socket group specifications.

// Validate and convert socket permission mode specifications.

// Create the command line configuration and merge it into our cumulative
// configuration.

// Create the creation specification.

// Connect to the daemon and defer closure of the connection.

// Perform the create operation.

// Print the session identifier.

// Success.

// createCommand is the create command.
var createCommand = &cobra.Command{
	Use:          "create <source> <destination>",
	Short:        "Create and start a new forwarding session",
	RunE:         createMain,
	SilenceUsage: true,
}

// createConfiguration stores configuration for the create command.
var createConfiguration struct {
	// help indicates whether or not to show help information and exit.
	help bool
	// name is the name specification for the session.
	name string
	// labels are the label specifications for the session.
	labels []string
	// paused indicates whether or not to create the session in a pre-paused
	// state.
	paused bool
	// noGlobalConfiguration specifies whether or not the global configuration
	// file should be ignored.
	noGlobalConfiguration bool
	// configurationFiles stores paths of additional files from which to load
	// default configuration.
	configurationFiles []string
	// socketOverwriteMode specifies the socket overwrite mode to use for the
	// session.
	socketOverwriteMode string
	// socketOverwriteModeSource specifies the socket overwrite mode to use for
	// the session, taking priority over socketOverwriteMode on source if
	// specified.
	socketOverwriteModeSource string
	// socketOverwriteModeDestination specifies the socket overwrite mode to use
	// for the session, taking priority over socketOverwriteMode on destination
	// if specified.
	socketOverwriteModeDestination string
	// socketOwner specifies the socket owner identifier to use new Unix domain
	// socket listeners, with endpoint-specific specifications taking priority.
	socketOwner string
	// socketOwnerSource specifies the socket owner identifier to use new Unix
	// domain socket listeners, taking priority over socketOwner on source if
	// specified.
	socketOwnerSource string
	// socketOwnerDestination specifies the socket owner identifier to use new
	// Unix domain socket listeners, taking priority over socketOwner on
	// destination if specified.
	socketOwnerDestination string
	// socketGroup specifies the socket owner identifier to use new Unix domain
	// socket listeners, with endpoint-specific specifications taking priority.
	socketGroup string
	// socketGroupSource specifies the socket owner identifier to use new Unix
	// domain socket listeners, taking priority over socketGroup on source if
	// specified.
	socketGroupSource string
	// socketGroupDestination specifies the socket owner identifier to use new
	// Unix domain socket listeners, taking priority over socketGroup on
	// destination if specified.
	socketGroupDestination string
	// socketPermissionMode specifies the socket permission mode to use for new
	// Unix domain socket listeners, with endpoint-specific specifications
	// taking priority.
	socketPermissionMode string
	// socketPermissionModeSource specifies the socket permission mode to use
	// for new Unix domain socket listeners on source, taking priority over
	// socketPermissionMode on source if specified.
	socketPermissionModeSource string
	// socketPermissionModeDestination specifies the socket permission mode to
	// use for new Unix domain socket listeners on destination, taking priority
	// over socketPermissionMode on destination if specified.
	socketPermissionModeDestination string
}

func init() {
	// Grab a handle for the command line flags.
	flags := createCommand.Flags()

	// Disable alphabetical sorting of flags in help output.
	flags.SortFlags = false

	// Manually add a help flag to override the default message. Cobra will
	// still implement its logic automatically.
	flags.BoolVarP(&createConfiguration.help, "help", "h", false, "Show help information")

	// Wire up name and label flags.
	flags.StringVarP(&createConfiguration.name, "name", "n", "", "Specify a name for the session")
	flags.StringSliceVarP(&createConfiguration.labels, "label", "l", nil, "Specify labels")

	// Wire up paused flags.
	flags.BoolVarP(&createConfiguration.paused, "paused", "p", false, "Create the session pre-paused")

	// Wire up general configuration flags.
	flags.BoolVar(&createConfiguration.noGlobalConfiguration, "no-global-configuration", false, "Ignore the global configuration file")
	flags.StringSliceVarP(&createConfiguration.configurationFiles, "configuration-file", "c", nil, "Specify additional files from which to load (and merge) default configuration parameters")

	// Wire up socket flags.
	flags.StringVar(&createConfiguration.socketOverwriteMode, "socket-overwrite-mode", "", "Specify socket overwrite mode (leave|overwrite)")
	flags.StringVar(&createConfiguration.socketOverwriteModeSource, "socket-overwrite-mode-source", "", "Specify socket overwrite mode for source (leave|overwrite)")
	flags.StringVar(&createConfiguration.socketOverwriteModeDestination, "socket-overwrite-mode-destination", "", "Specify socket overwrite mode for destination (leave|overwrite)")
	flags.StringVar(&createConfiguration.socketOwner, "socket-owner", "", "Specify socket owner")
	flags.StringVar(&createConfiguration.socketOwnerSource, "socket-owner-source", "", "Specify socket owner for source")
	flags.StringVar(&createConfiguration.socketOwnerDestination, "socket-owner-destination", "", "Specify socket owner for destination")
	flags.StringVar(&createConfiguration.socketGroup, "socket-group", "", "Specify socket group")
	flags.StringVar(&createConfiguration.socketGroupSource, "socket-group-source", "", "Specify socket group for source")
	flags.StringVar(&createConfiguration.socketGroupDestination, "socket-group-destination", "", "Specify socket group for destination")
	flags.StringVar(&createConfiguration.socketPermissionMode, "socket-permission-mode", "", "Specify socket permission mode")
	flags.StringVar(&createConfiguration.socketPermissionModeSource, "socket-permission-mode-source", "", "Specify socket permission mode for source")
	flags.StringVar(&createConfiguration.socketPermissionModeDestination, "socket-permission-mode-destination", "", "Specify socket permission mode for destination")
}
