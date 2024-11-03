package playground

import (
    "fmt"
    "testing"
)

type TestingStruct struct {
}

func (t *TestingStruct) pointerReceiverMethod() {
    fmt.Printf("pointerReceiverMethod invoked on TestingStruct %v\n", t)
}

func TestCallingPointerReceiverMethodForNilInstance(t *testing.T) {
    var testingStruct *TestingStruct
    testingStruct.pointerReceiverMethod()
}
