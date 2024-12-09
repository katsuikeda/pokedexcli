package pokecache

import (
	"log"
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu           *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, ok := c.cacheEntries[key]
	return entry.val, ok
}

func NewCache(interval time.Duration) *Cache {
	if interval <= 0 {
		log.Fatalln("caching interval must be greater than 0")
	}
	cache := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		mu:           &sync.RWMutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, entry := range c.cacheEntries {
		if time.Since(entry.createdAt) > interval {
			delete(c.cacheEntries, key)
		}
	}
}
