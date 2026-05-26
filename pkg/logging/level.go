package logging

// Level represents a log level. Its value hierarchy is designed to be ordered
// and comparable by value.
type Level uint

const (
	// LevelDisabled indicates that logging is completely disabled.
	LevelDisabled Level = iota
	// LevelError indicates that only fatal errors are logged.
	LevelError
	// LevelWarn indicates that both fatal and non-fatal errors are logged.
	LevelWarn
	// LevelInfo indicates that basic execution information is logged (in
	// addition to all errors).
	LevelInfo
	// LevelDebug indicates that advanced execution information is logged (in
	// addition to basic information and all errors).
	LevelDebug
	// LevelTrace indicates that low-level execution information is logged (in
	// addition to all other execution information and all errors).
	LevelTrace
)

// levelNames are the human-readable representations of log levels.
var levelNames = [6]string{
	"disabled",
	"error",
	"warn",
	"info",
	"debug",
	"trace",
}

// String provides a human-readable representation of a log level.
func (l Level) String() string { _ = "STUB: not implemented"; return "" }

// NameToLevel converts a string-based representation of a log level to the
// appropriate Level value. It returns a boolean indicating whether or not the
// conversion was valid. If the name is invalid, LevelDisabled is returned.
func NameToLevel(name string) (Level, bool) { _ = "STUB: not implemented"; return *new(Level), false }

// abbreviations is the range of abbreviations to use for log levels.
const abbreviations = "_EWIDT"

// abbreviation returns a one-byte prefix to use for the level in log lines.
func (l Level) abbreviation() byte { _ = "STUB: not implemented"; return 0 }

// abbreviationToLevel converts a one-byte prefix representation of a log level
// to the appropriate Level value. It returns a boolean indicating whether or
// not the conversion was valid. If the abbreviation is invalid, LevelDisabled
// is returned.
func abbreviationToLevel(abbreviation byte) (Level, bool) {
	_ = "STUB: not implemented"
	return *new(Level), false
}
