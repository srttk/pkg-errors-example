package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	err, ok := errors.Cause(errFunc()).(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := err.StackTrace()
	fmt.Printf("%+v", st[0:2])
}

func errFunc() error {
  return errors.New("some err")
}
