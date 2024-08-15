package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	t.Run("Cache is not nil", func(t *testing.T) {
		interval := time.Millisecond * 10
		cache := NewCache(interval)
		if cache.cache == nil {
			t.Error("cache is nil")
		}
	})

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	interval := time.Millisecond * 10
	cache := NewCache(interval)

	for _, c := range cases {
		t.Run(fmt.Sprintf("testing key %s", c.inputKey), func(t *testing.T) {
			cache.Add(c.inputKey, c.inputVal)
			actual, ok := cache.Get(c.inputKey)
			if !ok {
				t.Errorf("expected to find key value %s in cache", c.inputKey)
			}

			if string(actual) != string(c.inputVal) {
				t.Errorf("got %s, but expected %s", string(actual), string(c.inputVal))
			}
		})
	}
}

func TestReap(t *testing.T) {
	t.Run("should reap key value", func(t *testing.T) {
		interval := time.Millisecond * 10
		cache := NewCache(interval)

		keyOne := "key1"
		cache.Add(keyOne, []byte("val1"))

		time.Sleep(interval + time.Millisecond)

		_, ok := cache.Get(keyOne)

		if ok {
			t.Errorf("key %s should have been reaped", keyOne)
		}
	})

	t.Run("should not reap key value", func(t *testing.T) {
		interval := time.Millisecond * 10
		cache := NewCache(interval)

		keyOne := "key1"
		cache.Add(keyOne, []byte("val1"))

		time.Sleep(interval / 2)

		_, ok := cache.Get(keyOne)

		if !ok {
			t.Errorf("key %s should have not been reaped", keyOne)
		}
	})

}
