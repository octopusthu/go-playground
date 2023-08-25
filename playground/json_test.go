package playground

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestUnmarshalStringPointerField(t *testing.T) {
	type testingStruct struct {
		stringPointerField *string
	}

	normalStringStruct := "{\"stringPointerField\": \"normal value\"}"
	nilStringStruct := "{\"stringPointerField\": null}"
	zeroStringStruct := "{\"stringPointerField\": \"\"}"

	var unmarshalledNormalStringStruct testingStruct
	var unmarshalledNilStringStruct testingStruct
	var unmarshalledZeroStringStruct testingStruct

	_ = json.Unmarshal([]byte(normalStringStruct), &unmarshalledNormalStringStruct)
	_ = json.Unmarshal([]byte(nilStringStruct), &unmarshalledNilStringStruct)
	_ = json.Unmarshal([]byte(zeroStringStruct), &unmarshalledZeroStringStruct)
	fmt.Printf("unmarshalledNormalStringStruct: %v\n", unmarshalledNormalStringStruct)
	fmt.Printf("unmarshalledNilStringStruct: %v\n", unmarshalledNilStringStruct)
	fmt.Printf("unmarshalledZeroStringStruct: %v\n", unmarshalledZeroStringStruct)
}

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
