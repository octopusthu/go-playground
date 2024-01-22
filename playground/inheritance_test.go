package playground

import (
    "fmt"
    "testing"
)

type Event interface {
    GetCode() string
}

type AbstractEvent struct {
    code string
}

func (v AbstractEvent) GetCode() string {
    return v.code
}

type AbstractStreamEvent struct {
    AbstractEvent
    stream string
}

func (v AbstractStreamEvent) GetStream() string {
    return v.stream
}

type ConcreteEvent struct {
    AbstractStreamEvent
    concreteData string
}

// TestInheritance
//
// Conclusion: a variable has the type of its implementation type and its interface (in this case ConcreteEvent and Event),
// but not the types embedded by its implementation type  (in this case AbstractEvent and AbstractStreamEvent).
func TestInheritance(t *testing.T) {
    var concreteEvent Event = ConcreteEvent{
        AbstractStreamEvent: AbstractStreamEvent{
            AbstractEvent: AbstractEvent{
                code: "code",
            },
            stream: "stream",
        },
        concreteData: "concreteData",
    }
    var ok bool
    _, ok = concreteEvent.(ConcreteEvent)
    fmt.Printf("concreteEvent has the type of %s: %v\n", "ConcreteEvent", ok)
    _, ok = concreteEvent.(AbstractStreamEvent)
    fmt.Printf("concreteEvent has the type of %s: %v\n", "AbstractStreamEvent", ok)
    _, ok = concreteEvent.(AbstractEvent)
    fmt.Printf("concreteEvent has the type of %s: %v\n", "AbstractEvent", ok)
    _, ok = concreteEvent.(Event)
    fmt.Printf("concreteEvent has the type of %s: %v\n", "Event", ok)
}
