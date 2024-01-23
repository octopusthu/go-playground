package playground

import (
	"encoding/json"
	"fmt"
	"testing"
)

type testingStruct struct {
	foo string
	bar int
}

func TestSerialization(t *testing.T) {
	var nilPointer *testingStruct
	structVar := testingStruct{foo: "fooVal", bar: 100}
	pointerToStruct := &structVar
	var marshalled []byte
	marshalled, _ = json.Marshal(nilPointer)
	fmt.Printf("json.Marshal(nilPointer) = %v\n", marshalled)
	marshalled, _ = json.Marshal(structVar)
	fmt.Printf("json.Marshal(structVar) = %v\n", marshalled)
	marshalled, _ = json.Marshal(pointerToStruct)
	fmt.Printf("json.Marshal(pointerToStruct) = %v\n", marshalled)
}
