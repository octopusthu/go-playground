package playground

import (
	"fmt"
	"strconv"
	"testing"
)

func TestParseUint(t *testing.T) {
	fmt.Println(strconv.ParseUint("20230101", 10, 32))
}
