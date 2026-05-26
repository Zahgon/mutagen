package format

import (
	"golang.org/x/sys/unix"
)

const (
	// FormatAPFS represents an APFS filesystem format.
	FormatAPFS Format = iota + 1
	// FormatHFS represents an HFS (or variant thereof) filesystem format.
	FormatHFS
	// FormatFAT32 represents a FAT32 filesystem format.
	FormatFAT32
	// FormatExFAT represents a ExFAT filesystem format.
	FormatExFAT
)

// metadataRepresentsAPFS returns whether or not the specified filesystem
// metadata represents an APFS filesystem.
func metadataRepresentsAPFS(metadata *unix.Statfs_t) bool { _ = "STUB: not implemented"; return false }

// metadataRepresentsHFS returns whether or not the specified filesystem
// metadata represents an HFS filesystem. This also covers HFS variants.
func metadataRepresentsHFS(metadata *unix.Statfs_t) bool { _ = "STUB: not implemented"; return false }

// metadataRepresentsFAT32 returns whether or not the specified filesystem
// metadata represents a FAT32 filesystem.
func metadataRepresentsFAT32(metadata *unix.Statfs_t) bool { _ = "STUB: not implemented"; return false }

// metadataRepresentsExFAT returns whether or not the specified filesystem
// metadata represents a ExFAT filesystem.
func metadataRepresentsExFAT(metadata *unix.Statfs_t) bool { _ = "STUB: not implemented"; return false }

// formatFromStatfs extracts the filesystem format from the filesystem metadata.
func formatFromStatfs(metadata *unix.Statfs_t) Format {
	_ = "STUB: not implemented"
	// Check if this is a well-known filesystem format.
	return *new(Format)
}

// Otherwise classify it as unknown.
