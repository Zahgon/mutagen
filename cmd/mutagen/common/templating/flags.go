package templating

import (
	"text/template"

	"github.com/spf13/pflag"
)

// TemplateFlags stores command line formatting flags and provides for their
// registration and handling.
type TemplateFlags struct {
	// template stores the value of the --template flag.
	template string
	// templateFile stores the value of the --template-file flag.
	templateFile string
}

// Register registers the flags into the specified flag set.
func (f *TemplateFlags) Register(flags *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// LoadTemplate loads the template specified by the flags. If no template has
// been specified, then it returns nil with no error. Template literals
// specified via the command line will have a trailing newline added.
func (f *TemplateFlags) LoadTemplate() (*template.Template, error) {
	_ = "STUB: not implemented"
	// Figure out if there's a template to be processed. If not, then no valid
	// template has been specified and we can just return. If a template literal
	// was provided directly on the command line, then add a trailing newline to
	// make typical command line usage more friendly.
	return nil, nil
}

// Create the template and register built-in functions.

// Parse the template literal.
