package playground

import (
	"fmt"
	"testing"
)

func TestNilAndEmptySlices(t *testing.T) {
	var nilSlice []int8 = nil
	var emptySlice = make([]int8, 0)
	fmt.Printf("nil slice is nil: %v\n", nilSlice == nil)     // true
	fmt.Printf("empty slice is nil: %v\n", emptySlice == nil) // false
	fmt.Printf("len of nil slice is %d\n", len(nilSlice))     // 0
	fmt.Printf("len of empty slice is %d\n", len(emptySlice)) // 0
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
