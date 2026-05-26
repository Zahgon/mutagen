package encoding

// LoadAndUnmarshal provides the underlying loading and unmarshaling
// functionality for the encoding package. It reads the data at the specified
// path and then invokes the specified unmarshaling callback (usually a closure)
// to decode the data.
func LoadAndUnmarshal(path string, unmarshal func([]byte) error) error {
	_ = "STUB: not implemented"
	// Grab the file contents.
	return nil
}

// Perform the unmarshaling.

// Success.

// MarshalAndSave provide the underlying marshaling and saving functionality for
// the encoding package. It invokes the specified marshaling callback (usually a
// closure) and writes the result atomically to the specified path. The data is
// saved with read/write permissions for the user only.
func MarshalAndSave(path string, marshal func() ([]byte, error)) error {
	_ = "STUB: not implemented"
	// Marshal the message.
	return nil
}

// Write the file atomically with secure file permissions.

// Success.
