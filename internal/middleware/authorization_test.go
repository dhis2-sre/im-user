package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dhis2-sre/im-user/internal/errdef"

	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAuthorizationMiddleware_RequireAdministrator_Happy(t *testing.T) {
	var id uint = 1
	email := "someone@something.org"
	password := "password"

	userService := &mockUserService{}
	userService.On("FindById", mock.AnythingOfType("uint")).Return(&model.User{
		Model:    gorm.Model{ID: id},
		Email:    email,
		Password: password,
		Groups: []model.Group{
			{Name: model.AdministratorGroupName},
		},
	}, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("user", &model.User{
		Model: gorm.Model{ID: id},
	})
	authorization := NewAuthorization(userService)

	authorization.RequireAdministrator(c)

	assert.False(t, c.IsAborted())

	errs := c.Errors.Errors()
	assert.Equal(t, 0, len(errs))
}

func TestAuthorizationMiddleware_RequireAdministrator_NotInAdministratorGroup(t *testing.T) {
	var id uint = 1
	email := "someone@something.org"
	password := "password"

	userService := &mockUserService{}
	userService.On("FindById", mock.AnythingOfType("uint")).Return(&model.User{
		Model:    gorm.Model{ID: id},
		Email:    email,
		Password: password,
	}, nil)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("user", &model.User{
		Model: gorm.Model{ID: id},
	})
	authorization := NewAuthorization(userService)

	authorization.RequireAdministrator(c)

	assert.True(t, c.IsAborted())

	errs := c.Errors.Errors()
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, "administrator access denied", errs[0])
}

func TestAuthorizationMiddleware_RequireAdministrator_UserNotFound(t *testing.T) {
	id := uint(0)

	userService := &mockUserService{}
	errorMessage := "not found"
	userService.
		On("FindById", id).
		Return(nil, errors.New(errorMessage))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("user", &model.User{
		Model: gorm.Model{ID: id},
	})
	authorization := NewAuthorization(userService)

	authorization.RequireAdministrator(c)

	assert.True(t, c.IsAborted())

	errs := c.Errors.Errors()
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, errorMessage, errs[0])
}

func TestAuthorizationMiddleware_RequireAdministrator_UserNotFoundError(t *testing.T) {
	id := uint(0)

	userService := &mockUserService{}
	errorMessage := "not found"
	userService.
		On("FindById", id).
		Return(nil, errdef.NewNotFound(fmt.Errorf(errorMessage)))

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("user", &model.User{
		Model: gorm.Model{ID: id},
	})
	authorization := NewAuthorization(userService)

	authorization.RequireAdministrator(c)

	assert.True(t, c.IsAborted())

	errs := c.Errors.Errors()
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, errorMessage, errs[0])
	lastError := c.Errors.Last()
	assert.True(t, errdef.IsNotFound(lastError))
}

func TestAuthorizationMiddleware_RequireAdministrator_UserNotOnContext(t *testing.T) {
	userService := &mockUserService{}
	errorMessage := "unable to extract user from request context for unknown reason: "
	authorization := NewAuthorization(userService)

	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	authorization.RequireAdministrator(c)

	assert.True(t, c.IsAborted())

	errs := c.Errors.Errors()
	assert.Equal(t, 1, len(errs))
	assert.True(t, strings.HasPrefix(errs[0], errorMessage))
}

func TestAuthorizationMiddleware_RequireAdministrator_ExternalError(t *testing.T) {
	var id uint = 1
	email := "someone@something.org"
	password := "password"

	req, err := http.NewRequest(http.MethodPost, "/whatever", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer token")

	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = req

	userService := &mockUserService{}
	userService.
		On("FindById", mock.AnythingOfType("uint")).
		Return(&model.User{
			Model:    gorm.Model{ID: id},
			Email:    email,
			Password: password,
			Groups: []model.Group{
				{Name: model.AdministratorGroupName},
			},
		}, nil)
	authorization := NewAuthorization(userService)

	_ = c.Error(errors.New("some error which wasn't handled properly"))

	authorization.RequireAdministrator(c)

	assert.True(t, c.IsAborted())
}
