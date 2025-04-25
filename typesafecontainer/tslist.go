package typesafecontainer

import "container/list"

type TList[T any] struct {
	l *list.List
}

func NewList[T any]() *TList[T] {
	return &TList[T]{
		l: list.New(),
	}
}

func (tl *TList[T]) Back() *TElement[T] {
	return &TElement[T]{e: tl.l.Back()}
}

func (tl *TList[T]) Front() *TElement[T] {
	return &TElement[T]{e: tl.l.Front()}
}

func (tl *TList[T]) Init() *TList[T] {
	return &TList[T]{tl.l.Init()}
}

func (tl *TList[T]) InsertAfter(v T, mark *TElement[T]) *TElement[T] {
	return &TElement[T]{e: tl.l.InsertAfter(v, mark.e)}
}

func (tl *TList[T]) InsertBefore(v T, mark *TElement[T]) *TElement[T] {
	return &TElement[T]{e: tl.l.InsertBefore(v, mark.e)}
}

func (tl *TList[T]) Len() int {
	return tl.l.Len()
}

func (tl *TList[T]) MoveAfter(e, mark *TElement[T]) {
	tl.l.MoveAfter(e.e, mark.e)
}

func (tl *TList[T]) MoveBefore(e, mark *TElement[T]) {
	tl.l.MoveBefore(e.e, mark.e)
}

func (tl *TList[T]) MoveToBack(e *TElement[T]) {
	tl.l.MoveToBack(e.e)
}

func (tl *TList[T]) MoveToFront(e *TElement[T]) {
	tl.l.MoveToFront(e.e)
}

func (tl *TList[T]) PushBack(v T) *TElement[T] {
	return &TElement[T]{e: tl.l.PushBack(v)}
}

func (tl *TList[T]) PushBackList(other *TList[T]) {
	tl.l.PushBackList(other.l)
}

func (tl *TList[T]) PushFront(v T) *TElement[T] {
	return &TElement[T]{e: tl.l.PushFront(v)}
}

func (tl *TList[T]) PushFrontList(other *TList[T]) {
	tl.l.PushFrontList(other.l)
}

func (tl *TList[T]) Remove(e *TElement[T]) T {
	return tl.l.Remove(e.e).(T)
}

type TElement[T any] struct {
	e *list.Element
}

func (te *TElement[T]) Next() *TElement[T] {
	if t := te.e.Next(); t != nil {
		return &TElement[T]{e: t}
	}
	return nil
}

func (te *TElement[T]) Prev() *TElement[T] {
	if t := te.e.Prev(); t != nil {
		return &TElement[T]{e: t}
	}
	return nil
}

func (te *TElement[T]) Get() T {
	return te.e.Value.(T)
}
