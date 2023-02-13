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
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	u, ok := user.(*model.User)
	if !ok {
		err := fmt.Errorf("unable to cast user for unknown reason: %+v", c)
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return nil, err
	}

	return u, nil
}
