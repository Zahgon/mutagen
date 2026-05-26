//go:build !darwin && !linux

package format

import (
	"github.com/mutagen-io/mutagen/pkg/filesystem"
)

// QueryByPath queries the filesystem format for the specified path.
func QueryByPath(_ string) (Format, error) { _ = "STUB: not implemented"; return *new(Format), nil }

// Query queries the filesystem format for the specified directory.
func Query(_ *filesystem.Directory) (Format, error) {
	_ = "STUB: not implemented"
	return *new(Format), nil
}
