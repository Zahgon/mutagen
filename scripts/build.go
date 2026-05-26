package main

import (
	"archive/tar"
	"os"

	"github.com/klauspost/compress/gzip"

	"github.com/mutagen-io/mutagen/cmd"
)

const (
	// agentPackage is the Go package URL to use for building Mutagen agent
	// binaries.
	agentPackage = "github.com/mutagen-io/mutagen/cmd/mutagen-agent"
	// cliPackage is the Go package URL to use for building Mutagen binaries.
	cliPackage = "github.com/mutagen-io/mutagen/cmd/mutagen"

	// agentBuildSubdirectoryName is the name of the build subdirectory where
	// agent binaries are built.
	agentBuildSubdirectoryName = "agent"
	// cliBuildSubdirectoryName is the name of the build subdirectory where CLI
	// binaries are built.
	cliBuildSubdirectoryName = "cli"
	// releaseBuildSubdirectoryName is the name of the build subdirectory where
	// release bundles are built.
	releaseBuildSubdirectoryName = "release"

	// agentBaseName is the name of the Mutagen agent binary without any path or
	// extension.
	agentBaseName = "mutagen-agent"
	// cliBaseName is the name of the Mutagen binary without any path or
	// extension.
	cliBaseName = "mutagen"

	// minimumMacOSVersion is the minimum version of macOS that
	// we'll support. This is pinned to the oldest version of
	// macOS that Mutagen's minimum Go version supports. Go 1.25
	// requires macOS 12 Monterey or later.
	minimumMacOSVersion = "12"

	// minimumARMSupport is the value to pass to the GOARM environment variable
	// when building binaries. We currently specify support for ARMv5. This will
	// enable software-based floating point. For our use case, this is totally
	// fine, because we don't have any floating-point-heavy code, and the
	// resulting binary bloat is very minimal. This won't apply for arm64, which
	// always has hardware-based floating point support. For more information,
	// see: https://github.com/golang/go/wiki/GoArm
	minimumARMSupport = "5"
)

// Target specifies a GOOS/GOARCH combination.
type Target struct {
	// GOOS is the GOOS environment variable specification for the target.
	GOOS string
	// GOARCH is the GOARCH environment variable specification for the target.
	GOARCH string
}

// String generates a human-readable representation of the target.
func (t Target) String() string { _ = "STUB: not implemented"; return "" }

// Name generates a representation of the target that is suitable for paths and
// file names.
func (t Target) Name() string { _ = "STUB: not implemented"; return "" }

// ExecutableName formats executable names for the target.
func (t Target) ExecutableName(base string) string {
	_ = "STUB: not implemented"
	// If we're on Windows, append a ".exe" extension.
	return ""
}

// Otherwise return the base name unmodified.

// appendGoEnv modifies an environment specification to make the Go toolchain
// generate output for the target. It assumes that the resulting environment
// will be used with os/exec.Cmd and thus doesn't avoid duplicate variables.
func (t Target) appendGoEnv(environment []string) []string {
	_ = "STUB: not implemented"
	// Override GOOS/GOARCH.
	return nil
}

// If we're building a macOS binary on macOS, then we enable cgo because
// we'll need it to access the FSEvents API. We have to enable it explicitly
// because Go won't enable it when cross compiling between different Darwin
// architectures. We also need to tell the C compiler and external linker to
// support older versions of macOS. These flags will tell the C compiler to
// generate code compatible with the target version of macOS and tell the
// external linker what value to embed for the LC_VERSION_MIN_MACOSX flag in
// the resulting Mach-O binaries. Go's internal linker automatically
// defaults to a relatively liberal (old) value for this flag, but since
// we're using an external linker, it defaults to the current SDK version.
//
// For all other platforms, we disable cgo. This is essential for our Linux
// CI setup, because we build agent executables during testing that we then
// run inside Docker containers for our integration tests. These containers
// typically run Alpine Linux, and if the agent binary is linked to C
// libraries that only exist on the build system, then they won't work
// inside the container. We can't disable cgo on a global basis though,
// because it's needed for race condition testing. Another reason that it's
// good to disable cgo when building agent binaries during testing is that
// the release agent binaries will also have cgo disabled (except on macOS),
// and we'll want to faithfully recreate that.

