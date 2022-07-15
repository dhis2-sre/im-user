package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/dhis2-sre/im-user/internal/errdef"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		if errors.As(err, &errdef.Duplicate{}) {
			_ = c.AbortWithError(http.StatusBadRequest, err)
		} else if errdef.IsNotFound(err) {
			_ = c.AbortWithError(http.StatusNotFound, err)
		} else if errdef.IsUnauthorized(err) {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
		} else {
			id := uuid.New()
			log.Printf("unhandled error: %v, log id: %s", err, id)
			_ = c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("unhandled error, log id: %s", id.String()))
		}
	}
}
