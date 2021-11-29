package middleware

import (
	"github.com/dhis2-sre/im-users/internal/apperror"
	"github.com/dhis2-sre/im-users/pgk/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProvideAuthenticationMiddleware(userService user.Service) AuthenticationMiddleware {
	return AuthenticationMiddleware{
		userService,
	}
}

type AuthenticationMiddleware struct {
	userService user.Service
}

// BasicAuthentication Inspiration: https://www.pandurang-waghulde.com/custom-http-basic-authentication-using-gin/
func (m AuthenticationMiddleware) BasicAuthentication(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()

	if !ok {
		unauthorized := apperror.NewUnauthorized("Invalid Authorization header format")
		m.handleError(c, unauthorized)
		return
	}

	u, err := m.userService.SignIn(username, password)
	if err != nil {
		unauthorized := apperror.NewUnauthorized("Invalid credentials")
		m.handleError(c, unauthorized)
		return
	}

	c.Set("user", u)
	c.Next()
}

func (m AuthenticationMiddleware) handleError(c *gin.Context, e error) {
	c.Header("WWW-Authenticate", "Basic realm=\"DHIS2\"")
	_ = c.AbortWithError(http.StatusUnauthorized, e)
}
