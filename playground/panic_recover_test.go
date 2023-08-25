package playground

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPanicInGoRoutine(t *testing.T) {
	go func() {
		doPanic()
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Normal execution...")
}

func TestPanicInGoRoutineWithWaitGroup(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		doPanic()
	}()
	wg.Wait()
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

func TestReturnValueOfPanickedFunc(t *testing.T) {
	err := panicFuncWithReturnValue()
	fmt.Println("return value of panicked function: ", err)
}

func doPanic() {
	panic("This is a panic")
}

func panicFuncWithReturnValue() error {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered from panic: ", v)
		}
	}()
	panic("This is a panic")
	return nil
}
