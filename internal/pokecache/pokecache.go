package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	m        map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		m:        make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	c.m[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mutex.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	entry, ok := c.m[key]
	c.mutex.Unlock()

	if ok {
		return entry.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.m {
			if time.Now().Sub(entry.createdAt) > c.interval {
				delete(c.m, key)
			}
		}
		c.mutex.Unlock()
	}
}
