package assert

import (
	"fmt"
	"reflect"
	"testing"
)

type test struct {
	t     *testing.T
	value interface{}
}

type Assertion interface {
	Check(v interface{}) AssertError
}

func That(t *testing.T, value interface{}) *test {
	return &test{
		t:     t,
		value: value,
	}
}

// TODO: Recover?
func (t *test) Is(a Assertion) *test {
	if err := a.Check(t.value); err != nil {
		t.t.Fatal(err)
	}
	return t
}

// ----------------------------------------------------------------------------
// Nil Assertion
// ----------------------------------------------------------------------------

func (t *test) IsNil() *test {
	return t.Is(Nil())
}

type IsNil int

func Nil() *IsNil {
	return new(IsNil)
}

func (n *IsNil) Check(actual interface{}) AssertError {
	switch v := reflect.ValueOf(actual); v.Kind() {
	default:
		if nil != actual {
			return EqualityErr(nil, actual)
		}
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr,
		reflect.Slice:
		if !v.IsNil() {
			return EqualityErr(nil, v)
		}
	}
	return nil
}

// ----------------------------------------------------------------------------
// Equals Assertion
// ----------------------------------------------------------------------------

func (t *test) Equals(v interface{}) *test {
	return t.Is(&IsEqual{v})
}

type IsEqual struct {
	v interface{}
}

func (e *IsEqual) Check(actual interface{}) AssertError {
	if e.v != actual {
		return EqualityErr(e.v, actual)
	}
	return nil
}

// ----------------------------------------------------------------------------
// Boolean Assertions
// ----------------------------------------------------------------------------

func (t *test) IsTrue() *test {
	return t.Equals(true)
}

func (t *test) IsFalse() *test {
	return t.Equals(false)
}

// ----------------------------------------------------------------------------
// HasLen Assertion
// ----------------------------------------------------------------------------

func (t *test) Haslen(length int) *test {
	return t.Is(&Lengthed{length})
}

type Lengthed struct {
	length int
}

func (l *Lengthed) Check(val interface{}) AssertError {
	switch v := reflect.ValueOf(val); v.Kind() {
	default:
		return RuntimeErr(fmt.Sprintf("Type %v has no length", v))
	case reflect.Map, reflect.Array, reflect.Slice, reflect.Chan,
		reflect.String:
		if l.length != v.Len() {
			return LengthErr(l.length, v.Len(), val)
		}
	}
	return nil
}

// ----------------------------------------------------------------------------
// Contains Assertion
// ----------------------------------------------------------------------------

func (t *test) Contains(i interface{}) *test {
	return t.Is(&Containing{i})
}

type Containing struct {
	v interface{}
}

func (c *Containing) Check(val interface{}) AssertError {
	return nil
}

func (c *Containing) containsString(exp, act string) bool {
	return false
}
