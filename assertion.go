package assert

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type test struct {
	t     *testing.T
	value interface{}
}

type Assertion interface {
	// TODO: Can probably just be error
	Check(v interface{}) AssertError
}

func That(t *testing.T, value interface{}) *test {
	return &test{
		t:     t,
		value: value,
	}
}

func (t *test) Is(a Assertion) *test {
	return t.is(a)
}

// Interal `is` association for built-in checkers. This is because the call
// stack has a differnet depth, and it would look like errors came from inside
// this file.
// TODO: Recover?
func (t *test) is(a Assertion) *test {
	if err := a.Check(t.value); err != nil {
		msg := err.Error()
		if _, file, line, ok := runtime.Caller(2); ok {
			header := fmt.Sprintf("\n%s:%d\n", file, line)
			if loc, err := readLine(file, line); err == nil {
				header += fmt.Sprintf("  %s\n", strings.TrimSpace(loc))
			}
			msg = header + msg
		}
		t.t.Fatal(msg)
	}
	return t
}

// ----------------------------------------------------------------------------
// Nil Assertion
// ----------------------------------------------------------------------------

func (t *test) IsNil() *test {
	return t.is(Nil())
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
// Not Nil Assertion
// ----------------------------------------------------------------------------

func (t *test) NotNil() *test {
	return t.is(NotNil())
}

type IsNotNil int

func NotNil() *IsNotNil {
	return new(IsNotNil)
}

func (n *IsNotNil) Check(actual interface{}) AssertError {
  if actual != nil {
    return nil
  }
  return EqualityErr("not(nil)", actual)
}

// ----------------------------------------------------------------------------
// Equals Assertion
// ----------------------------------------------------------------------------

func (t *test) Equals(v interface{}) *test {
	return t.is(&IsEqual{v})
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
	return t.is(&IsEqual{true})
}

func (t *test) IsFalse() *test {
	return t.is(&IsEqual{false})
}

// ----------------------------------------------------------------------------
// HasLen Assertion
// ----------------------------------------------------------------------------

func (t *test) HasLen(length int) *test {
	return t.is(&HasLen{length})
}

type HasLen struct {
	length int
}

func (l *HasLen) Check(val interface{}) AssertError {
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
	return t.is(&Contains{i})
}

type Contains struct {
	needle interface{}
}

func (c *Contains) Check(hay interface{}) AssertError {
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

// ----------------------------------------------------------------------------
// Helpers
// ----------------------------------------------------------------------------

func readLine(fname string, lineNo int) (string, error) {
	file, err := os.Open(fname)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for ; scanner.Scan(); i++ {
		if i == lineNo {
			return scanner.Text(), nil
		}
	}
	return "", scanner.Err()
}
