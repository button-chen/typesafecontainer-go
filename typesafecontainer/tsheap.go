package typesafecontainer

import (
	"container/heap"
	"sort"
)

type THeap[T any] struct {
	h *WrapInterface[T]
}

func NewHeap[T any](h Interface[T]) *THeap[T] {
	th := &THeap[T]{
		h: &WrapInterface[T]{h},
	}
	heap.Init(th.h)
	return th
}

func (th *THeap[T]) Fix(i int) {
	heap.Fix(th.h, i)
}

func (th *THeap[T]) Pop() T {
	return heap.Pop(th.h).(T)
}

func (th *THeap[T]) Push(x T) {
	heap.Push(th.h, x)
}

func (th *THeap[T]) Remove(i int) T {
	return heap.Remove(th.h, i).(T)
}

type Interface[T any] interface {
	sort.Interface
	Push(x T)
	Pop() T
}

type WrapInterface[T any] struct {
	h Interface[T]
}

func (w WrapInterface[T]) Len() int {
	return w.h.Len()
}

func (w WrapInterface[T]) Less(i, j int) bool {
	return w.h.Less(i, j)
}

func (w WrapInterface[T]) Swap(i, j int) {
	w.h.Swap(i, j)
}

func (w *WrapInterface[T]) Push(x any) {
	w.h.Push(x.(T))
}

func (w *WrapInterface[T]) Pop() any {
	return w.h.Pop()
}
