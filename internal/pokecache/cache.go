package pokecache

import (
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache() Cache {
	interval := time.Minute * 5
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.Reaploop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	//cache := make(map[string]cacheEntry)

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {

	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) Reap(interval time.Duration) {
	t := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(t) {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) Reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Reap(interval)
	}
}
