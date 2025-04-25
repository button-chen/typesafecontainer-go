package typesafecontainer

import (
	"fmt"
	"testing"
)

func TestTMap(t *testing.T) {
	ssm := TMap[string, string]{}

	ssm.Store("foo", "bar")
	fmt.Println(ssm.Load("foo"))
}
