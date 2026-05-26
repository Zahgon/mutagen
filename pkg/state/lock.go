package state

import (
	"sync"
)

// TrackingLock provides locking facilities with automatic state tracking
// notifications.
type TrackingLock struct {
	// lock is the underlying mutex.
	lock sync.Mutex
	// tracker is the underlying tracker.
	tracker *Tracker
}

// NewTrackingLock creates a new tracking lock with the specified tracker.
func NewTrackingLock(tracker *Tracker) *TrackingLock { _ = "STUB: not implemented"; return nil }

// Lock locks the tracking lock.
func (l *TrackingLock) Lock() {
	_ = "STUB: not implemented"

	// Unlock unlocks the tracking lock and triggers a state update notification.
	return
}

func (l *TrackingLock) Unlock() { _ = "STUB: not implemented"; return }

// UnlockWithoutNotify unlocks the tracking lock without triggering a state
// update notification.
func (l *TrackingLock) UnlockWithoutNotify() { _ = "STUB: not implemented"; return }
