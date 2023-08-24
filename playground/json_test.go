package playground

import (
	"encoding/json"
	"fmt"
	"testing"
)

type strct struct {
	MapField map[string]any `json:"mapField"`
}

func TestUnmarshalIntField(t *testing.T) {
	v := strct{MapField: map[string]any{"intValField": 10}}
	bytes, _ := json.Marshal(v)
	fmt.Println("marshalled: ", string(bytes))
	var unmarshalled strct
	_ = json.Unmarshal(bytes, &unmarshalled)
	fmt.Println("unmarshalled: ", unmarshalled)
	var float64Assertion = unmarshalled.MapField["intValField"].(float64)
	fmt.Println("float64Assertion: ", float64Assertion)
	fmt.Println("converted int: ", int(float64Assertion))
	fmt.Println("converted uint32: ", uint32(float64Assertion))
}
