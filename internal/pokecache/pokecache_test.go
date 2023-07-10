package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddCache(t *testing.T) {
	cache := NewCache()

	cases := []struct { // table for better reading the original code bellow
		inputKey   string
		inputValue []byte
	}{
		{
			inputKey:   "key1",
			inputValue: []byte("value1"),
		},
		{
			inputKey:   "key2",
			inputValue: []byte("value2"),
		},
	}
	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputValue)
		val, ok := cache.get(cas.inputKey)
		if !ok {
			t.Error("key do not exist")
		}
		if string(val) != string(cas.inputValue) {
			t.Error("value doesn't match !")
		}
	}

	//cache.Add("key1", []byte("value1"))
	////fmt.Println(cache)
	//val, ok := cache.get("key1")
	//if !ok {
	//	t.Error("key do not exist")
	//}
	//if string(val) != "value1" {
	//	t.Error("value doesn't match !")
	//}

}

func TestReap(t *testing.T) {
	cache := NewCache()
	cache.Add("key1", []byte("val1"))
	cache.Add("key2", []byte("val2"))
	fmt.Println(cache)
	cache.Reap()
	time.Sleep(time.Millisecond * 5000)
	fmt.Println(cache)
}
