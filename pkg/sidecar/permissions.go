package sidecar

import (
	"github.com/mutagen-io/mutagen/pkg/filesystem"
)

// SetVolumeOwnershipAndPermissionsIfEmpty will set the ownership and
// permissions on a sidecar volume if (and only if) the volume is empty.
func SetVolumeOwnershipAndPermissionsIfEmpty(name string, ownership *filesystem.OwnershipSpecification, mode filesystem.Mode) error {
	_ = "STUB: not implemented"
	// Open the volumes directory and defer its closure.
	return nil
}

// Open the volume mount point and defer its closure.

// Check if the volume is empty. If not, then we're done.

// Set permissions on the volume.
