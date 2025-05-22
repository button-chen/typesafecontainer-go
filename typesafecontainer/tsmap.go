package typesafecontainer

import "sync"

type TMap[K comparable, V any] struct {
	m sync.Map
}

func (tm *TMap[K, V]) Clear() {
	tm.m.Clear()
}

func (tm *TMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return tm.m.CompareAndDelete(key, old)
}

func (tm *TMap[K, V]) CompareAndSwap(key K, old V, new V) (swapped bool) {
	return tm.m.CompareAndSwap(key, old, new)
}

func (tm *TMap[K, V]) Delete(key K) {
	tm.m.Delete(key)
}

func (tm *TMap[K, V]) Load(key K) (V, bool) {
	v, ok := tm.m.Load(key)
	return tm.cast(v), ok
}

func (tm *TMap[K, V]) LoadAndDelete(key K) (V, bool) {
	v, ok := tm.m.LoadAndDelete(key)
	return tm.cast(v), ok
}

func (tm *TMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	v, ok := tm.m.LoadOrStore(key, value)
	return tm.cast(v), ok
}

func (tm *TMap[K, V]) Store(key K, value V) {
	tm.m.Store(key, value)
}

func (tm *TMap[K, V]) Swap(key K, value V) (V, bool) {
	v, ok := tm.m.Swap(key, value)
	return tm.cast(v), ok
}

func (tm *TMap[K, V]) Range(f func(key K, value V) bool) {
	tm.m.Range(func(k, v any) bool {
		return f(k.(K), v.(V))
	})
}

func (tm *TMap[K, V]) cast(val any) V {
	v, ok := val.(V)
	if ok {
		return v
	}
	var zv V
	return zv
}

type StringStringMap = TMap[string, string]
type StringTMap[T any] = TMap[string, T]
