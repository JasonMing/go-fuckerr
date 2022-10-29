package try

import (
	"errors"
	"fmt"
)

// Throw raise a panic which will be caught by outside Catching[n]
// function then convert back to error in result.
func Throw(v ...any) {
	if len(v) == 1 {
		if err, ok := v[0].(error); ok {
			panic(notPanic{err})
		}
	}
	panic(notPanic{errors.New(fmt.Sprint(v...))})
}

// Throwf raise a panic with formatted string which will be caught
// by outside Catching[n] function then convert back to error in result.
func Throwf(format string, a ...any) {
	panic(notPanic{errors.New(fmt.Sprintf(format, a...))})
}
