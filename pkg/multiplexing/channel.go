package multiplexing

import (
	"time"
)

// isClosed checks if a signaling channel is closed.
func isClosed(channel <-chan struct{}) bool { _ = "STUB: not implemented"; return false }

// wasPopulatedWithTime checks if a time signaling channel was populated with a
// time value and drains it if so.
func wasPopulatedWithTime(channel <-chan time.Time) bool { _ = "STUB: not implemented"; return false }
