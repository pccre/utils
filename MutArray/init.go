package MutArray

import (
	"sync"

	"github.com/pccre/utils/shared"
)

type MutArray[T any] struct {
	Array []T
	Mut   *sync.RWMutex
}

func (a *MutArray[T]) Set(i int, v T) {
	a.Mut.Lock()
	a.Array[i] = v
	a.Mut.Unlock()
}

func (a *MutArray[T]) Append(v ...T) {
	a.Mut.Lock()
	a.Array = append(a.Array, v...)
	a.Mut.Unlock()
}

func (a *MutArray[T]) Get(i int) T {
	a.Mut.RLock()
	defer a.Mut.RUnlock()
	return a.Array[i]
}

func (a *MutArray[T]) Remove(i int) {
	a.Mut.Lock()
	a.Array = shared.Remove(a.Array, i)
	a.Mut.Unlock()
}
