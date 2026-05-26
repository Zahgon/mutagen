package multiplexing

import (
	"time"
)

// Configuration encodes multiplexer configuration.
type Configuration struct {
	// StreamReceiveWindow is the size (in bytes) of the stream receive window
	// (which also sets the size of the local per-stream receive buffer). If
	// less than or equal to 0, then no inbound data will be allowed and all
	// writes on the other end of the stream will block until write deadline
	// expiration or a call to CloseWrite or Close. The default value is 64 kB.
	StreamReceiveWindow int
	// WriteBufferCount is the number of write buffers to use within the
	// multiplexer. Each write buffer contains a fixed amount of internal
	// storage, currently 65548 bytes. If less than or equal to 0, then a single
	// write buffer will be created. The default is 5.
	WriteBufferCount int
	// AcceptBacklog is the maximum number of concurrent pending inbound open
	// requests that will be allowed. If less than or equal to 0, then it will
	// be set to 1. The default value is 10.
	AcceptBacklog int
	// HeartbeatTransmitInterval is the interval on which heartbeats will be
	// transmitted. If less than or equal to 0, then heartbeat transmission will
	// be disabled. The default interval is 5 seconds.
	HeartbeatTransmitInterval time.Duration
	// MaximumHeartbeatReceiveInterval is the maximum amount of time that the
	// multiplexer will be allowed to operate without receiving a heartbeat
	// message from the remote. If less than or equal to 0, remote heartbeats
	// will be processed but not required. The default interval is 10 seconds.
	MaximumHeartbeatReceiveInterval time.Duration
}

// DefaultConfiguration returns the default multiplexer configuration.
func DefaultConfiguration() *Configuration { _ = "STUB: not implemented"; return nil }

// 65535 bytes

// normalize normalizes out-of-range configuration values.
func (c *Configuration) normalize() { _ = "STUB: not implemented"; return }
