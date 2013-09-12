package assert

import (
	"bytes"
	"fmt"
	"github.com/mgutz/ansi"
)

var (
	colorExpected = ansi.ColorFunc("green")
	colorActual   = ansi.ColorFunc("red")
)

type AssertError interface {
	Error() string
}

type valueError struct {
	expected interface{}
	actual   interface{}
}

// ----------------------------------------------------------------------------
// Equality Error
// ----------------------------------------------------------------------------

type EqualityError valueError

func EqualityErr(expected, actual interface{}) AssertError {
	return &EqualityError{expected, actual}
}

func (e *EqualityError) Expected() interface{} {
	return e.expected
}

func (e *EqualityError) Actual() interface{} {
	return e.actual
}

func (e *EqualityError) Error() string {
	buf := new(bytes.Buffer)
	buf.WriteString("\n")
	buf.WriteString("Expected: ")
	buf.WriteString(colorExpected(fmt.Sprintf("%v", e.expected)))
	buf.WriteString("\n")
	buf.WriteString("Actual:   ")
	buf.WriteString(colorActual(fmt.Sprintf("%v", e.actual)))

	return buf.String()
}

func expActString(expected, actual interface{}) {
}

// ----------------------------------------------------------------------------
// Length Error
// ----------------------------------------------------------------------------

type LengthError struct {
	expected int
	actual   int
	l        interface{}
}

func LengthErr(expected, actual int, l interface{}) AssertError {
	return &LengthError{expected, actual, l}
}

func (e *LengthError) Error() string {
	expected := colorExpected(fmt.Sprintf("%d", e.expected))
	actual := colorActual(fmt.Sprintf("%d", e.actual))
	buf := new(bytes.Buffer)

	buf.WriteString("\n")
	buf.WriteString(fmt.Sprintf("Expected length: %s\n", expected))
	buf.WriteString(fmt.Sprintf("Actual length:   %s\n", actual))
	buf.WriteString(fmt.Sprintf("Received:        %v\n", e.l))

	return buf.String()
}

// ----------------------------------------------------------------------------
// Runtime Error
// ----------------------------------------------------------------------------

type RuntimeError struct {
	msg string
}

func RuntimeErr(msg string) *RuntimeError {
	return &RuntimeError{msg}
}

func (e *RuntimeError) Expected() interface{} {
	return nil
}

func (e *RuntimeError) Actual() interface{} {
	return nil
}

func (e *RuntimeError) Error() string {
	buf := new(bytes.Buffer)
	buf.WriteString("\n")
	buf.WriteString("Runtime Error: ")
	buf.WriteString(colorActual(e.msg))

	return buf.String()
}
