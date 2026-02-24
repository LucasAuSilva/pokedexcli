package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	val 			[]byte
	createdAt time.Time
}

type Cache struct {
	mu sync.Mutex
	entries map[string]cacheEntry
	interval time.Duration
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if len(c.entries) < 1 {
		return
	}
	for k, e := range c.entries {
		if time.Since(e.createdAt) > c.interval {
			delete(c.entries, k)
		} 
	} 
}

func NewCache(interval time.Duration) *Cache {
	entries := make(map[string]cacheEntry)

	ticker := time.NewTicker(interval)

	cache := Cache{
		entries: entries,
		interval: interval,
	}

	go func() {
		for range ticker.C {
			cache.reapLoop()
		}
	}()

	return &cache
}
