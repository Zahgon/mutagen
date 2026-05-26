package forwarding

import (
	"github.com/mutagen-io/mutagen/pkg/filesystem"
)

// DefaultVersion is the default session version.
const DefaultVersion Version = Version_Version1

// Supported indicates whether or not the session version is supported.
func (v Version) Supported() bool { _ = "STUB: not implemented"; return false }

// DefaultSocketOverwriteMode returns the default socket overwrite mode for the
// session version.
func (v Version) DefaultSocketOverwriteMode() SocketOverwriteMode {
	_ = "STUB: not implemented"
	return *new(SocketOverwriteMode)
}

// DefaultSocketPermissionMode returns the default socket permission mode for
// the session version.
func (v Version) DefaultSocketPermissionMode() filesystem.Mode {
	_ = "STUB: not implemented"
	return *new(filesystem.Mode)
}

// DefaultSocketOwnerSpecification returns the default socket owner
// specification for the session version.
func (v Version) DefaultSocketOwnerSpecification() string { _ = "STUB: not implemented"; return "" }

// DefaultSocketGroupSpecification returns the default socket group
// specification for the session version.
func (v Version) DefaultSocketGroupSpecification() string { _ = "STUB: not implemented"; return "" }
