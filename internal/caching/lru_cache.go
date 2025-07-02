package caching

import (
	"slices"
	"time"
)

type LruCache[T any] struct {
	maxSize int
	store   []lruElement[T]
}

type lruElement[T any] struct {
	lastUsed time.Time
	key      string
	value    T
}

func (l *LruCache[T]) Get(key string) (value T, ok bool) {
	index := slices.IndexFunc(l.store, func(element lruElement[T]) bool {
		return element.key == key
	})
	if index == -1 {
		var empty T
		return empty, false
	}
	element := l.store[index]
	element.lastUsed = time.Now()
	return element.value, true
}

func (l *LruCache[T]) Add(key string, value T) {
	if len(l.store) == l.maxSize {
		// Order by least used -> most used
		slices.SortFunc(l.store, func(a, b lruElement[T]) int {
			if a.lastUsed.Before(b.lastUsed) {
				return -1
			}
			if a.lastUsed.Equal(b.lastUsed) {
				return 0
			}
			// Else must be After!
			return 1
		})
		l.store = slices.Delete(l.store, 0, 1) // Delete first element (least used)
	}
	l.store = append(l.store, lruElement[T]{
		lastUsed: time.Now(),
		key:      key,
		value:    value,
	})
}
