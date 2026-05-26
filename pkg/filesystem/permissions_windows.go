package filesystem

import (
	"golang.org/x/sys/windows"
)

// OwnershipSpecification is an opaque type that encodes specification of file
// and/or directory ownership.
type OwnershipSpecification struct {
	// ownerSID encodes the Windows owner SID associated with the ownership
	// specification. It may represent either a user SID or a group SID. A nil
	// value indicates the absence of specification.
	ownerSID *windows.SID
	// groupSid encodes the Windows group SID associated with the ownership
	// specification. A nil value indicates the absence of specification.
	groupSID *windows.SID
}

// NewOwnershipSpecification parsers owner and group specifications and resolves
// their system-level identifiers.
func NewOwnershipSpecification(owner, group string) (*OwnershipSpecification, error) {
	_ = "STUB: not implemented"
	// Attempt to parse and look up owner, if specified. On Windows, an owner
	// can be either a user or a group.
	return nil, nil
}

// Verify that this SID represents either a user or a group.

// Convert the retrieved SID to a string.

// Verify that this name represents either a user or a group and
// retrieve the associated SID.

// Convert the retrieved SID to a string.

// Attempt to parse and look up group, if specified.

// Success.

// SetPermissionsByPath sets the permissions on the content at the specified
// path. Ownership information is set first, followed by permissions extracted
// from the mode using ModePermissionsMask. Ownership setting can be skipped
// completely by providing a nil OwnershipSpecification or a specification with
// both components unset. An OwnershipSpecification may also include only
// certain components, in which case only those components will be set.
// Permission setting can be skipped by providing a mode value that yields 0
// after permission bit masking.
func SetPermissionsByPath(path string, ownership *OwnershipSpecification, mode Mode) error {
	_ = "STUB: not implemented"
	// Set ownership information, if specified.
	return nil
}

// Compute the information that we're going to set.

// Set the information.

// Set permissions, if specified.

// Success.
