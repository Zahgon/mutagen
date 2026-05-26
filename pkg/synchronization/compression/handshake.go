package compression

import (
	"errors"
	"io"
)

// errAlgorithmUnsupported indicates that the algorithm requested by the client
// is unsupported by the server.
var errAlgorithmUnsupported = errors.New("algorithm unsupported")

// ClientHandshake performs a client-side compression handshake on the stream.
// It transmits the desired compression algorithm and verifies that this
// algorithm is supported by the server.
func ClientHandshake(stream io.ReadWriter, algorithm Algorithm) error {
	_ = "STUB: not implemented"
	// Verify that the algorithm can be encoded into a single byte.
	return nil
}

// Convert the algorithm specification.

// Transmit the data.

// Receive the response.

// Handle the response.

// ServerHandshake performs a server-side compression handshake on the stream.
// It receives the desired compression algorithm from the client, verifies that
// this algorithm is supported, and transmits a response to the client.
func ServerHandshake(stream io.ReadWriter) (Algorithm, error) {
	_ = "STUB: not implemented"
	// Receive the algorithm specification.
	return *new(Algorithm), nil
}

// Convert the algorithm specification and ensure that it's supported.

// Format and transmit the response.

// Handle unsupported algorithms.

// Success.
