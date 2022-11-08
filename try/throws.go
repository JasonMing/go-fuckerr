package try

import (
	"fmt"

	"github.com/pkg/errors"
)

// Throw raise a panic which will be caught by outside Catching[n]
// function then convert back to error in result.
func Throw(v ...any) {
	if len(v) == 0 {
		panic(throwable{errors.New("")})
	}
	if len(v) == 1 {
		if err, ok := v[0].(error); ok {
			panic(throwable{errors.WithStack(err)})
		}
	}
	if err, ok := v[0].(error); ok {
		panic(throwable{errors.Wrap(err, fmt.Sprint(v[1:]...))})
	} else {
		panic(throwable{errors.New(fmt.Sprint(v...))})
	}
}

// Throwf raise a panic with formatted string which will be caught
// by outside Catching[n] function then convert back to error in result.
func Throwf(format string, a ...any) {
	panic(throwable{errors.New(fmt.Sprintf(format, a...))})
}
