package try

// CatchingAll run block and catch all panics.
func CatchingAll[R any](block func() R) (R, any) {
	var r any
	ret := func() R {
		defer tryRecoverAll(&r)
		return block()
	}()
	return ret, r
}

// CatchingAll0 run block and catch all panics.
func CatchingAll0(block func()) any {
	var r any
	func() {
		defer tryRecoverAll(&r)
		block()
	}()
	return r
}

// CatchingAll2 run block and catch all panics.
func CatchingAll2[R1 any, R2 any](block func() (R1, R2)) (R1, R2, any) {
	var r any
	r1, r2 := func() (R1, R2) {
		defer tryRecoverAll(&r)
		return block()
	}()
	return r1, r2, r
}

func tryRecoverAll(r *any) {
	*r = recover()
}
