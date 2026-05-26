package forward

import (
	"github.com/mutagen-io/mutagen/cmd/mutagen/common"

	"github.com/mutagen-io/mutagen/pkg/forwarding"
	"github.com/mutagen-io/mutagen/pkg/url"
)

const (
	// emptyLabelValueDescription is a human-friendly description representing
	// an empty label value. It contains characters which are invalid for use in
	// label values, so it won't be confused for one.
	emptyLabelValueDescription = "<empty>"
)

// printEndpoint prints the configuration for a forwarding endpoint.
func printEndpoint(name string, url *url.URL, configuration *forwarding.Configuration, state *forwarding.EndpointState, version forwarding.Version, mode common.SessionDisplayMode) {
	_ = "STUB: not implemented"
	// Print the endpoint header.
	return
}

// Print the URL.

// Print configuration information if desired.

// Print configuration header.

// Compute and print the socket overwrite mode.

// Compute and print the socket owner.

// Compute and print the socket group.

// Compute and print the socket permission mode.

// At this point, there's no other status information that will be displayed
// for non-list modes, so we can save ourselves some checks and return if
// we're in a monitor mode.

// Print connection status.

// printSession prints the configuration and status of a forwarding session and
// its endpoints.
func printSession(state *forwarding.State, mode common.SessionDisplayMode) {
	_ = "STUB: not implemented"
	// Print name, if any.
	return
}

// Print the session identifier.

// Print extended information, if desired.

// Print labels, if any.

// Print the configuration header and configuration.
// TODO: Implement this in the future if we have any session-level
// configuration behaviors. We currently only have endpoint-level
// configuration behaviors and thus we don't print any session-level
// configuration information like we do for synchronization.

// Compute and print source-specific configuration.

// Compute and print beta-specific configuration.

// At this point, there's no other status information that will be displayed
// for non-list modes, so we can save ourselves some checks and return if
// we're in a monitor mode.

// Print the last error, if any.

// Print the session status .

// Print connection statistics if we're forwarding.
