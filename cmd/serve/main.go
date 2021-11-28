package main

import (
	"github.com/dhis2-sre/im-users/pgk/config"
	"github.com/dhis2-sre/im-users/pgk/database"
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

	router := r.Group(c.BasePath)
	router.GET("/health", handler.Health)
	router.POST("/signup", errorHandler(handler.Signup))

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func errorHandler(fn func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := fn(c)
		if err == nil {
			return
		}
		c.JSON(helper.ToHttpStatusCode(err), err.Error())
		c.Abort()
	}
}
