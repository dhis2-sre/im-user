package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetUserFromContext_Happy(t *testing.T) {
	var id uint = 0
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("user", &model.User{
		Model: gorm.Model{ID: id},
	})

	user, err := GetUserFromContext(c)

	assert.NoError(t, err)
	assert.Equal(t, id, user.ID)
}

func TestGetUserFromContext_NoUserOnContext(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	_, err := GetUserFromContext(c)

	assert.ErrorContains(t, err, "unable to extract user from request context for unknown reason")
}

func TestGetUserFromContext_InvalidUserOnContext(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("user", "not a user struct")

	_, err := GetUserFromContext(c)

	assert.ErrorContains(t, err, "unable to cast user for unknown reason")
}
