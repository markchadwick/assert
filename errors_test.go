package assert

import (
	"testing"
)

func TestErrorPreamble(t *testing.T) {
	msg := preamble("expected", 3)
	exp := "expected int: "

	if msg != exp {
		t.Fatalf("Expected: '%s', got '%s'", exp, msg)
	}
}

func TestErrorPreambleNil(t *testing.T) {
	msg := preamble("expected", nil)
	exp := "expected: "

	if msg != exp {
		t.Fatalf("Expected: '%s', got '%s'", exp, msg)
	}
}

func TestErrNotEqual(t *testing.T) {
	msg := errNotEqual("three", 3).Error()
	exp := "expected string: three\n" +
		"received int:    3"

	if msg != exp {
		t.Fatalf("Expected: '%s', got '%s'", exp, msg)
	}
}

func TestErrEqual(t *testing.T) {
	msg := errEqual("three", 3).Error()
	exp := "expected not string: three\n" +
		"received int:        3"

	if msg != exp {
		t.Fatalf("Expected: '%s', got '%s'", exp, msg)
	}
}
