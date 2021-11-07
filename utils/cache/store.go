package cache

import (
	"time"
)

type updaterFunc func() ([]byte, bool)

type Store interface {
	// Get gets an item from the cache.
	Get(key string) ([]byte, bool)

	// Set sets an item in the cache with the given TTL.
	Set(key string, content []byte, duration time.Duration)

	// TTL returns the remaining TTL for an item.
	TTL(key string) int64

	// Expired returns whether a cached item has expired.
	Expired(key string) bool

	// PurgeExpired deletes any expired items from the cache.
	PurgeExpired() (purged uint)

	// Update triggers an update of a cached item using a mutator function.
	Update(key string, duration time.Duration, mutator updaterFunc) bool

	// WaitOrUpdate will wait for another routine to update an item or update it using the mutator function.
	WaitOrUpdate(key string, duration time.Duration, mutator updaterFunc) (stale bool)

	// Remove all items from the cache
	Flush()
}
