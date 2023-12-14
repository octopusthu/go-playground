package playground

import (
	"fmt"
	"testing"
)

type parent interface {
	parentMethod()
}
type parentImpl struct{}

func (parentImpl) parentMethod() {
	fmt.Println("parentImpl.parentMethod() invoked!")
}

type child interface {
	childMethod()
}
type childImplEmbeddingParentInterface struct {
	parent
}
type childImplEmbeddingParentImplStruct struct {
	parentImpl
}

func (childImplEmbeddingParentInterface) childMethod()  {}
func (childImplEmbeddingParentImplStruct) childMethod() {}

func TestTypeAssertionOfEmbeddedInterface(t *testing.T) {
	var o parent = childImplEmbeddingParentInterface{parent: parentImpl{}}
	if e, ok := o.(childImplEmbeddingParentInterface); ok {
		e.parentMethod()
		fmt.Println("o has the concrete type of e!")
	} else {
		t.Error("failed to assert that o has the concrete type of e!")
	}
}
func TestTypeAssertionOfEmbeddedStruct(t *testing.T) {
	var o parent = childImplEmbeddingParentImplStruct{parentImpl: parentImpl{}}
	if e, ok := o.(childImplEmbeddingParentImplStruct); ok {
		e.parentMethod()
		fmt.Println("o has the concrete type of e!")
	} else {
		t.Error("failed to assert that o has the concrete type of e!")
	}
}
