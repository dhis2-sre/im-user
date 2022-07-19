package middleware

import (
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
		if c.Writer.Status() != http.StatusOK {
			_, _ = c.Writer.WriteString(err.Error())
			return
		}

		if errdef.IsDuplicated(err) {
			_ = c.AbortWithError(http.StatusBadRequest, err)
		} else if errdef.IsNotFound(err) {
			_ = c.AbortWithError(http.StatusNotFound, err)
		} else if errdef.IsUnauthorized(err) {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
		} else {
			id := uuid.New()
			log.Printf("unhandled error: %v, log id: %s\n", err, id)
			_ = c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("something went wrong. We'll look into it if you send us the id %q :)", id))
		}
	}
}
