package cache

import (
	"runtime/debug"
	"sync"
	"time"

	logger "github.com/synycboom/bsc-evm-compatible-bridge-api/utils/log"
)

var _ Store = &MemoryStore{}

type MemoryStore struct {
	items  map[string]Item
	ulocks map[string]*sync.WaitGroup
	mtx    *sync.RWMutex
	umtx   *sync.RWMutex
}

// NewMemStorage creates a new in-memory storage
func NewMemStorage() Store {
	return &MemoryStore{
		items:  make(map[string]Item),
		ulocks: make(map[string]*sync.WaitGroup),
		mtx:    &sync.RWMutex{},
		umtx:   &sync.RWMutex{},
	}
}

// Get a cached content by key
func (ms *MemoryStore) Get(key string) ([]byte, bool) {
	ms.mtx.RLock()

	if item, ok := ms.items[key]; ok {
		ms.mtx.RUnlock()
		if item.Expired() {
			ms.mtx.Lock()
			delete(ms.items, key)
			logger.Logger.Debugf("MemoryStore: Expiring stale item with key: %s", key)
			ms.mtx.Unlock()
			return nil, false
		}
		return item.Content, true
	} else {
		ms.mtx.RUnlock()
	}
	return nil, false
}

// Set a cached content by key
func (ms *MemoryStore) Set(key string, content []byte, duration time.Duration) {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()

	ms.items[key] = Item{
		Content:    content,
		Expiration: time.Now().Add(duration).UnixNano(),
	}

	logger.Logger.Debugf("MemoryStore: Set a new item with key: %s", key)
}

func (ms *MemoryStore) TTL(key string) int64 {
	ms.mtx.RLock()
	defer ms.mtx.RUnlock()

	if item, ok := ms.items[key]; ok {
		if item.Expired() {
			return -1
		}
		return (item.Expiration - time.Now().UnixNano()) /
			time.Millisecond.Nanoseconds()
	}
	return -1
}

func (ms *MemoryStore) Expired(key string) bool {
	return ms.TTL(key) <= 0
}

func (ms *MemoryStore) PurgeExpired() (purged uint) {
	ms.mtx.Lock()
	defer ms.mtx.Unlock()
	for key, val := range ms.items {
		if val.Expired() {
			delete(ms.items, key)
			purged++
		}
	}
	return purged
}

func (ms *MemoryStore) Update(key string, duration time.Duration, updater updaterFunc) bool {
	ms.umtx.RLock()
	if _, ok := ms.ulocks[key]; ok {
		ms.umtx.RUnlock()
		return false
	}
	ms.umtx.RUnlock()
	ms.umtx.Lock()
	wg := &sync.WaitGroup{}
	ms.ulocks[key] = wg
	ms.umtx.Unlock()
	defer func() {
		ms.umtx.Lock()
		delete(ms.ulocks, key)
		ms.umtx.Unlock()
	}()
	wg.Add(1)
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Errorf("MemoryStore: Panic when update key %s, stack %s", key, string(debug.Stack()))
		}
	}()
	if content, updated := updater(); updated {
		if len(content) > 0 {
			ms.Set(key, content, duration)
		}
	}
	return true
}

// MemoryStore can tolerate with two request update the same key, which is rarely happened but make fine grit lock
func (ms *MemoryStore) WaitOrUpdate(key string, duration time.Duration, updater updaterFunc) (didUpdate bool) {
	ms.umtx.RLock()

	if wg, ok := ms.ulocks[key]; ok {
		ms.umtx.RUnlock()
		wg.Wait()
	} else { // first one here; Update will create WaitGroup
		ms.umtx.RUnlock()
	}

	// if stale is true, no other routines updated the cache & we should do that
	if stale := ms.TTL(key) <= 0; stale {
		didUpdate = ms.Update(key, duration, updater)
	}

	return didUpdate
}

func (ms *MemoryStore) Flush() {
	ms.mtx.Lock()
	ms.umtx.Lock()
	defer ms.mtx.Unlock()
	defer ms.umtx.Unlock()
	ms.items = make(map[string]Item)
	ms.ulocks = make(map[string]*sync.WaitGroup)
}
