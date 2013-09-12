package doc

import (
	assert ".."
	"testing"
)

func TestSimpleStuff(t *testing.T) {
	assert.That(t, 3).Equals(3)

	assert.That(t, "Hi, Billy!").
		Contains("Billy").
		HasLen(10)
}

// ----------------------------------------------------------------------------
// Custom Checks
// ----------------------------------------------------------------------------

type IsThree struct{}

func Three() *IsThree {
	return &IsThree{}
}

func (it *IsThree) Check(actual interface{}) assert.AssertError {
	if actual != 3 {
		return assert.EqualityErr(3, actual)
	}
	return nil
}

func TestIsThree(t *testing.T) {
	assert.That(t, 3).Is(Three())
}

func TestIsNotThree(t *testing.T) {
	assert.That(t, "two").Is(Three())
}
