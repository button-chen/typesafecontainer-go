package typesafecontainer

import (
	"fmt"
	"testing"
)

func TestTRing(t *testing.T) {

	r := NewRing[string](5)
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Set(fmt.Sprintf("index: %d", i))
		r = r.Next()
	}

	r.Do(func(p string) {
		fmt.Println(p)
	})
}
