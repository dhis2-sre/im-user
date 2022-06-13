package handler

import (
	"fmt"
	"net/http"

	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/gin-gonic/gin"
)

func GetUserFromContext(c *gin.Context) (*model.User, error) {
	user, exists := c.Get("user")

	if !exists {
		err := fmt.Errorf("unable to extract user from request context for unknown reason: %+v", c)
		c.AbortWithError(http.StatusInternalServerError, err) // nolint:errcheck
		return nil, err
	}

	return user.(*model.User), nil
}
