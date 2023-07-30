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

func TestStructFieldValues(t *testing.T) {
	mockStruct := MockStruct{}
	fmt.Printf("mockStruct: %+v\n", mockStruct)
}
