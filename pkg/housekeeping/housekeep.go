package housekeeping

import (
	"time"
)

const (
	// maximumAgentIdlePeriod is the maximum period of time that an agent binary
	// is allowed to sit on disk without being executed before being deleted.
	maximumAgentIdlePeriod = 30 * 24 * time.Hour
	// maximumCacheAge is the maximum allowed cache age.
	maximumCacheAge = 7 * 24 * time.Hour
	// maximumStagingRootAge is the maximum allowed staging root age.
	maximumStagingRootAge = 7 * 24 * time.Hour
)

// Housekeep invokes housekeeping functions on the Mutagen data directory.
func Housekeep() {
	_ = "STUB: not implemented"
	// Perform housekeeping on agent binaries, but only if we're not in a
	// Mutagen sidecar container. Sidecar containers are particularly
	// susceptible to stale agent access times due to the fact that the agent is
	// baked into the sidecar image and the sidecar image is typically unpacked
	// via OverlayFS on top of ext4 with either relatime or noatime.
	return
}

// Perform housekeeping on caches.

// Perform housekeeping on staging roots.

// housekeepAgents performs housekeeping of agent binaries.
func housekeepAgents() {
	_ = "STUB: not implemented"
	// Compute the path to the agents directory. If we fail, just abort. We
	// don't attempt to create the directory, because if it doesn't exist, then
	// we don't need to do anything and we'll just bail when we fail to list the
	// agent directory below.
	return
}

// Get the list of locally installed agent versions. If we fail, just abort.

// Compute the name of the agent binary.

// Grab the current time.

// Loop through each agent version, compute the time it was last launched,
// and remove it if longer than the maximum allowed period. Skip contents
// where failures are encountered.

// TODO: Ensure that the name matches the expected format. Be mindful of
// the fact that it might contain a tag.

// housekeepCaches performs housekeeping of caches.
func housekeepCaches() {
	_ = "STUB: not implemented"
	// Compute the path to the caches directory. If we fail, just abort. We
	// don't attempt to create the directory, because if it doesn't exist, then
	// we don't need to do anything and we'll just bail when we fail to list the
	// caches directory contents below.
	// TODO: Move this logic into paths.go? Need to keep it in sync with
	// pathForCache.
	return
}

// Get the list of caches. If we fail, just abort.

// Grab the current time.

// Loop through each cache and remove those older than a certain age. Ignore
// any failures.

// housekeepStaging performs housekeeping of staging roots.
func housekeepStaging() {
	_ = "STUB: not implemented"
	// Compute the path to the staging directory (the top-level directory
	// containing all staging roots). If we fail, just abort. We don't attempt
	// to create the directory, because if it doesn't exist, then we don't need
	// to do anything and we'll just bail when we fail to list the staging
	// directory contents below.
	// TODO: Move this logic into paths.go? Need to keep it in sync with
	// pathForStagingRoot and pathForStaging.
	return
}

// Get the list of staging roots. If we fail, just abort.

// Grab the current time.

// Loop through each staging root and remove those older than a certain
// age. Ignore any failures. This is a little bit more cavalier than cache
// housekeeping because removal is non-atomic and theoretically a given
// staging root could be in use. However, a session's staging root is wiped
// on each successful synchronization cycle, so by using a large maximum
// staging root age, we're only going to run into trouble if the staging
// portion of a synchronization cycle starts up, after having failed a long
// time ago, at the precise moment that we're housekeeping. In that case, it
// would try to use the existing staging directory from the failed
// synchronization cycle and there might be a conflict. But even in that
// statistically unlikely case, the worst case scenario would be triggering
// an additional synchronization cycle.
