package playground

import (
	"fmt"
	"testing"
)

type embedded interface {
	embeddedMethod()
}
type embeddedImpl struct{}

func (embeddedImpl) embeddedMethod() {
	fmt.Println("embeddedImpl.embeddedMethod() invoked!")
}

type outer interface {
	outerMethod()
}
type outerImpl struct {
	embedded
}

func (outerImpl) outerMethod() {}

func TestTypeAssertionOfEmbeddedField(t *testing.T) {
	var o outer = outerImpl{embedded: embeddedImpl{}}
	if e, ok := o.(embedded); ok {
		e.embeddedMethod()
		fmt.Println("o has the concrete type of e!")
	} else {
		t.Error("failed to assert that o has the concrete type of e!")
	}
}
