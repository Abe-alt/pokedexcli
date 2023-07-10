package pokecache

import (
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func add(key string, value []byte) {
	cache := make(map[string]cacheEntry)

	cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

}
