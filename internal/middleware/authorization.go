package middleware

import (
	"github.com/dhis2-sre/im-user/internal/apperror"
	"github.com/dhis2-sre/im-user/internal/handler"
	"github.com/dhis2-sre/im-user/pkg/user"
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
	u, err := handler.GetUserFromContext(c)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	userWithGroups, _ := m.userService.FindById(u.ID)

	if !userWithGroups.IsAdministrator() {
		log.Printf("User tried to access administrator restricted endpoint: %+v\n", u)
		unauthorized := apperror.NewUnauthorized("Administrator access denied")
		_ = c.Error(unauthorized)
		c.Abort()
		return
	}

	// Extra precaution to ensure that no errors has occurred, and it's safe to call c.Next()
	if len(c.Errors.Errors()) > 0 {
		c.Abort()
		return
	} else {
		c.Next()
	}
}
