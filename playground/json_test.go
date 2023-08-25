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
