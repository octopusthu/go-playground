package playground

import (
	"fmt"
	"runtime"
	"testing"
)

func TestPrintCurrentStack(t *testing.T) {
	stack := make([]byte, 4096)
	length := runtime.Stack(stack, true)
	fmt.Printf("stack: \n%s\n", stack[:length])
}
