package assert

import (
	"errors"
	"reflect"
	"testing"
)

func TestErrNotEqual(t *testing.T) {
	msg := errNotEqual("three", 3).Error()
	exp := "" +
		"expected string: \"three\"\n" +
		"   received int: 3"

	if msg != exp {
		t.Fatalf("Expected: '%s', got '%s'", exp, msg)
	}
}

func TestErrEqual(t *testing.T) {
	msg := errEqual("three", 3).Error()
	exp := "" +
		"expected not string: \"three\"\n" +
		"       received int: 3"

	if msg != exp {
		t.Fatalf("Expected: '%s', got '%s'", exp, msg)
	}
}

func TestStringValueString(t *testing.T) {
	e := NewError()
	That(t, e.valueStr("foo")).Equals(`"foo"`)
}

func TestIntValueString(t *testing.T) {
	e := NewError()
	That(t, e.valueStr(3)).Equals("3")
}

func TestErrorValueString(t *testing.T) {
	err := errors.New("that's strange")
	e := NewError()
	That(t, e.valueStr(err)).Equals("that's strange")
}

func TestValueValueString(t *testing.T) {
	val := reflect.ValueOf(3)
	e := NewError()
	That(t, e.valueStr(val)).Equals("3")
}

func TestErrorMessage(t *testing.T) {
	orig := errors.New("error message")
	err := Assert(orig).IsNil().Check()
	That(t, err).NotNil()

	e := NewError()
	That(t, e.valueStr(orig)).Equals("error message")

	msg := err.Error()
	exp := "" +
		"              expected: <nil>\n" +
		"received reflect.Value: error message"

	if msg != exp {
		t.Fatalf("Expected: '%s', got '%s'", exp, msg)
	}
}
