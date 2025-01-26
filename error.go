package flags

import (
	"fmt"
)

type FlagError struct {
	name string
	real string
}

func (e *FlagError) Error() string {
	return e.real
}

func Err(name string, real string) func(args ...any) error {
	return func(args ...any) error {
		return &FlagError{name: name, real: fmt.Sprintf(real, args...)}
	}
}

func (e *FlagError) Is(target error) bool {
	t, ok := target.(*FlagError)
	if !ok {
		return e.name == target.Error()
	}

	return t.name == e.name
}

type MegaError struct {
	name string
	errs []error
}

func (e *MegaError) Error() string {
	res := e.name + ":"
	for _, err := range e.errs {
		res += "\n	- " + err.Error()
	}

	return res
}

func (e *MegaError) Unwrap() []error {
	return e.errs
}

func Mega(name string, errs []error) error {
	return &MegaError{name: name, errs: errs}
}
