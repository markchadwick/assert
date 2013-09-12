package assert

import (
	"testing"
)

func TestRuntimeError(t *testing.T) {
	err := RuntimeErr("crrrap!")

	if err.Expected() != nil {
		t.Fatal("Expectedshould have been nil")
	}

	if err.Actual() != nil {
		t.Fatal("Actual should have been nil")
	}
}
