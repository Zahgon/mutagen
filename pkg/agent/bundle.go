package agent

const (
	// BundleName is the base name of the agent bundle.
	BundleName = "mutagen-agents.tar.gz"
)

// BundleLocation encodes an expected location for the agent bundle.
type BundleLocation uint8

const (
	// BundleLocationDefault indicates that the ExecutableForPlatform function
	// should expect to find the agent bundle in the same directory as the
	// current executable or (if the current executable resides in the "bin"
	// directory of a Filesystem Hierarchy Standard layout) in the libexec
	// directory.
	BundleLocationDefault BundleLocation = iota
	// BundleLocationBuildDirectory indicates that the ExecutableForPlatform
	// function should expect to find the agent bundle in the Mutagen build
	// directory. This mode is only used during integration testing. It is
	// required because test executables will be built in temporary directories.
	BundleLocationBuildDirectory
)

// ExpectedBundleLocation specifies the expected agent bundle location. It is
// set by the time that init functions have completed. After that, it should
// only be set at the process entry point, before any code calls into the agent
// package.
var ExpectedBundleLocation BundleLocation

// ExecutableForPlatform attempts to locate the agent bundle and extract an
// agent executable for the specified target platform. If no output path is
// specified, then the extracted file will be in a temporary location accessible
// to only the user, will have an appropriate extension for the target platform,
// and will have the executability bit set if it makes sense. The path to the
// extracted file will be returned, and the caller is responsible for cleaning
// up the file if this function returns a nil error.
func ExecutableForPlatform(goos, goarch, outputPath string) (string, error) {
	_ = "STUB: not implemented"
	// Compute the path to the location in which we expect to find the agent
	// bundle.
	return "", nil
}

// Add the executable directory as a search path.

// If the executable is in what appears to be a Filesystem Hierarchy
// Standard layout, then add the libexec directory as a search path.

// Loop until we find a bundle file. If we fail to locate a bundle, then
// abort. If we succeed, then defer its closure.

// Create a decompressor and defer its closure.

// Create an archive reader.

// Scan until we find a matching header.

// Check if we have a valid header. If not, there was no match.

// If an output path has been specified, then open the path for writing,
// otherwise create a temporary file.

// Copy data into the file.

// If we're not on Windows and our target system is not Windows, mark the
// file as executable. This will save us an additional "chmod +x" command
// during agent installation. Note that the mechanism we use here
// (os.File.Chmod) does not work on Windows (only the path-based os.Chmod is
// supported there), but this is fine because this code wouldn't make sense
// to use on Windows in any scenario (where executability bits don't exist).

// Close the file.

// Success.
