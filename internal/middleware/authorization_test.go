package middleware

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAuthorizationMiddleware_RequireAdministrator_Happy(t *testing.T) {
	id := uint(1)
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

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user", &model.User{
		Model: gorm.Model{ID: id},
	})
	authorization := NewAuthorization(userService)

	authorization.RequireAdministrator(c)

	errs := c.Errors.Errors()
	assert.Equal(t, 0, len(errs))
}

func TestAuthorizationMiddleware_RequireAdministrator_NotInAdminstratorGroup(t *testing.T) {
	id := uint(1)
	email := "someone@something.org"
	password := "password"

	userService := &mockUserService{}
	userService.On("FindById", mock.AnythingOfType("uint")).Return(&model.User{
		Model:    gorm.Model{ID: id},
		Email:    email,
		Password: password,
	}, nil)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user", &model.User{
		Model: gorm.Model{ID: id},
	})
	authorization := NewAuthorization(userService)

	authorization.RequireAdministrator(c)

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

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("user", &model.User{
		Model: gorm.Model{ID: id},
	})
	authorization := NewAuthorization(userService)

	authorization.RequireAdministrator(c)

	errs := c.Errors.Errors()
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, errorMessage, errs[0])
}
