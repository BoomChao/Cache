package fifo

import (
	"testing"

	"github.com/matryer/is"
)

func TestSetGet(t *testing.T) {
	is := is.New(t)

	cache := New(24, nil)
	cache.DelOldest()

	cache.Set("k1", 1)
	v := cache.Get("k1")
	is.Equal(v, 1)

	cache.Del("k1")
	is.Equal(0, cache.Len())
	// cache.Set("k2", time.Now())
}

func TestOnEvicated(t *testing.T) {
	is := is.New(t)

	keys := make([]string, 0, 8)
	onEvicated := func(key string, value interface{}) {
		keys = append(keys, key)
	}
	cache := New(16, onEvicated)

	cache.Set("k1", 1)
	cache.Set("k2", 2)
	cache.Get("k1")
	cache.Set("k3", 3)
	cache.Get("k1")
	cache.Set("k4", 4)

	expected := []string{"k1", "k2"}
	is.Equal(expected, keys)
	is.Equal(2, cache.Len())
}
