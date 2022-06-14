package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dhis2-sre/im-user/pkg/token"
	"github.com/dhis2-sre/im-user/pkg/user"
	"github.com/gin-gonic/gin"
)

type AuthenticationMiddleware struct {
	userService  user.Service
	tokenService token.Service
}

func NewAuthentication(userService user.Service, tokenService token.Service) AuthenticationMiddleware {
	return AuthenticationMiddleware{
		userService,
		tokenService,
	}
}

// BasicAuthentication Inspiration: https://www.pandurang-waghulde.com/custom-http-basic-authentication-using-gin/
func (m AuthenticationMiddleware) BasicAuthentication(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		m.handleError(c, errors.New("invalid Authorization header format"))
		return
	}

	u, err := m.userService.SignIn(username, password)
	if err != nil {
		m.handleError(c, errors.New("invalid credentials"))
		return
	}

	c.Set("user", u)
	c.Next()
}

func (m AuthenticationMiddleware) handleError(c *gin.Context, e error) {
	// Trigger username/password prompt
	c.Header("WWW-Authenticate", "Basic realm=\"DHIS2\"")
	_ = c.AbortWithError(http.StatusUnauthorized, e)
}

func (m AuthenticationMiddleware) TokenAuthentication(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	authorizationHeader = strings.TrimPrefix(authorizationHeader, "Bearer ")

	u, err := m.tokenService.ValidateAccessToken(authorizationHeader)
	if err != nil {
		_ = c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// Extra precaution to ensure that no errors has occurred, and it's safe to call c.Next()
	if len(c.Errors.Errors()) > 0 {
		c.Abort()
		return
	} else {
		c.Set("user", u)
		c.Next()
	}
}
