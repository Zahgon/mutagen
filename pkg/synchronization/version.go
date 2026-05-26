package synchronization

import (
	"github.com/mutagen-io/mutagen/pkg/filesystem"
	"github.com/mutagen-io/mutagen/pkg/filesystem/behavior"
	"github.com/mutagen-io/mutagen/pkg/synchronization/compression"
	"github.com/mutagen-io/mutagen/pkg/synchronization/core"
	"github.com/mutagen-io/mutagen/pkg/synchronization/core/ignore"
	"github.com/mutagen-io/mutagen/pkg/synchronization/hashing"
)

// DefaultVersion is the default session version.
const DefaultVersion Version = Version_Version1

// Supported indicates whether or not the session version is supported.
func (v Version) Supported() bool { _ = "STUB: not implemented"; return false }

// DefaultSynchronizationMode returns the default synchronization mode for the
// session version.
func (v Version) DefaultSynchronizationMode() core.SynchronizationMode {
	_ = "STUB: not implemented"
	return *new(core.SynchronizationMode)
}

// DefaultHashingAlgorithm returns the default hashing algorithm for the session
// version.
func (v Version) DefaultHashingAlgorithm() hashing.Algorithm {
	_ = "STUB: not implemented"
	return *new(hashing.Algorithm)
}

// DefaultMaximumEntryCount returns the default maximum entry count for the
// session version.
func (v Version) DefaultMaximumEntryCount() uint64 { _ = "STUB: not implemented"; return 0 }

// DefaultMaximumStagingFileSize returns the default maximum staging file size
// for the session version.
func (v Version) DefaultMaximumStagingFileSize() uint64 { _ = "STUB: not implemented"; return 0 }

// DefaultProbeMode returns the default probe mode for the session version.
func (v Version) DefaultProbeMode() behavior.ProbeMode {
	_ = "STUB: not implemented"
	return *new(behavior.ProbeMode)
}

// DefaultScanMode returns the default scan mode for the session version.
func (v Version) DefaultScanMode() ScanMode { _ = "STUB: not implemented"; return *new(ScanMode) }

// DefaultStageMode returns the default staging mode for the session version.
func (v Version) DefaultStageMode() StageMode { _ = "STUB: not implemented"; return *new(StageMode) }

// DefaultSymbolicLinkMode returns the default symbolic link mode for the
// session version.
func (v Version) DefaultSymbolicLinkMode() core.SymbolicLinkMode {
	_ = "STUB: not implemented"
	return *new(core.SymbolicLinkMode)
}

// DefaultWatchMode returns the default watch mode for the session version.
func (v Version) DefaultWatchMode() WatchMode { _ = "STUB: not implemented"; return *new(WatchMode) }

// DefaultWatchPollingInterval returns the default watch polling interval for
// the session version.
func (v Version) DefaultWatchPollingInterval() uint32 { _ = "STUB: not implemented"; return 0 }

// DefaultIgnoreSyntax returns the default ignore syntax for the session
// version.
func (v Version) DefaultIgnoreSyntax() ignore.Syntax {
	_ = "STUB: not implemented"
	return *new(ignore.Syntax)
}

// DefaultIgnoreVCSMode returns the default VCS ignore mode for the session
// version.
func (v Version) DefaultIgnoreVCSMode() ignore.IgnoreVCSMode {
	_ = "STUB: not implemented"
	return *new(ignore.IgnoreVCSMode)
}

// DefaultPermissionsMode returns the default permissions mode for the session
// version.
func (v Version) DefaultPermissionsMode() core.PermissionsMode {
	_ = "STUB: not implemented"
	// NOTE: Due to the hack listed in Configuration.EnsureValid (regarding the
	// computation of the default permissions mode), it would be advisable to
	// keep the default here the same for all session versions. If we want this
	// behavior to differ in the future, then we'd need to thread the session
	// version information into Configuration.EnsureValid, because the default
	// can affect the validation of default file and directory modes. This hack
	// could be replaced by looser validation on the default file and directory
	// modes, at least in the scenario where a default permissions mode is used
	// (which is most cases, unfortunately), but since we don't have any
	// foreseeable reason to change this default across future session versions,
	// we're best off keeping the stricter validation for now. We could also
	// change the signature of Configuration.EnsureValid to accept a session
	// version, but that rapidly spirals into other APIs and it's not even clear
	// how to enforce that the daemon's default session version is what's being
	// used for validation in the command line interface or external tools.
	return *new(core.PermissionsMode)
}

// DefaultFileMode returns the default file permission mode for the session
// version.
func (v Version) DefaultFileMode() filesystem.Mode {
	_ = "STUB: not implemented"
	return *new(filesystem.Mode)
}

// DefaultDirectoryMode returns the default directory permission mode for the
// session version.
func (v Version) DefaultDirectoryMode() filesystem.Mode {
	_ = "STUB: not implemented"
	return *new(filesystem.Mode)
}

// DefaultOwnerSpecification returns the default owner specification for the
// session version.
func (v Version) DefaultOwnerSpecification() string { _ = "STUB: not implemented"; return "" }

// DefaultGroupSpecification returns the default owner group specification for
// the session version.
func (v Version) DefaultGroupSpecification() string { _ = "STUB: not implemented"; return "" }

// DefaultCompressionAlgorithm returns the default compression algorithm for the
// session version.
func (v Version) DefaultCompressionAlgorithm() compression.Algorithm {
	_ = "STUB: not implemented"
	return *new(compression.Algorithm)
}
