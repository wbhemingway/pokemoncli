package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mux      *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache:    make(map[string]cacheEntry),
		mux:      &sync.Mutex{},
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	data, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return data.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		delTime := time.Now().Add(-c.interval)
		c.mux.Lock()

		for key, val := range c.cache {
			if val.createdAt.Before(delTime) {
				delete(c.cache, key)
			}
		}
		c.mux.Unlock()
	}
}
