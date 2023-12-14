package playground

import (
    "fmt"
    "testing"
)

type multiInheritanceParent interface {
    multiInheritanceParentMethod()
}
type multiInheritanceAbstractImpl struct{}

func (multiInheritanceAbstractImpl) multiInheritanceParentMethod() {
    fmt.Println("parentImpl.parentMethod() invoked!")
}

type multiInheritanceConcreteImpl struct {
    multiInheritanceAbstractImpl
}

func TestTypeAssertionOfMultiInheritance(t *testing.T) {
    var o multiInheritanceParent = multiInheritanceConcreteImpl{
        multiInheritanceAbstractImpl: multiInheritanceAbstractImpl{}}
    var ok bool
    _, ok = o.(multiInheritanceAbstractImpl)
    fmt.Printf("o has the concrete type of multiInheritanceAbstractImpl: %v\n", ok)
    _, ok = o.(multiInheritanceConcreteImpl)
    fmt.Printf("o has the concrete type of multiInheritanceConcreteImpl: %v\n", ok)
}
