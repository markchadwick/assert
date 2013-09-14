package assert

import (
	"fmt"
	"reflect"
	"strings"
)

type Check interface {
	Check(interface{}) error
}

// ----------------------------------------------------------------------------
// Nil Check
// ----------------------------------------------------------------------------

type IsNil int

func (*IsNil) Check(actual interface{}) error {
	switch v := reflect.ValueOf(actual); v.Kind() {
	default:
		if nil != actual {
			return errNotEqual(nil, actual)
		}
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr,
		reflect.Slice:
		if !v.IsNil() {
			return errNotEqual(nil, v)
		}
	}
	return nil
}

func (a *Assertion) IsNil() *Assertion {
	return a.Is(new(IsNil))
}

// ----------------------------------------------------------------------------
// Not-Nil Check
// ----------------------------------------------------------------------------

type NotNil int

func (*NotNil) Check(actual interface{}) error {
	if actual != nil {
		return nil
	}
	return errEqual(nil, actual)
}

func (a *Assertion) NotNil() *Assertion {
	return a.Is(new(NotNil))
}

// ----------------------------------------------------------------------------
// Equals check
// ----------------------------------------------------------------------------

type Equals struct {
	value interface{}
}

func (e *Equals) Check(actual interface{}) error {
	if e.value != actual {
		return errNotEqual(e.value, actual)
	}
	return nil
}

func (a *Assertion) Equals(v interface{}) *Assertion {
	return a.Is(&Equals{v})
}

// ----------------------------------------------------------------------------
// Not-Equals check
// ----------------------------------------------------------------------------

type NotEquals struct {
	value interface{}
}

func (e *NotEquals) Check(actual interface{}) error {
	if e.value == actual {
		return errEqual(e.value, actual)
	}
	return nil
}

func (a *Assertion) NotEquals(v interface{}) *Assertion {
	return a.Is(&NotEquals{v})
}

// ----------------------------------------------------------------------------
// Boolean checks
// ----------------------------------------------------------------------------

func (a *Assertion) IsTrue() *Assertion {
	return a.Is(&Equals{true})
}

func (a *Assertion) IsFalse() *Assertion {
	return a.Is(&Equals{false})
}

// ----------------------------------------------------------------------------
// HasLen Check
// ----------------------------------------------------------------------------

type HasLen struct {
	length int
}

func (l *HasLen) Check(val interface{}) error {
	switch v := reflect.ValueOf(val); v.Kind() {
	default:
		return fmt.Errorf("Type %v has no length", v)
	case reflect.Map, reflect.Array, reflect.Slice, reflect.Chan,
		reflect.String:
		if l.length != v.Len() {
			return fmt.Errorf("Expected %v to have length %d, was %d",
				val, l.length, v.Len())
		}
	}
	return nil
}

func (a *Assertion) HasLen(length int) *Assertion {
	return a.Is(&HasLen{length})
}

// ----------------------------------------------------------------------------
// Contains Check
// ----------------------------------------------------------------------------

type Contains struct {
	needle interface{}
}

func (c *Contains) Check(hay interface{}) error {
	switch n := c.needle.(type) {
	case string:
		return c.checkString(n, hay)
	}
	return fmt.Errorf("Type %T(%v) can't contain values", c.needle, c.needle)
}

func (c *Contains) checkString(needle string, hay interface{}) error {
	switch h := hay.(type) {
	case string:
		if strings.Index(h, needle) < 0 {
			// TODO Improve error message
			return fmt.Errorf("String '%s' didn't contain '%s'", hay, needle)
		} else {
			return nil
		}
	}
	return fmt.Errorf("Type %T(%v) not a string", hay, hay)
}

func (a *Assertion) Contains(needle interface{}) *Assertion {
	return a.Is(&Contains{needle})
}
