package playground

import (
	"fmt"
	"strconv"
	"testing"
)

func TestParseUint(t *testing.T) {
	fmt.Println(strconv.ParseUint("20230101", 10, 32))
}

func TestUintAsResultOfNegativeDifference(t *testing.T) {
	var uintResult uint64
	var minuend, subtrahend uint64
	minuend = 1
	subtrahend = 2
	uintResult = minuend - subtrahend
	fmt.Printf("%d - %d = (uint64)%d\n", minuend, subtrahend, uintResult)
}

func TestParseFloatStringAsInt(t *testing.T) {
	parsed, _ := strconv.ParseInt("0.1", 10, 64)
	fmt.Printf("0.1: %d", parsed)
}

func TestConvertIntToFloat(t *testing.T) {
	var a = 5
	var b = (float64)(a)
	fmt.Printf("b: %f\n", b)
}