// Set up ARM target support. See notes for definition of minimumARMSupport.
// We don't need to unset any existing GOARM variables since they simply
// won't be used if we're not targeting (non-64-bit) ARM systems.

// Done.

// IsCrossTarget determines whether or not the target represents a
// cross-compilation target (i.e. not the native target for the current Go
// toolchain).
func (t Target) IsCrossTarget() bool { _ = "STUB: not implemented"; return false }

// IncludeAgentInSlimBuildModes indicates whether or not the target should have
// an agent binary included in the agent bundle in slim and release-slim modes.
func (t Target) IncludeAgentInSlimBuildModes() bool { _ = "STUB: not implemented"; return false }

// BuildBundleInReleaseSlimMode indicates whether or not the target should have
// a release bundle built in release-slim mode.
func (t Target) BuildBundleInReleaseSlimMode() bool { _ = "STUB: not implemented"; return false }

// Build executes a module-aware build of the specified package URL, storing the
// output of the build at the specified path.
func (t Target) Build(url, output string, enableSSPLEnhancements, disableDebug bool) error {
	_ = "STUB: not implemented"
	// Compute the build command. If we don't need debugging, then we use the -s
	// and -w linker flags to omit the symbol table and debugging information.
	// This shaves off about 25% of the binary size and only disables debugging
	// (stack traces are still intact). For more information, see:
	// https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick
	// In this case, we also trim the code paths stored in the executable, as
	// there's no use in having the full paths available.
	return nil
}

// Create the build command.

// Set the environment.

// Forward input, output, and error streams.

// Run the build.

