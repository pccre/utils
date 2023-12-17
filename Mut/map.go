package Mut

import "sync"

type Map[K comparable, V any] struct {
	Map map[K]V
	Mut *sync.RWMutex
}

func (m *Map[K, V]) Set(key K, value V) {
	m.Mut.Lock()
	m.Map[key] = value
	m.Mut.Unlock()
}

func (m *Map[K, V]) Get(key K) V {
	m.Mut.RLock()
	defer m.Mut.RUnlock()
	return m.Map[key]
}
