package caching

type CacheManager[T any] struct {
	Cache Cache[T]
}

func (cacheManager *CacheManager[T]) GetOrAddFromFunc(key string, fetch func(key string) T) T {
	value, ok := cacheManager.Cache.Get(key)
	if ok {
		return value
	}
	element := fetch(key)
	cacheManager.Cache.Add(key, element)
	return element
}

func NewCacheManager[T any](cache Cache[T]) *CacheManager[T] {
	return &CacheManager[T]{
		Cache: cache,
	}
}
