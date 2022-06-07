package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}
		if c.Writer.Status() != http.StatusOK {
			c.Writer.WriteString(err.Error()) // nolint:errcheck
			return
		}

		c.String(http.StatusInternalServerError, err.Error())
	}
}
