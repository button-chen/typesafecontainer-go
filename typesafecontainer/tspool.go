package typesafecontainer

import (
	"sync"
)

type TPool[T any] struct {
	pool sync.Pool
}

func NewPool[T any](newfn func() T) *TPool[T] {
	return &TPool[T]{
		pool: sync.Pool{New: func() any { return newfn() }},
	}
}

func (tp *TPool[T]) Get() T {
	return tp.pool.Get().(T)
}

func (tp *TPool[T]) Put(x T) {
	tp.pool.Put(x)
}
