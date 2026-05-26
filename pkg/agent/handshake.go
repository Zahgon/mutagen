package agent

import (
	"io"
)

// magicNumberBytes is a type capable of holding a Mutagen magic byte sequence.
type magicNumberBytes [3]byte

// serverMagicNumber is a byte sequence that is sent by an endpoint server to
// identify the start of a Mutagen protocol stream. It is intentionally composed
// of bytes that are not (all) printable ASCII characters. The purpose of adding
// this magic number to the beginning of streams is to work around agent-style
// transports where the underlying transport executable may write error output
// to standard output, which would otherwise be interpreted as version
// information. By identifying this magic number, we can be sure that we're
// talking to a Mutagen stream before we start exchanging version information.
var serverMagicNumber = magicNumberBytes{0x05, 0x27, 0x87}

// clientMagicNumber serves the same purpose as serverMagicNumber, but it is
// send by the endpoint client to the endpoint server. It is not as necessary,
// since whatever connects to the server should already know what it's doing,
// but it serves as an extra sanity check in the world of agent-style
// transports.
var clientMagicNumber = magicNumberBytes{0x87, 0x27, 0x05}

// sendMagicNumber sends the Mutagen magic byte sequence to the specified
// writer.
func sendMagicNumber(writer io.Writer, magicNumber magicNumberBytes) error {
	_ = "STUB: not implemented"
	return nil
}

// receiveAndCompareMagicNumber reads a Mutagen magic byte sequence from the
// specified reader and verifies that it matches what's expected.
func receiveAndCompareMagicNumber(reader io.Reader, expected magicNumberBytes) (bool, error) {
	_ = "STUB: not implemented"
	// Read the bytes.
	return false, nil
}

// Compare the bytes.

// ClientHandshake performs a client-side handshake on the stream.
func ClientHandshake(stream io.ReadWriter) error {
	_ = "STUB: not implemented"
	// Receive the server's magic number.
	return nil
}

// Send our magic number to the server.

// Success.

// ServerHandshake performs a server-side handshake on the stream.
func ServerHandshake(stream io.ReadWriter) error {
	_ = "STUB: not implemented"
	// Send our magic number to the client.
	return nil
}

// Receive the client's magic number. We treat a mismatch of the magic
// number as a transport error as well, because it indicates that we're not
// actually talking to a Mutagen client.

// Success.
