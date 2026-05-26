package url

import (
	"os"
)

const (
	// alphaSpecificEnvironmentVariablePrefix is the prefix to use when checking
	// for alpha-specific environment variables.
	alphaSpecificEnvironmentVariablePrefix = "MUTAGEN_ALPHA_"
	// betaSpecificEnvironmentVariablePrefix is the prefix to use when checking
	// for beta-specific environment variables.
	betaSpecificEnvironmentVariablePrefix = "MUTAGEN_BETA_"
	// sourceSpecificEnvironmentVariablePrefix is the prefix to use when
	// checking for source-specific environment variables.
	sourceSpecificEnvironmentVariablePrefix = "MUTAGEN_SOURCE_"
	// destinationSpecificEnvironmentVariablePrefix is the prefix to use when
	// checking for destination-specific environment variables.
	destinationSpecificEnvironmentVariablePrefix = "MUTAGEN_DESTINATION_"
)

// lookupEnv is the environment variable lookup function to use. It is a
// variable so that it can be swapped out during testing.
var lookupEnv = os.LookupEnv

// getEnvironmentVariable returns the value for the specified environment
// variable, as well as whether or not it was found. Endpoint-specific variables
// take precedence over non-specific variables.
func getEnvironmentVariable(name string, kind Kind, first bool) (string, bool) {
	_ = "STUB: not implemented"
	// Validate the variable name.
	return "", false
}

// Check for an endpoint-specific variant.

// Check for the general variant.
