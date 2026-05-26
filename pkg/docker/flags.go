package docker

// DaemonConnectionFlags encodes top-level Docker command line flags that
// control the Docker daemon connection. These flags are shared between the
// Docker CLI and Docker Compose. These flags can be loaded from Mutagen URL
// parameters or used as command line flag storage. The zero value of this
// structure is a valid value corresponding to the absence of any of these
// flags.
type DaemonConnectionFlags struct {
	// Config stores the value of the --config flag.
	Config string
	// Host stores the value of the -H/--host flag.
	Host string
	// Context stores the value of the -c/--context flag.
	Context string
	// TLS indicates the presence of the --tls flag.
	TLS bool
	// TLSCACert stores the value of the --tlscacert flag.
	TLSCACert string
	// TLSCert stores the value of the --tlscert flag.
	TLSCert string
	// TLSKey stores the value of the --tlskey flag.
	TLSKey string
	// TLSVerify indicates the presence of the --tlsverify flag.
	TLSVerify bool
}

// LoadDaemonConnectionFlagsFromURLParameters loads top-level Docker daemon
// connection flags from Mutagen URL parameters.
func LoadDaemonConnectionFlagsFromURLParameters(parameters map[string]string) (*DaemonConnectionFlags, error) {
	_ = "STUB: not implemented"
	// Create a zero-valued result (corresponding to no flags).
	return nil, nil
}

// Validate and convert parameters.

// Success.

// ToFlags reconstitues top-level daemon connection flags so that they can be
// passed to a Docker CLI or Docker Compose command.
func (f *DaemonConnectionFlags) ToFlags() []string {
	_ = "STUB: not implemented"
	// Set up the result.
	return nil
}

// Add flags as necessary.

// Done.

// ToURLParameters converts top-level daemon connection flags to parameters that
// can be embedded in a Mutagen URL. These parameters can be converted back
// using LoadDaemonConnectionFlagsFromURLParameters.
func (f *DaemonConnectionFlags) ToURLParameters() map[string]string {
	_ = "STUB: not implemented"
	// Create an empty set of parameters.
	return nil
}

// Add parameters as necessary.

// Done.
