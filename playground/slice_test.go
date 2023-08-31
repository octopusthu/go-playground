package playground

import (
	"fmt"
	"testing"
)

func TestLenOfNilSlice(t *testing.T) {
	var nilSlice []int8 = nil
	fmt.Printf("len of a nil slice is %d\n", len(nilSlice))
}

func TestForRangeNilSlice(t *testing.T) {
	var nilSlice []int8
	for _, element := range nilSlice {
		fmt.Println(element)
	}
	fmt.Println("OK.")
}

func TestAppendToNilSlice(t *testing.T) {
	var nilSlice []int8
	fmt.Printf("nilSlice before appending: %v, is nil? %v\n", nilSlice, nilSlice == nil)
	nilSlice = append(nilSlice, nilSlice...)
	fmt.Printf("nilSlice after appending: %v, is nil? %v\n", nilSlice, nilSlice == nil)
}
