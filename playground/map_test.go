package playground

import (
	"fmt"
	"testing"
)

func TestGetNonexistentKey(t *testing.T) {
	m := map[int8]int8{1: 1, 2: 2}
	fmt.Println("get nonexistent key of map[int8]int8: ", m[3])
	n := map[int8]any{1: 1, 2: 2}
	fmt.Println("get nonexistent key of map[int8]any: ", n[3])
}
