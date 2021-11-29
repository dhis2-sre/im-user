package main

import (
	"github.com/dhis2-sre/im-users/pgk/config"
	"github.com/dhis2-sre/im-users/pgk/database"
	"github.com/dhis2-sre/im-users/pgk/health"
	"github.com/dhis2-sre/im-users/pgk/helper"
	"github.com/dhis2-sre/im-users/pgk/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	c := config.ProvideConfig()

	db, err := database.ProvideDatabase(c)
	if err != nil {
		log.Fatal(err)
	}

	repository := user.ProvideRepository(db)
	service := user.ProvideService(repository)
	handler := user.ProvideHandler(service)

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(errorHandler())

	router := r.Group(c.BasePath)
	router.GET("/health", health.Health)
	router.POST("/signup", handler.Signup)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func errorHandler() gin.HandlerFunc {
	return errorHandlerT(gin.ErrorTypeAny)
}

func errorHandlerT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) > 0 {
			// TODO: Handle more than one error
			err := detectedErrors[0].Err
			c.JSON(helper.ToHttpStatusCode(err), err.Error())
			c.Abort()
			return
		}
	}
}
