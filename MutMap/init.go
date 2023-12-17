package MutMap

import "sync"

type MutMap[T any, K comparable] struct {
	Map map[K]T
	Mut *sync.RWMutex
}

func (m *MutMap[T, K]) Set(key K, value T) {
	m.Mut.Lock()
	m.Map[key] = value
	m.Mut.Unlock()
}

func (m *MutMap[T, K]) Get(key K) T {
	m.Mut.RLock()
	defer m.Mut.RUnlock()
	return m.Map[key]
}
