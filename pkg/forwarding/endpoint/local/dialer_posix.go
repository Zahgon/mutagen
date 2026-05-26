//go:build !windows

package local

import (
	"context"
	"net"
)

// dialWindowsNamedPipe returns an "unsupported" error on POSIX systems.
func dialWindowsNamedPipe(_ context.Context, _ string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
