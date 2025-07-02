package caching

type Cache[T any] interface {
	Get(key string) (value T, ok bool)
	Add(key string, value T)
}

type ManualCache[T any] interface {
	Cache[T]
	Invalidate(key string)
	InvalidateAll()
}

// BasicCache is a cache that implements Cache / ManualCache
type BasicCache[T any] struct {
	store map[string]T
}

func NewBasicCache[T any]() *BasicCache[T] {
	return &BasicCache[T]{
		store: make(map[string]T),
	}
}

func (cache *BasicCache[T]) Get(key string) (value T, ok bool) {
	t, ok := cache.store[key]
	return t, ok
}

func (cache *BasicCache[T]) Add(key string, value T) {
	cache.store[key] = value
}

func (cache *BasicCache[T]) Invalidate(key string) {
	delete(cache.store, key)
}

func (cache *BasicCache[T]) InvalidateAll() {
	clear(cache.store)
}
