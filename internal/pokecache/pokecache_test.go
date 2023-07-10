package pokecache

import (
	"testing"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddCache(t *testing.T) {
	cache := NewCache()
	cache.Add("key1", []byte("value1"))
	//fmt.Println(cache)
	val, ok := cache.get("key1")
	if !ok {
		t.Error("key do not exist")
	}
	if string(val) != "value1" {
		t.Error("value doesn't match !")
	}

}
