package try

import (
	"errors"
	"testing"
)

func TestCatchingExplicitError(t *testing.T) {
	_, e := Catching(func() any {
		Throw(errors.New("my error"))
		panic("never happened")
	})
	if e == nil || e.Error() != "my error" {
		t.Error("Expected caught error is \"my error\"")
	}
}

func TestCatchingImplicitError(t *testing.T) {
	_, e := Catching(func() any {
		Throw("my error")
		panic("never happened")
	})
	if e == nil {
		t.Error("Expected caught error is \"my error\"")
	}
}

func TestCatchingRegularPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != "my panic" {
			t.Error("Expected panic \"my panic\"")
		}
	}()
	_, _ = Catching(func() any {
		panic("my panic")
	})
	t.Error("`Catching()` should not recover a regular panic")
}
