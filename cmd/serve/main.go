package main

import (
	"github.com/dhis2-sre/im-users/internal/middleware"
	"github.com/dhis2-sre/im-users/pgk/config"
	"github.com/dhis2-sre/im-users/pgk/health"
	"github.com/dhis2-sre/im-users/pgk/storage"
	"github.com/dhis2-sre/im-users/pgk/token"
	"github.com/dhis2-sre/im-users/pgk/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	c := config.ProvideConfig()

	db, err := storage.ProvideDatabase(c)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.ProvideRepository(db)
	userService := user.ProvideService(userRepository)

	authenticationMiddleware := middleware.ProvideAuthenticationMiddleware(userService)

	redis := storage.ProvideRedis(c)
	tokenRepository := token.ProvideTokenRepository(redis)
	tokenService := token.ProvideTokenService(c, tokenRepository)

	userHandler := user.ProvideHandler(userService, tokenService)

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.ErrorHandler())

	router := r.Group(c.BasePath)
	router.GET("/health", health.Health)
	router.POST("/signup", userHandler.Signup)
	router.POST("/signin", authenticationMiddleware.BasicAuthentication, userHandler.SignIn)
	router.POST("/refresh", userHandler.RefreshToken)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