// targets encodes which combinations of GOOS and GOARCH we want to use for
// building agent and CLI binaries. We don't build every target at the moment,
// but we do list all potential targets here and comment out those we don't
// support. This list is created from https://golang.org/doc/install/source.
// Unfortunately there's no automated way to construct this list, but that's
// fine since we have to manually groom it anyway.
var targets = []Target{
	// Define AIX targets.
	{"aix", "ppc64"},

	// Define Android targets. We disable support for Android since it doesn't
	// have a clearly defined use case as a target platform, though there might
	// be certain development scenarios where it would make sense as an endpoint
	// (via a third-party SSH server on the device).
	// {"android", "386"},
	// {"android", "amd64"},
	// {"android", "arm"},
	// {"android", "arm64"},

	// Define macOS targets.
	{"darwin", "amd64"},
	{"darwin", "arm64"},

	// Define DragonFlyBSD targets.
	{"dragonfly", "amd64"},

	// Define FreeBSD targets.
	{"freebsd", "386"},
	{"freebsd", "amd64"},
	{"freebsd", "arm"},
	// The freebsd/arm64 port was added in Go 1.14, but for some reason isn't
	// documented at https://golang.org/doc/install/source.
	{"freebsd", "arm64"},

	// Define illumos targets. We disable explicit support for illumos because
	// it's already effectively supported by our Solaris target. illumos is (at
	// least for Mutagen's purposes) an ABI-compatible superset of Solaris, so
	// there's no need for a separate build. Within the Go toolchain, runtime,
	// and standard library, most of illumos' support is provided by the Solaris
	// port. The "illumos" target even implies the "solaris" build constraint.
	// As such, the Solaris binaries should work fine for illumos distributions.
	// Also, the uname command on illumos returns the same kernel name ("SunOS")
	// as Solaris, so our probing wouldn't be able to identify illumos anyway.
	// {"illumos", "amd64"},

	// Define iOS/iPadOS/watchOS/tvOS targets. We disable support for these
	// since they don't make sense as target platforms.
	// The ios/amd64 port was added in Go 1.16, but for some reason isn't
	// documented at https://golang.org/doc/install/source.
	// {"ios", "amd64"},
	// {"ios", "arm64"},

	// Define WebAssembly targets. We disable support for WebAssembly since it
	// doesn't make sense as a target platform.
	// {"js", "wasm"},
	// {"wasip1", "wasm"},

	// Define Linux targets.
	{"linux", "386"},
	{"linux", "amd64"},
	{"linux", "arm"},
	{"linux", "arm64"},
	// TODO: Assess whether or not we want to support LoongArch. Support was
	// added in Go 1.19, but it sounds like most real-world deployments use a
	// Linux kernel that's too old to support binaries compiled by the official
	// Go toolchain. If this situation changes, then it's certainly worth
	// enabling support. The code does build successfully on this architecture.
	// In this case, we'll also need to update platform detection with the
	// appropriate uname -m value.
	// {"linux", "loong64"},
	{"linux", "mips"},
	{"linux", "mipsle"},
	{"linux", "mips64"},
	{"linux", "mips64le"},
	{"linux", "ppc64"},
	{"linux", "ppc64le"},
	{"linux", "riscv64"},
	{"linux", "s390x"},

	// Define NetBSD targets.
	{"netbsd", "386"},
	{"netbsd", "amd64"},
	{"netbsd", "arm"},
	// The netbsd/arm64 port was added in Go 1.16, but for some reason isn't
	// documented at https://golang.org/doc/install/source.
	{"netbsd", "arm64"},

	// Define OpenBSD targets.
	{"openbsd", "386"},
	{"openbsd", "amd64"},
	{"openbsd", "arm"},
	{"openbsd", "arm64"},
	// The openbsd/mips64 port was added in Go 1.16, but for some reason isn't
	// documented at https://golang.org/doc/install/source. It is currently
	// disabled on https://build.golang.org/ due to https://go.dev/issue/36435.
	// Moreover, it seems to be broken when using the Go sys subrepository after
	// v0.1.0 - the Go linker crashes with a segfault (possibly due to the
	// aforementioned issue). Until the picture around this port clarifies a
	// bit, it's not worth supporting.
	// {"openbsd", "mips64"},
	// The openbsd/ppc64 port was added in Go 1.22, but for some reason isn't
	// documented at https://golang.org/doc/install/source. Let's wait and see
	// if there's sufficient demand for it. It is still considered experimental.
	// {"openbsd", "ppc64"},
	// The openbsd/riscv64 port was added in Go 1.23, but for some reason isn't
	// documented at https://golang.org/doc/install/source. Let's wait and see
	// if there's sufficient demand for it. It is still considered experimental.
	// {"openbsd", "riscv64"},

	// Define Plan 9 targets. We disable support for Plan 9 because it's missing
	// too many system calls and other APIs necessary for Mutagen to build. It
	// might make sense to support Plan 9 as an endpoint for certain development
	// scenarios, but it will take a significant amount of work just to get the
	// Mutagen agent to build.
	// {"plan9", "386"},
	// {"plan9", "amd64"},
	// {"plan9", "arm"},

	// Define Solaris targets.
	{"solaris", "amd64"},

	// Define Windows targets. The windows/arm port was removed
	// in Go 1.26.
	{"windows", "386"},
	{"windows", "amd64"},
	{"windows", "arm64"},
}

// macOSCodeSign performs macOS code signing on the specified path using the
// specified signing identity. It performs code signing in a manner suitable for
// later submission to Apple for notarization.
func macOSCodeSign(path, identity string) error {
	_ = "STUB: not implemented"
	// Create the code signing command.
	//
	// We include the --force flag because the Go toolchain won't touch binaries
	// if they don't need to be rebuilt and thus we might have a signature from
	// a previous build. In that case, the code signing operation will fail
	// without --force. When --force is specified, any existing signature will
	// be overwritten, unless it's using the same code signing identity, in
	// which case it will simply be left in place (which is actually optimal for
	// for repeated local usage). Note that the --force flag is not required to
	// override ad hoc signatures (which the Go toolchain will add by default
	// darwin/arm64 binaries).
	//
	// The --options runtime and --timestamp flags are required to enable the
	// hardened runtime (which doesn't affect Mutagen binaries) and to use a
	// secure signing timestamp, both of which are required for notarization.
	return nil
}

