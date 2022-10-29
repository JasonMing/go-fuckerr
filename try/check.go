package try

// Check Throw if err != nil.
func Check(err error) {
	if err != nil {
		Throw(err)
	}
}

// Must return ret if err != nil, or Throw err.
func Must[T any](ret T, err error) T {
	if err != nil {
		Throw(err)
	}
	return ret
}

// Must2 return 2 values(ret1, ret2) if err != nil, or Throw err.
func Must2[T1 any, T2 any](ret1 T1, ret2 T2, err error) (T1, T2) {
	if err != nil {
		Throw(err)
	}
	return ret1, ret2
}
