package try

// NoErr Provide a idiomatic nil on return statement.
//
// For Example:
//
//	return result, NoErr
var NoErr error = nil

// Catching run block and catch all panics which are raised by Throw[f] function.
func Catching[R any](block func() R) (R, error) {
	var err error
	ret := func() R {
		defer tryRecover(&err)
		return block()
	}()
	return ret, err
}

// Catching0 run block and catch all panics which are raised by Throw[f] function.
func Catching0(block func()) error {
	var err error
	func() {
		defer tryRecover(&err)
		block()
	}()
	return err
}

// Catching2 run block and catch all panics which are raised by Throw[f] function.
func Catching2[R1 any, R2 any](block func() (R1, R2)) (R1, R2, error) {
	var err error
	r1, r2 := func() (R1, R2) {
		defer tryRecover(&err)
		return block()
	}()
	return r1, r2, err
}

func tryRecover(err *error) {
	switch r := recover().(type) {
	case nil:
		return
	case notPanic:
		*err = r.Err
	case *notPanic:
		*err = r.Err
	default:
		panic(r) // is not notPanic, re-panic it
	}
}

type notPanic struct {
	Err error
}

func (this *notPanic) Error() string {
	return this.Err.Error()
}

// IsNotPanic check whether err is raised by Throw[f].
func IsNotPanic(err any) bool {
	switch err.(type) {
	case notPanic, *notPanic:
		return true
	default:
		return false
	}
}