// Forward input, output, and error streams.

// Run code signing.

// archiveBuilderCopyBufferSize determines the size of the copy buffer used when
// generating archive files.
// TODO: Figure out if we should set this on a per-machine basis. This value is
// taken from Go's io.Copy method, which defaults to allocating a 32k buffer if
// none is provided.
const archiveBuilderCopyBufferSize = 32 * 1024

type ArchiveBuilder struct {
	file       *os.File
	compressor *gzip.Writer
	archiver   *tar.Writer
	copyBuffer []byte
}

func NewArchiveBuilder(bundlePath string) (*ArchiveBuilder, error) {
	_ = "STUB: not implemented"
	// Open the underlying file.
	return nil, nil
}

// Create the compressor.

// Success.

func (b *ArchiveBuilder) Close() error {
	_ = "STUB: not implemented"
	// Close in the necessary order to trigger flushes.
	return nil
}

// Success.

func (b *ArchiveBuilder) Add(name, path string, mode int64) error {
	_ = "STUB: not implemented"
	// If the name is empty, use the base name.
	return nil
}

// Open the file and ensure its cleanup.

// Compute its size.

// Write the header for the entry.

// Copy the file contents.

// Success.

// copyFile copies the contents at sourcePath to a newly created file at
// destinationPath that inherits the permissions of sourcePath.
func copyFile(sourcePath, destinationPath string) error {
	_ = "STUB: not implemented"
	// Open the source file and defer its closure.
	return nil
}

// Grab source file metadata.

// Remove the destination.

// Create the destination file and defer its closure. We open with exclusive
// creation flags to ensure that we're the ones creating the file so that
// its permissions are set correctly.

// Copy contents.

// Success.

var usage = `usage: build [-h|--help] [-m|--mode=<mode>] [--sspl]
       [--macos-codesign-identity=<identity>]

The mode flag accepts four values: 'local', 'slim', 'release', and
'release-slim'. 'local' will build CLI and agent binaries only for the current
platform. 'slim' will build the CLI binary for only the current platform and
agents for a common subset of platforms. 'release' will build CLI and agent
binaries for all platforms and package for release. 'release-slim' is the same
as release but only builds release bundles for a small subset of platforms. The
default mode is 'slim'.

If --sspl is specified, then SSPL-licensed enhancements will be included in the
build output. By default, only MIT-licensed code is included in builds.

If --macos-codesign-identity specifies a non-empty value, then it will be used
to perform code signing on all macOS binaries in a fashion suitable for
notarization by Apple. The codesign utility must be able to access the
associated certificate and private keys in Keychain Access without a password if
this script is operated in a non-interactive mode.
`

// build is the primary entry point.
func build() error {
	_ = "STUB: not implemented"
	// Parse command line arguments.
	return nil
}

// The only platform really suited to cross-compiling for every other
// platform at the moment is macOS. This is because FSEvents is used for
// file monitoring and that is a C-based API, not accessible purely via
// system calls. All of the other platforms can operate with pure Go
// compilation.

// If a macOS code signing identity has been specified, then make sure we're
// in a mode where that makes sense.

// Compute the path to the Mutagen source directory.

// Verify that we're running inside the Mutagen source directory, otherwise
// we can't rely on Go modules working.

// Compute the path to the build directory and ensure that it exists.

// Create the necessary build directory hierarchy.

// Compute the local target.

// Compute agent targets.

// Compute CLI targets.

// Determine whether or not to disable debugging information in binaries.
// Doing so saves significant space, but is only suited to release builds.

// Build agent binaries.

// Build CLI binaries.

// Build the agent bundle.

// Build release bundles if necessary.

// Update status.

// Compute paths.

// Build the release bundle.

// Relocate the CLI binary for the current platform.

// Success.

func main() {
	if err := build(); err != nil {
		cmd.Fatal(err)
	}
}
