package typesafecontainer

import "container/ring"

type TRing[T any] struct {
	r *ring.Ring
}

func NewRing[T any](n int) *TRing[T] {
	return &TRing[T]{
		r: ring.New(n),
	}
}

func (tr *TRing[T]) Get() T {
	return tr.r.Value.(T)
}

func (tr *TRing[T]) Set(v T) {
	tr.r.Value = v
}

func (tr *TRing[T]) Do(f func(T)) {
	tr.r.Do(func(a any) {
		f(a.(T))
	})
}

func (tr *TRing[T]) Len() int {
	return tr.r.Len()
}

func (tr *TRing[T]) Link(s *TRing[T]) *TRing[T] {
	return &TRing[T]{r: tr.r.Link(s.r)}
}

func (tr *TRing[T]) Move(n int) *TRing[T] {
	return &TRing[T]{r: tr.r.Move(n)}
}

func (tr *TRing[T]) Next() *TRing[T] {
	return &TRing[T]{r: tr.r.Next()}
}

func (tr *TRing[T]) Prev() *TRing[T] {
	return &TRing[T]{r: tr.r.Prev()}
}

func (tr *TRing[T]) Unlink(n int) *TRing[T] {
	return &TRing[T]{r: tr.r.Unlink(n)}
}
