package try

import (
	"fmt"

	"github.com/pkg/errors"
)

type throwable struct {
	error
}

func (this *throwable) Is(target error) bool {
	return errors.Is(this.error, target)
}

func (this *throwable) As(target any) bool {
	return errors.As(this.error, target)
}

func (this *throwable) Unwrap() error {
	return errors.Unwrap(this.error)
}

func (this *throwable) Cause() error {
	return errors.Cause(this.error)
}

func (this *throwable) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('-') {
			_, _ = fmt.Fprintf(s, "%v", this.error)
		} else {
			_, _ = fmt.Fprintf(s, "%+v", this.error)
		}
	default:
		_, _ = fmt.Fprintf(s, "%"+string(verb), this.error)
	}
}
