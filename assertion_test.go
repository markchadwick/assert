package assert

import (
	"fmt"
	"strings"
	"testing"
)

func TestThat(t *testing.T) {
	test := That(t, 3)
	if test.t != t {
		t.Fatalf("expected %v, got %v", t, test.t)
	}

	if test.value != 3 {
		t.Fatalf("expected %v, got %v", 3, test.value)
	}
}

type SimpleError struct {
	expected interface{}
	actual   interface{}
}

func (se *SimpleError) Expected() interface{} {
	return se.expected
}

func (se *SimpleError) Actual() interface{} {
	return se.actual
}

func (se *SimpleError) Error() string {
	return fmt.Sprintf("Expected '%v', got '%v'", se.expected, se.actual)
}

type SimpleAssert struct {
	v interface{}
}

func Simple(v interface{}) *SimpleAssert {
	return &SimpleAssert{v}
}

func (s *SimpleAssert) Check(actual interface{}) AssertError {
	if s.v != actual {
		return &SimpleError{s.v, actual}
	}
	return nil
}

func TestSimpleAssertion(t *testing.T) {
	simple := &SimpleAssert{"hi"}
	err := simple.Check("hi")
	if err != nil {
		t.Fatal(err)
	}

	err = simple.Check("there")
	expected := "Expected 'hi', got 'there'"
	if err.Error() != expected {
		t.Fatalf("Got '%s', expected '%s'", err.Error(), expected)
	}
}

func TestSimpleIs(t *testing.T) {
	t2 := new(testing.T)
	That(t2, "hi").Is(Simple("hi"))
	if t2.Failed() {
		t.Errorf("Should not have failed")
	}

	That(t2, "hi").Is(Simple("there"))
	if !t2.Failed() {
		t.Errorf("Should have failed")
	}
}

func TestFailEquality(t *testing.T) {
	t2 := new(testing.T)
	That(t2, "hi").Is(Nil())
	if !t2.Failed() {
		t.Error("Should have failed")
	}
}

// ----------------------------------------------------------------------------
// Nil Assertion
// ----------------------------------------------------------------------------

func TestSimpleNil(t *testing.T) {
	err := Nil().Check(nil)
	if err != nil {
		t.Fatalf("Should not have failed: %s", err.Error())
	}
}

func TestTypedNil(t *testing.T) {
	var nothing *test
	if nothing != nil {
		t.Fatal("My assumptions were wrong!")
	}

	err := Nil().Check(nothing)
	if err != nil {
		t.Fatalf("Should not have failed: %s", err.Error())
	}
}

func TestTestNil(t *testing.T) {
	That(t, nil).IsNil()
}

// ----------------------------------------------------------------------------
// Equals Assertion
// ----------------------------------------------------------------------------

func TestSimpleEquals(t *testing.T) {
	if err := (&IsEqual{3}).Check(3); err != nil {
		t.Fatal(err)
	}
	if err := (&IsEqual{3}).Check(5); err == nil {
		t.Fatal()
	}

	if err := (&IsEqual{"hello"}).Check("hello"); err != nil {
		t.Fatal(err)
	}

	if err := (&IsEqual{true}).Check(true); err != nil {
		t.Fatal(err)
	}
	if err := (&IsEqual{true}).Check(false); err == nil {
		t.Fatal()
	}
}

// Different types should fail equality matches
func TestEqualsDifferentNativeTypes(t *testing.T) {
	if err := (&IsEqual{int64(3)}).Check(int64(3)); err != nil {
		t.Fatal(err)
	}

	if err := (&IsEqual{int32(3)}).Check(int64(3)); err == nil {
		t.Fatal()
	}
}

// Typed nils should fail equality matches
func TestEqualsNil(t *testing.T) {
	if err := (&IsEqual{nil}).Check(nil); err != nil {
		t.Fatal(err)
	}

	var nothing *test
	if err := (&IsEqual{nothing}).Check(nil); err == nil {
		t.Fatal()
	}
}

func TestEqualsMessage(t *testing.T) {
	msg := (&IsEqual{"bon"}).Check("jour").Error()
	if strings.Index(msg, "bon") < 0 {
		t.Fatalf("String '%s' should have contained 'bon'", msg)
	}
	if strings.Index(msg, "jour") < 0 {
		t.Fatalf("String '%s' should have contained 'jour'", msg)
	}
}

// TODO: This panics
func TestEqualsSlice(t *testing.T) {
	defer func() {
		recover()
	}()

	s := []byte{0, 1, 2}

	if err := (&IsEqual{s}).Check(s); err != nil {
		t.Fatal(err)
	}
}

func TestTestEquals(t *testing.T) {
	That(t, 5).Equals(5)
	That(t, "hi").Equals("hi")
	That(t, true).Equals(true)
	That(t, false).Equals(false)
}

// ----------------------------------------------------------------------------
// Boolean Assertions
// ----------------------------------------------------------------------------

func TestIsTrue(t *testing.T) {
	That(t, true).IsTrue()
}

func TestIsFalse(t *testing.T) {
	That(t, false).IsFalse()
}

// ----------------------------------------------------------------------------
// HasLen Assertion
// ----------------------------------------------------------------------------

func TestSimpleLengthed(t *testing.T) {
	if err := (&Lengthed{0}).Check([]string{}); err != nil {
		t.Fatal(err)
	}

	if err := (&Lengthed{1}).Check([]string{}); err == nil {
		t.Fatal()
	}

	err := (&Lengthed{0}).Check([]string{"one", "two", "three"}).Error()

	if strings.Index(err, "[one two three]") < 0 {
		t.Fatalf("Expected to have what was received...that's deep")
	}
}

func TestNoLength(t *testing.T) {
	err := (&Lengthed{0}).Check(666)
	if err == nil {
		t.Fatal("Should have failed")
	}
	if strings.Index(err.Error(), "Type <int Value> has no length") < 0 {
		t.Fatal("Should have failed with a good reason")
	}
}

// ----------------------------------------------------------------------------
// Contains Assertion
// ----------------------------------------------------------------------------

func TestContains(t *testing.T) {
	That(t, "johnny").Contains("john")
}

func TestContainsMap(t *testing.T) {
	t.Logf("pending")
}

func TestContainsArray(t *testing.T) {
	t.Logf("pending")
}

func TestContainsSlice(t *testing.T) {
	t.Logf("pending")
}

func TestContainsChan(t *testing.T) {
	t.Logf("pending")
}

func TestContainsString(t *testing.T) {
	if err := (&Containing{"Hello, Steve"}).Check("Steve"); err != nil {
		t.Fatal(err)
	}

	if err := (&Containing{"Hello, Steve"}).Check("Jimmy"); err == nil {
		t.Fatal("String didn't contain 'Jimmy'")
	}
}
