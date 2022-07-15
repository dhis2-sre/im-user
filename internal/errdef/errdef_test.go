package errdef_test

import (
	"errors"
	"testing"

	"github.com/dhis2-sre/im-user/internal/errdef"
	"github.com/stretchr/testify/assert"
)

func TestAsDuplicate(t *testing.T) {
	assert.False(t, errors.As(errors.New("some error"), &errdef.Duplicate{}))
	assert.True(t, errors.As(errdef.NewDuplicated(errors.New("some error")), &errdef.Duplicate{}))
}

func TestIsUnauthorized(t *testing.T) {
	assert.False(t, errdef.IsUnauthorized(errors.New("some error")))
	assert.True(t, errdef.IsUnauthorized(errdef.NewUnauthorized(errors.New("some error"))))
}

func TestIsNotFound(t *testing.T) {
	assert.False(t, errdef.IsNotFound(errors.New("some error")))
	assert.True(t, errdef.IsNotFound(errdef.NewNotFound(errors.New("some error"))))
}
