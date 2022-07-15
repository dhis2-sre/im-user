package errdef

import (
	"errors"
	"fmt"
)

type duplicate struct{ error }

func NewDuplicated(format string, a ...any) error {
	return duplicate{fmt.Errorf(format, a...)}
}

func IsDuplicated(err error) bool {
	var e duplicate
	return errors.As(err, &e)
}

type unauthorized struct{ error }

func NewUnauthorized(format string, a ...any) error {
	return unauthorized{fmt.Errorf(format, a...)}
}

func IsUnauthorized(err error) bool {
	var e unauthorized
	return errors.As(err, &e)
}

type notFound struct{ error }

func (e notFound) NotFound() {}

func (e notFound) Unwrap() error {
	return e.error
}

// NotFound creates an error representing a resource that could not be found.
func NotFound(err error) error {
	if err == nil || IsNotFound(err) {
		return err
	}
	return notFound{err}
}

// IsNotFound returns true if err is an error representing a resource that could not be found and false otherwise.
func IsNotFound(err error) bool {
	var e notFound
	return errors.As(err, &e)
}
