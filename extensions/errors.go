// Package extensions implements readable error stacks.
package extensions

import (
	"fmt"
)

type TraceableError struct {
	parent *TraceableError
	err    error
	op     string
}

// ErrStack adds an operation and error to the top level
// of an error stack.
func ErrStack(op string, err error) *TraceableError {
	if err == nil {
		return &TraceableError{}
	}

	newErrorStack := TraceableError{
		op:  op,
		err: err,
	}

	// check if there is a previous TraceableError in stack.
	if derr, ok := err.(*TraceableError); ok { //nolint:errorlint
		derr.parent = &newErrorStack
	}

	return &newErrorStack
}

func (e TraceableError) Unwrap() error {
	return e.err
}

func (e TraceableError) Error() string {
	// check for empty error
	if (TraceableError{}) == e {
		return ""
	}

	return fmt.Sprintf("%s()%s %s", e.op, e.separator(), e.err.Error())
}

func (e TraceableError) separator() string {
	if e.isRoot() {
		return "->"
	}

	return ":"
}

func (e TraceableError) isRoot() bool {
	_, ok := e.err.(*TraceableError) //nolint:errorlint

	return ok
}
