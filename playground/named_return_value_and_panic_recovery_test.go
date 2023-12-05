package playground

import (
	"errors"
	"log"
	"testing"
)

const (
	writtenByRecover = "written by recover"
)

func TestNamedReturnValueAndPanicRecovery(t *testing.T) {
	err := testNamedReturnValueAndPanicRecovery()
	if err.Error() != writtenByRecover {
		t.Error("unexpected err: ", err)
	}
}

func testNamedReturnValueAndPanicRecovery() (err error) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("recovered from panic: ", v)
			err = errors.New(writtenByRecover)
		}
	}()
	err = errors.New("written by normal execution")
	panic("panic!")
	return err // not necessarily
}
