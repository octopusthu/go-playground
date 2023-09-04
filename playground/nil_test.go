package playground

import (
	"fmt"
	"testing"
)

func TestForNil(t *testing.T) {
	var nilArr []int8 = nil
	for i, elem := range nilArr {
		fmt.Println(i, elem)
	}
}

func TestPrintNil(t *testing.T) {
	var pointer *int = nil
	fmt.Printf("nil pointer: %+v\n", pointer)
}
