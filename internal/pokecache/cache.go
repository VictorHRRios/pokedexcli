package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	mu       *sync.RWMutex
	interval time.Duration
}

func NewCache(duration time.Duration) Cache {
	c := Cache{
		interval: duration,
		mu:       &sync.RWMutex{},
		cache:    map[string]cacheEntry{},
	}
	c.ReapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.cache[key]
	if !ok {
		c.cache[key] = cacheEntry{time.Now(), val}
	} else {
		fmt.Println("Value already exists")
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.cache[key]
	if !ok {
		fmt.Println("Value does not exists")
		return nil, ok
	}
	return val.val, ok
}

func (c *Cache) ReapLoop() {
	ticker := time.NewTicker(c.interval)
	go func() {
		for range ticker.C {
			c.mu.Lock()
			now := time.Now()
			for key, val := range c.cache {
				if now.Sub(val.createdAt) > c.interval {
					delete(c.cache, key)
				}

			}
			c.mu.Unlock()
		}
	}()
}
