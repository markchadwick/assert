# Assert
Not-too-complicated test assertions for Go.

## Usage

```go
package mine

import (
  "github.com/markchadwick/assert"
  "testing"
)

func TestSimpleStuff(t *testing.T) {
  assert.That(t, 3).Equals(3)

  assert.That(t, "Hi, Billy!").
    Contains("Billy").
    HasLen(10)
}
```

## Built-in Checks
Currently a bit in flux, but a quick overview,

* Is(a Assertion)
* IsNil()
* NotNil()
* Equals(v interface{})
* IsTrue()
* IsFalse()
* HasLen(length int)
* Contains(v interface{}) // Only works for a few types

## Using Custom checks
A custom check just needs a `Check(actual interface{}) error` method. Below is
an example of a custom checker that ensures values are the integer 3.

```go
type IsThree struct {}

func Three() *IsThree {
  return &IsThree{}
}

func (it *IsThree) Check(actual interface{}) assert.AssertError {
  if actual != 3 {
    return assert.EqualityErr(3, actual)
  }
  return nil
}
```

And using it from your tests would look like the following:

```go
func TestIsThree(t *testing.T) {
  assert.That(t, 3).Is(Three())
}

func TestIsNotThree(t *testing.T) {
  assert.That(t, "two").Is(Three())
}
```

```
--- FAIL: TestIsNotThree-2 (0.00 seconds)
        assertion.go:48:
                /doc/doc_test.go:38
                  assert.That(t, "two").Is(Three())

                Expected: 3
                Actual:   two
FAIL
```
