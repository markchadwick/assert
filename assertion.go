package assert

import (
	"fmt"
	"testing"
)

type CheckAdded func(check Check)

type Assertion struct {
	value      interface{}
	checks     []Check
	CheckAdded CheckAdded
}

func Assert(value interface{}) *Assertion {
	return &Assertion{
		value:  value,
		checks: make([]Check, 0),
	}
}

func That(t *testing.T, value interface{}) *Assertion {
	assert := Assert(value)
	assert.CheckAdded = func(check Check) {
		if err := assert.CheckOne(check); err != nil {
			assert.failTest(t, err)
		}
	}
	return assert
}

func (a *Assertion) Is(check Check) *Assertion {
	a.checks = append(a.checks, check)
	if a.CheckAdded != nil {
		a.CheckAdded(check)
	}
	return a
}

func (a *Assertion) Check() error {
	for _, check := range a.checks {
		if err := a.CheckOne(check); err != nil {
			return err
		}
	}
	return nil
}

func (a *Assertion) CheckOne(check Check) error {
	return check.Check(a.value)
}

// TODO: Line numbers are janked in error messages
func (a *Assertion) failTest(t *testing.T, err error) {
	t.Fatal(fmt.Errorf("\n%s", err.Error()))
}
