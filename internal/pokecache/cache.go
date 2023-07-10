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
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	return c
}

func (c *Cache) Add(key string, value []byte) {
	//cache := make(map[string]cacheEntry)

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}

}

func (c *Cache) get(key string) ([]byte, bool) {

	val, ok := c.cache[key]
	return val.val, ok
}

func (c *Cache) Reap() {
	t := time.Now().UTC().Add(time.Millisecond * 2000)
	for k, v := range c.cache {
		if v.createdAt.Before(t) {
			delete(c.cache, k)
		}
	}
}

func (c *Cache) reaploop() {
	ticker := time.NewTicker(time.Minute * 1)
	for range ticker.C {
		c.Reap()
	}
}

//func (c *Cache) reaploop(now time.Time, last time.Duration) {
//
//	for k, v := range c.cache {
//		if v.createdAt.Before(now.Add(-last)) {
//			delete(c.cache, k)
//		}
//	}
//
//}
