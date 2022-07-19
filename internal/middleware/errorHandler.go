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
			c.String(http.StatusBadRequest, err.Error())
		} else if errdef.IsNotFound(err) {
			c.String(http.StatusNotFound, err.Error())
		} else if errdef.IsUnauthorized(err) {
			c.String(http.StatusUnauthorized, err.Error())
		} else {
			id := uuid.New()
			log.Printf("unhandled error: %q, log id: %s\n", err, id)
			err := fmt.Errorf("something went wrong. We'll look into it if you send us the id %q :)", id)
			c.String(http.StatusInternalServerError, err.Error())
		}
	}
}
