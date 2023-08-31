package playground

import (
	"fmt"
	"testing"
)

func TestLenOfNilSlice(t *testing.T) {
	var nilSlice []int8 = nil
	fmt.Println("len of a nil slice is: ", len(nilSlice))
}

func TestForRangeNilSlice(t *testing.T) {
	var nilSlice []int8
	for _, element := range nilSlice {
		fmt.Println(element)
	}
	fmt.Println("OK.")
}
