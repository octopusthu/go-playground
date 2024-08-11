package playground

import (
	"fmt"
	"testing"
	"time"
)

type MockStruct struct {
	IntField    int8
	StringField string
	TimeField   time.Time
}

var (
	mockStruct1 = MockStruct{1, "1", time.Now()}
	mockStruct2 = MockStruct{2, "2", time.Now()}
	mockStruct3 = MockStruct{3, "3", time.Now()}
)

func TestStructFieldValues(t *testing.T) {
	mockStruct := MockStruct{}
	fmt.Printf("mockStruct: %+v\n", mockStruct)
}

func TestNilSlice(t *testing.T) {
	var nilSlice []int8
	fmt.Printf("len(nilSlice) = %d\n", len(nilSlice))
	for _, _ = range nilSlice {
		fmt.Println("OMG!")
	}
}
