package Mut

import "sync"

type Map[T any, K comparable] struct {
	Map map[K]T
	Mut *sync.RWMutex
}

func (m *Map[T, K]) Set(key K, value T) {
	m.Mut.Lock()
	m.Map[key] = value
	m.Mut.Unlock()
}

func (m *Map[T, K]) Get(key K) T {
	m.Mut.RLock()
	defer m.Mut.RUnlock()
	return m.Map[key]
}
