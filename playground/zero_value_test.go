package playground

import (
	"fmt"
	"testing"
)

func TestZeroValueOfAnAnyField(t *testing.T) {
	type s struct {
		a any
	}
	var v = s{}
	fmt.Println("the zero value of an any field in a struct is: ", v.a)
}
