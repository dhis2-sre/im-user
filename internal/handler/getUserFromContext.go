package handler

import (
	"fmt"
	"github.com/dhis2-sre/im-users/internal/apperror"
	"github.com/dhis2-sre/im-users/pkg/model"
	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) (*model.User, error) {
	user, exists := c.Get("user")

	if !exists {
		message := fmt.Sprintf("Unable to extract user from request context for unknown reason: %+v", c)
		err := apperror.NewInternal(message)
		return nil, err
	}

	return user.(*model.User), nil
}
