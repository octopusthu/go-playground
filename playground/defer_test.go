package playground

import (
	"fmt"
	"testing"
)

func TestSetStructFieldAfterDefer(t *testing.T) {
	type s struct {
		a string
	}
	str := s{a: "initial value"}
	defer func() {
		fmt.Println("value evaluated by defer is: ", str.a)
	}()
	str.a = "modified value"
}

func TestPanicAfterDefer(t *testing.T) {
	defer func() {
		fmt.Println("defer statement executed")
	}()
	fmt.Println("before panic")
	panic("panicked")
}
