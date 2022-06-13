package errdef_test

import (
	"errors"
	"testing"

	"github.com/dhis2-sre/im-user/internal/errdef"
	"github.com/stretchr/testify/assert"
)

func TestIsNotFound(t *testing.T) {
	assert.False(t, errdef.IsNotFound(errors.New("some error")))
	assert.True(t, errdef.IsNotFound(errdef.NotFound(errors.New("some error"))))
}
