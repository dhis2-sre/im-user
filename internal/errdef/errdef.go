package errdef

import "errors"

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
