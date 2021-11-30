package middleware

import (
	"github.com/dhis2-sre/im-users/internal/apperror"
	"github.com/dhis2-sre/im-users/pgk/helper"
	"github.com/dhis2-sre/im-users/pgk/user"
	"github.com/gin-gonic/gin"
	"log"
)

func ProvideAuthorization(userService user.Service) AuthorizationMiddleware {
	return AuthorizationMiddleware{userService}
}

type AuthorizationMiddleware struct {
	userService user.Service
}

func (m AuthorizationMiddleware) RequireAdministrator(c *gin.Context) {
	u, err := helper.GetUserFromContext(c)
	if err != nil {
		_ = c.Error(err)
		return
	}

	userWithGroups, _ := m.userService.FindById(u.ID)

	if !userWithGroups.IsAdministrator() {
		log.Printf("User tried to access administrator restricted endpoint: %+v\n", u)
		unauthorized := apperror.NewUnauthorized("Administrator access denied")
		_ = c.Error(unauthorized)
		return
	}

	c.Next()
}
