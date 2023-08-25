package playground

import (
	"fmt"
	"testing"
)

func TestLenOfNilSlice(t *testing.T) {
	var nilSlice []int8 = nil
	fmt.Println("len of a nil slice is: ", len(nilSlice))
}
