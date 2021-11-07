package services

import (
	"time"

	"github.com/synycboom/bsc-evm-compatible-bridge-api/utils/cache"
)

// CacheService clears out stale entries from the cache, to keep it tidy and fast
type cacheService struct {
	BaseService

	ticker *time.Ticker
	cache  cache.Store
}

func NewCacheService(cache cache.Store, interval time.Duration,
) *cacheService {
	ticker := time.NewTicker(interval)
	ps := &cacheService{
		ticker: ticker,
		cache:  cache,
	}

	ps.BaseService = *NewBaseService(nil, "CacheService", ps)
	return ps
}

// Start the service.
// If it's already started or stopped, will return an error.
// If OnStart() returns an error, it's returned by Start()
// implements BaseService
func (ps *cacheService) OnStart() error {
	log.Info("CacheService starting.")
	go func() {
		for t := range ps.ticker.C {
			ps.Tick(t)
		}
	}()
	ps.Tick(time.Now())
	return nil
}

// Stop the service.
// If it's already stopped, will return an error.
// OnStop must never error.
// implements BaseService
func (ps *cacheService) OnStop() {
	log.Info("CacheService stopping.")
	ps.ticker.Stop()
}

// Reset the service.
// Panics by default - must be overwritten to enable reset.
// implements BaseService
func (ps *cacheService) OnReset() error {
	log.Info("CacheService resetting.")
	return nil
}

// Tick is called each `interval` (typically 1h)
func (ps *cacheService) Tick(t time.Time) {
	log.Debug("CacheService tick.")

	purged := ps.cache.PurgeExpired()
	log.Debugf("CacheService purged %d stale items this round.", purged)
}
