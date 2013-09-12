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

## Using Custom checks
A custom check just needs a `Check(actual interface{}) error` method. Below is
an example of a custom checker that ensures values are the integer 3.

```go
type IsThree struct {}

func Three() *IsThree {
  return &IsThree{}
}

func (it *IsThree) Check(actual interface{}) assert.AssertError {
  if actual == 3 {
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
