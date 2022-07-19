package errdef

import (
	"errors"
)

type duplicated struct{ error }

func NewDuplicated(err error) error {
	return duplicated{err}
}

func IsDuplicated(err error) bool {
	var e duplicated
	return errors.As(err, &e)
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
