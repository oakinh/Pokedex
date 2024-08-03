package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	entries  map[string]cacheEntry
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		mu:       sync.Mutex{},
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := cacheEntry{createdAt: time.Now(), val: val}
	c.entries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for key, entry := range c.entries {
		afterInterval := entry.createdAt.Add(time.Second * c.interval)
		if entry.createdAt.After(afterInterval) {
			delete(c.entries, key)
		}
	}

}
