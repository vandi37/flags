package flags

import (
	"fmt"
)

type flagError struct {
	name string
	real string
}

func (e *flagError) Error() string {
	return e.real
}

func err(name string, real string) func(args ...any) error {
	return func(args ...any) error {
		return &flagError{name: name, real: fmt.Sprintf(real, args...)}
	}
}

func (e *flagError) Is(target error) bool {
	t, ok := target.(*flagError)
	if !ok {
		return e.name == target.Error()
	}

	return t.name == e.name
}

type megaError struct {
	name string
	errs []error
}

func (e *megaError) Error() string {
	res := e.name + ":"
	for _, err := range e.errs {
		res += "\n	- " + err.Error()
	}

	return res
}

func (e *megaError) Unwrap() []error {
	return e.errs
}

func mega(name string, errs []error) error {
	return &megaError{name: name, errs: errs}
}
