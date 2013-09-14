package assert

import (
	"strings"
	"testing"
)

// ----------------------------------------------------------------------------
// Nil Check

func TestSimpleNil(t *testing.T) {
	check := new(IsNil)

	err := check.Check(nil)
	if err != nil {
		t.Fatalf("Should not have failed: %s", err.Error())
	}
}

func TestTypedNil(t *testing.T) {
	var nothing *Assertion // type doesn't matter here
	if nothing != nil {
		t.Fatal("My assumptions were wrong!")
	}

	err := new(IsNil).Check(nothing)
	if err != nil {
		t.Fatalf("Should not have failed: %s", err.Error())
	}
}

func TestTestNil(t *testing.T) {
	That(t, nil).IsNil()
}

func TestNilMessage(t *testing.T) {
	msg := Assert(3).IsNil().Check().Error()
	expected := "expected:     <nil>\n" +
		"received int: 3"
	if msg != expected {
		t.Fatalf("Didn't like the error message: %s", msg)
	}
}

// ----------------------------------------------------------------------------
// Not-Nil Check

func TestNotNil(t *testing.T) {
	err := new(NotNil).Check(3)
	That(t, err).IsNil()

	msg := new(NotNil).Check(nil).Error()
	exp := "expected not: <nil>\n" +
		"received:     <nil>"
	if msg != exp {
		t.Fatalf("Expected: '%s', got '%s'", exp, msg)
	}
}

func TestNotNilTest(t *testing.T) {
	t2 := new(testing.T)

	That(t2, 3).NotNil()
	if t2.Failed() {
		t.Fatalf("should not have failed")
	}

	That(t2, nil).NotNil()
	if !t2.Failed() {
		t.Fatalf("should not have failed")
	}
}

// ----------------------------------------------------------------------------
// Equals Check

func TestSimpleEquals(t *testing.T) {
	if err := (&Equals{3}).Check(3); err != nil {
		t.Fatal(err)
	}
	if err := (&Equals{3}).Check(5); err == nil {
		t.Fatal()
	}

	if err := (&Equals{"hello"}).Check("hello"); err != nil {
		t.Fatal(err)
	}

	if err := (&Equals{true}).Check(true); err != nil {
		t.Fatal(err)
	}
	if err := (&Equals{true}).Check(false); err == nil {
		t.Fatal()
	}
}

// Different types should fail equality matches
func TestEqualsDifferentNativeTypes(t *testing.T) {
	if err := (&Equals{int64(3)}).Check(int64(3)); err != nil {
		t.Fatal(err)
	}

	if err := (&Equals{int32(3)}).Check(int64(3)); err == nil {
		t.Fatal()
	}
}

// Typed nils should fail equality matches
func TestEqualsNil(t *testing.T) {
	if err := (&Equals{nil}).Check(nil); err != nil {
		t.Fatal(err)
	}

	var nothing *Assertion // type doesn't matter here
	if err := (&Equals{nothing}).Check(nil); err == nil {
		t.Fatal()
	}
}

func TestEqualsMessage(t *testing.T) {
	msg := (&Equals{"bon"}).Check("jour").Error()
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

	if err := (&Equals{s}).Check(s); err != nil {
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
// Not-equals Check

func TestNotEquals(t *testing.T) {
	err := (&NotEquals{"bon"}).Check("jour")
	That(t, err).IsNil()

	That(t, 123).NotEquals(false)

	err = (&NotEquals{123}).Check(123)
	exp := "expected not int: 123\n" +
		"received int:     123"

	That(t, err.Error()).Equals(exp)
}

// ----------------------------------------------------------------------------
// Boolean checks

func TestIsTrue(t *testing.T) {
	That(t, true).IsTrue()
}

func TestIsFalse(t *testing.T) {
	That(t, false).IsFalse()
}

// ----------------------------------------------------------------------------
// HasLen

func TestSimpleHasLen(t *testing.T) {
	if err := (&HasLen{0}).Check([]string{}); err != nil {
		t.Fatal(err)
	}

	if err := (&HasLen{1}).Check([]string{}); err == nil {
		t.Fatal()
	}

	err := (&HasLen{0}).Check([]string{"one", "two", "three"}).Error()

	if strings.Index(err, "[one two three]") < 0 {
		t.Fatalf("Expected to have what was received...that's deep")
	}
}

func TestNoLength(t *testing.T) {
	err := (&HasLen{0}).Check(666)
	if err == nil {
		t.Fatal("Should have failed")
	}
	if strings.Index(err.Error(), "Type <int Value> has no length") < 0 {
		t.Fatal("Should have failed with a good reason")
	}
}

// ----------------------------------------------------------------------------
// Contains

func TestContains(t *testing.T) {
	That(t, "johnny").Contains("john")
}

func TestContainsBadType(t *testing.T) {
	err := (&Contains{3}).Check("Hi")
	if err == nil {
		t.Fail()
	}
	That(t, err.Error()).Equals("Type int(3) can't contain values")

	err = (&Contains{"Hi"}).Check(3)
	if err == nil {
		t.Fail()
	}
	That(t, err.Error()).Equals("Type int(3) not a string")
}

func TestContainsMap(t *testing.T) {
	t.Skip()
	if err := (&Contains{"one"}).Check(map[string]int{"one": 1}); err != nil {
		t.Fatal(err)
	}
}
