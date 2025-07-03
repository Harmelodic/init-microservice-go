package caching

import "slices"

// LfuCache is a Cache
// It doesn't implement ManualCache methods, but uses "Least Frequently Used" cache invalidation
type LfuCache[T any] struct {
	maxSize int
	store   []lfuElement[T]
}

type lfuElement[T any] struct {
	timesAccessed int
	key           string
	value         T
}

func NewLfuCache[T any](maxSize int) *LfuCache[T] {
	return &LfuCache[T]{
		maxSize: maxSize,
		store:   make([]lfuElement[T], maxSize),
	}
}

func (l *LfuCache[T]) Get(key string) (value T, ok bool) {
	index := slices.IndexFunc(l.store, func(element lfuElement[T]) bool {
		return element.key == key
	})
	if index == -1 {
		var empty T
		return empty, false
	}
	element := l.store[index]
	element.timesAccessed++
	return element.value, true
}

func (l *LfuCache[T]) Add(key string, value T) {
	if len(l.store) == l.maxSize {
		slices.SortFunc(l.store, func(a, b lfuElement[T]) int {
			return a.timesAccessed - b.timesAccessed
		})
		l.store = slices.Delete(l.store, 0, 1)
	}
	l.store = append(l.store, lfuElement[T]{
		timesAccessed: 0,
		key:           key,
		value:         value,
	})
}
