package errdef

import (
	"errors"
)

func NewDuplicated(err error) error {
	return Duplicate{err}
}

type Duplicate struct{ error }

func (d Duplicate) As(target any) bool {
	return errors.As(d, &target)
}

type unauthorized struct{ error }

func NewUnauthorized(err error) error {
	return unauthorized{err}
}

func IsUnauthorized(err error) bool {
	var e unauthorized
	return errors.As(err, &e)
}

type notFound struct{ error }

// NewNotFound creates an error representing a resource that could not be found.
func NewNotFound(err error) error {
	return notFound{err}
}

// IsNotFound returns true if err is an error representing a resource that could not be found and false otherwise.
func IsNotFound(err error) bool {
	var e notFound
	return errors.As(err, &e)
}
