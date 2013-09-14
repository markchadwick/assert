package assert

import (
	"errors"
	"testing"
)

func TestEmptyAssertion(t *testing.T) {
	err := Assert(nil).Check()
	if err != nil {
		t.Fatal(err)
	}
}

func TestThat(t *testing.T) {
	test := That(t, 3)

	err := test.Check()
	if err != nil {
		t.Fatal(err)
	}
}

type AlwaysPass struct {
	checks int
}

func (c *AlwaysPass) Check(actual interface{}) error {
	c.checks++
	return nil
}

type AlwaysFail struct {
	checks int
	last   error
}

func (c *AlwaysFail) Check(actual interface{}) error {
	c.checks++
	c.last = errors.New("I said always failing")
	return c.last
}

func TestPassCheck(t *testing.T) {
	a := Assert(nil).Is(&AlwaysPass{})

	if err := a.Check(); err != nil {
		t.Fatal(err)
	}
}

func TestFailCheck(t *testing.T) {
	a := Assert(nil).Is(&AlwaysFail{})

	if err := a.Check(); err == nil {
		t.Fatalf("Should have failed")
	}
}

func TestPassCheckTesting(t *testing.T) {
	t2 := new(testing.T)
	That(t2, nil).Is(&AlwaysPass{})

	if t2.Failed() {
		t.Fatalf("Test should not have failed")
	}
}

func TestFailCheckTesting(t *testing.T) {
	t2 := new(testing.T)
	That(t2, nil).Is(&AlwaysFail{})

	if !t2.Failed() {
		t.Fatalf("Test should have failed")
	}
}

func TestChaining(t *testing.T) {
	pass := &AlwaysPass{}
	fail := &AlwaysFail{}

	a := Assert(nil).
		Is(pass).
		Is(fail)

	if pass.checks != 0 {
		t.Fatalf("called pass %d times, wanted 0", pass.checks)
	}
	if fail.checks != 0 {
		t.Fatalf("called fail %d times, wanted 0", fail.checks)
	}

	err := a.Check()
	if err == nil {
		t.Fatalf("Should have failed")
	}

	if pass.checks != 1 {
		t.Fatalf("called pass %d times, wanted 1", pass.checks)
	}
	if fail.checks != 1 {
		t.Fatalf("called fail %d times, wanted 1", fail.checks)
	}

	if err != fail.last {
		t.Fatalf("Not sure who's failure this is: %s", err.Error())
	}
}

func TestChainingTesting(t *testing.T) {
	t2 := new(testing.T)

	That(t2, nil).Is(&AlwaysPass{})

	if t2.Failed() {
		t.Fatal("should not have failed")
	}

	That(t2, nil).
		Is(&AlwaysPass{}).
		Is(&AlwaysFail{})

	if !t2.Failed() {
		t.Fatal("should have failed")
	}
}
