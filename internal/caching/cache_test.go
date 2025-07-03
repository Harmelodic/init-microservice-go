package caching

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicCache(t *testing.T) {
	cache := NewBasicCache[string]()

	testKey, testValue := "something", "some_value"
	testKey2, testValue2 := "something2", "some_value2"

	value, ok := cache.Get(testKey)
	assert.False(t, ok)
	assert.NotEqual(t, testValue, value)

	cache.Add(testKey, testValue)

	value, ok = cache.Get(testKey)
	assert.True(t, ok)
	assert.Equal(t, testValue, value)

	cache.Invalidate(testKey)
	value, ok = cache.Get(testKey)
	assert.False(t, ok)
	assert.NotEqual(t, testValue, value)

	cache.Add(testKey, testValue)
	cache.Add(testKey2, testValue2)

	cache.InvalidateAll()
	_, ok = cache.Get(testKey)
	assert.False(t, ok)
	_, ok = cache.Get(testKey2)
	assert.False(t, ok)
	assert.Empty(t, cache.store)
}
