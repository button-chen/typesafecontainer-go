package typesafecontainer

import (
	"bytes"
	"testing"
)

func TestTPool(t *testing.T) {
	var pool = NewPool(func() *bytes.Buffer {
		return new(bytes.Buffer)
	})

	buff := pool.Get()
	pool.Put(buff)
}
