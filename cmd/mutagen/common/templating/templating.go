package templating

import (
	"text/template"

	"github.com/mutagen-io/mutagen/pkg/platform/terminal"
)

// jsonify is the built-in JSON encoder that's made available to templates.
func jsonify(value any) (string, error) {
	_ = "STUB: not implemented"
	// Create a buffer to store the output.
	return "", nil
}

// Create and configure a JSON encoder.

// Marshal the value.

// Convert the encoded JSON to a string.

// Remove the trailing newline that's automatically added by Encode.

// Success.

// builtins are the builtin functions supported in output templates.
var builtins = template.FuncMap{
	"json":          jsonify,
	"shellSanitize": terminal.NeutralizeControlCharacters,
	// TODO: Figure out what other functions we want to include here, if any.
}
