//go:build !windows

package filesystem

// OwnershipSpecification is an opaque type that encodes specification of file
// and/or directory ownership.
type OwnershipSpecification struct {
	// ownerID encodes the POSIX user ID associated with the ownership
	// specification. A value of -1 indicates the absence of specification. The
	// availability of -1 as a sentinel value for omission is guaranteed by the
	// POSIX definition of chmod.
	ownerID int
	// groupID encodes the POSIX user ID associated with the ownership
	// specification. A value of -1 indicates the absence of specification. The
	// availability of -1 as a sentinel value for omission is guaranteed by the
	// POSIX definition of chmod.
	groupID int
}

// NewOwnershipSpecification parsers owner and group specifications and resolves
// their system-level identifiers.
func NewOwnershipSpecification(owner, group string) (*OwnershipSpecification, error) {
	_ = "STUB: not implemented"
	// Attempt to parse and look up owner user, if specified.
	return nil, nil
}

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

// Set permissions, if specified.

// Success.
