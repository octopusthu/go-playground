package playground

import (
    "fmt"
    "testing"
)

type stringerTest struct {
    foo string
    bar int
}

func (v stringerTest) String() string {
    return fmt.Sprintf("stringerTest(foo: %s, bar: %d)", v.foo, v.bar)
}

func TestStringer(t *testing.T) {
    v := stringerTest{"fooValue", 1}
    fmt.Println("printed directly: ", v)
    fmt.Printf("printed by percent v: %v\n", v)
    fmt.Printf("printed by percent +v: %+v\n", v)
}
