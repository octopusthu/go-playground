package playground

import (
	"fmt"
	"testing"
)

func TestPanicInGoRoutine(t *testing.T) {
	go func() {
		doPanic()
	}()
	fmt.Println("Normal execution...")
}

func TestPanic(t *testing.T) {
	doPanic()
}

func TestPanicAndRecover(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("Recovered from panic: ", v)
		}
	}()
	doPanic()
}

func doPanic() {
	panic("This is a panic")
}
