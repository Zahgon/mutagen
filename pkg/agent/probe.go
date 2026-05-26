package agent

// unameSToGOOS maps uname -s output values to their corresponding GOOS values.
// Although some Windows environments (Cygwin, MSYS, and MinGW) support uname,
// their values are handled by unameSIsWindowsPosix because they are so varied
// (their value depends on the POSIX environment and its version, the system
// architecture, and the NT kernel version).
var unameSToGOOS = map[string]string{
	"AIX":       "aix",
	"Darwin":    "darwin",
	"DragonFly": "dragonfly",
	"FreeBSD":   "freebsd",
	"Linux":     "linux",
	"NetBSD":    "netbsd",
	"OpenBSD":   "openbsd",
	"SunOS":     "solaris",
	// TODO: Add more obscure uname -s values as necessary, e.g.
	// debian/kFreeBSD, which returns "GNU/kFreeBSD".
}

// unameSIsWindowsPosix determines whether or not a uname -s output value
// represents a Windows POSIX environment.
func unameSIsWindowsPosix(value string) bool { _ = "STUB: not implemented"; return false }

// unameMToGOARCH maps uname -m output values to their corresponding GOARCH
// values.
var unameMToGOARCH = map[string]string{
	"i386":     "386",
	"i486":     "386",
	"i586":     "386",
	"i686":     "386",
	"x86_64":   "amd64",
	"amd64":    "amd64",
	"armv5l":   "arm",
	"armv6l":   "arm",
	"armv7l":   "arm",
	"armv8l":   "arm64",
	"aarch64":  "arm64",
	"arm64":    "arm64",
	"mips":     "mips",
	"mipsel":   "mipsle",
	"mips64":   "mips64",
	"mips64el": "mips64le",
	"ppc64":    "ppc64",
	"ppc64le":  "ppc64le",
	"riscv64":  "riscv64",
	"s390x":    "s390x",
	// TODO: Add any more obscure uname -m variations that we might encounter.
}

// osEnvToGOOS maps the value of the "OS" environment variable on Windows to the
// corresponding GOOS. There's only one supported value, but we keep things this
// way for symmetry and extensibility.
var osEnvToGOOS = map[string]string{
	"Windows_NT": "windows",
}

// processorArchitectureEnvToGOARCH maps the value of the
// "PROCESSOR_ARCHITECTURE" environment variable on Windows to the corresponding
// GOARCH.
var processorArchitectureEnvToGOARCH = map[string]string{
	"x86":   "386",
	"AMD64": "amd64",
	"ARM":   "arm",
	"ARM64": "arm64",
	// TODO: Add IA64 (that's the key) if Go ever supports Itanium, though
	// they've pretty much stated that this will never happen:
	// https://groups.google.com/forum/#!topic/golang-nuts/RgGF1Dudym4
}

// probePOSIX performs platform probing over an agent transport, working under
// the assumption that the remote system is a POSIX system.
func probePOSIX(transport Transport) (string, string, error) {
	_ = "STUB: not implemented"
	// Try to invoke uname and print kernel and machine name.
	return "", "", nil
}

// Parse uname output.

// Translate GOOS. Windows POSIX systems typically include their NT version
// number in their uname -s output, so we have to handle those specially.

// Translate GOARCH. On AIX systems, uname -m returns the machine's serial
// number, but modern AIX only runs on 64-bit PowerPC anyway (and that's
// all we support), so we set the architecture directly.

// Success.

// probeWindows performs platform probing over an agent transport, working under
// the assumption that the remote system is a Windows system.
func probeWindows(transport Transport) (string, string, error) {
	_ = "STUB: not implemented"
	// Attempt to dump the remote environment.
	return "", "", nil
}

// Parse the output block into a series of KEY=value specifications.

// Extract the OS and PROCESSOR_ARCHITECTURE environment variables.

// Translate to GOOS.

// Translate to GOARCH.

// Success.

// probe attempts to identify the properties of the target platform (namely
// GOOS, GOARCH, and whether or not it's a POSIX environment (which it might be
// even on Windows)) using the specified transport.
func probe(transport Transport, prompter string) (string, string, bool, error) {
	_ = "STUB: not implemented"
	// Attempt to probe for a POSIX platform. This might apply to certain
	// Windows environments as well.
	return "", "", false, nil
}

// If that fails, attempt a Windows fallback.

// Failure.
