package caching

import "slices"

// FifoCache is a Cache
// It doesn't implement ManualCache methods, but uses "First In First Out" cache invalidation
type FifoCache[T any] struct {
	maxSize int
	store   []fifoElement[T]
}

type fifoElement[T any] struct {
	key   string
	value T
}

func NewFifoCache[T any](maxSize int) *FifoCache[T] {
	return &FifoCache[T]{
		maxSize: maxSize,
		store:   make([]fifoElement[T], maxSize),
	}
}

func (f *FifoCache[T]) Get(key string) (value T, ok bool) {
	index := slices.IndexFunc(f.store, func(element fifoElement[T]) bool {
		return element.key == key
	})
	if index == -1 {
		var result T
		return result, false
	}
	return f.store[index].value, true
}

func (f *FifoCache[T]) Add(key string, value T) {
	if len(f.store) == f.maxSize {
		f.store = f.store[1:]
	}
	f.store = append(f.store, fifoElement[T]{
		key:   key,
		value: value,
	})
}
