package typesafecontainer

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {

	l := NewList[string]()
	e4 := l.PushBack("foo")
	e1 := l.PushFront("bar")
	l.InsertBefore("amy", e4)
	l.InsertAfter("ruby", e1)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Get())
	}
}
