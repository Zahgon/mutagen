//go:build go1.25

package mutagen

import (
	"fmt"
	"io"
)

const (
	// VersionMajor represents the current major version of Mutagen.
	VersionMajor = 0
	// VersionMinor represents the current minor version of Mutagen.
	VersionMinor = 19
	// VersionPatch represents the current patch version of Mutagen.
	VersionPatch = 0
	// VersionTag represents a tag to be appended to the Mutagen version string.
	// It must not contain spaces. If empty, no tag is appended to the version
	// string.
	VersionTag = "dev"
)

// DevelopmentModeEnabled indicates that development mode is active. This is
// regulated via VersionTag and should not be set or updated explicitly.
const DevelopmentModeEnabled = VersionTag == "dev"

// Version provides a stringified version of the current Mutagen version.
var Version string

// init performs global initialization.
func init() {
	// Compute the stringified version.
	if VersionTag != "" {
		Version = fmt.Sprintf("%d.%d.%d-%s", VersionMajor, VersionMinor, VersionPatch, VersionTag)
	} else {
		Version = fmt.Sprintf("%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
	}
}

// versionBytes is a type that can be used to send and receive version
// information over the wire.
type versionBytes [12]byte

// sendVersion writes the current version to the specified writer. Version tag
// components are neither transmitted nor received.
func sendVersion(writer io.Writer) error {
	_ = "STUB: not implemented"
	// Compute the version bytes.
	return nil
}

// Transmit the bytes.

// receiveVersion reads version information from the specified reader. Version
// tag components are neither transmitted nor received.
func receiveVersion(reader io.Reader) (uint32, uint32, uint32, error) {
	_ = "STUB: not implemented"
	// Read the bytes.
	return 0, 0, 0, nil
}

// Decode components.

// Done.

// ClientVersionHandshake performs the client side of a version handshake,
// returning an error if the received server version is not compatible with the
// client version.
//
// TODO: Add some ability to support version skew in this function.
func ClientVersionHandshake(stream io.ReadWriteCloser) error {
	_ = "STUB: not implemented"
	// Receive the server's version.
	return nil
}

// Send our version to the server.

// Ensure that our Mutagen versions are compatible. For now, we enforce that
// they're equal.
// TODO: Once we lock-in an internal protocol that we're going to support
// for some time, we can allow some version skew. On the client side in
// particular, we'll probably want to look out for the specific "locked-in"
// server protocol that we support and instantiate some frozen client
// implementation from that version.

// Success.

// ServerVersionHandshake performs the server side of a version handshake,
// returning an error if the received client version is not compatible with the
// server version.
//
// TODO: Add some ability to support version skew in this function.
func ServerVersionHandshake(stream io.ReadWriteCloser) error {
	_ = "STUB: not implemented"
	// Send our version to the client.
	return nil
}

// Receive the client's version.

// Ensure that our versions are compatible. For now, we enforce that they're
// equal.

// Success.
