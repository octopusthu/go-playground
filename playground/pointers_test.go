package playground

import (
	"fmt"
	"testing"
)

func TestBoolPointer(t *testing.T) {
	var boolPointerToNil *bool
	if true == *boolPointerToNil {
		fmt.Println("ok")
	}
}

func TestSetPointedToValue(t *testing.T) {
	var pointerToString *string
	fmt.Println("initial pointer value: ", pointerToString)
	*pointerToString = "modified pointer value"
	fmt.Println("modified pointer value: ", pointerToString)
	fmt.Println("modified pointed to value: ", *pointerToString)
}
