package main

import (
	"github.com/dhis2-sre/im-users/internal/middleware"
	"github.com/dhis2-sre/im-users/pgk/config"
	"github.com/dhis2-sre/im-users/pgk/group"
	"github.com/dhis2-sre/im-users/pgk/health"
	"github.com/dhis2-sre/im-users/pgk/model"
	"github.com/dhis2-sre/im-users/pgk/storage"
	"github.com/dhis2-sre/im-users/pgk/token"
	"github.com/dhis2-sre/im-users/pgk/user"
	"github.com/dhis2-sre/im-users/swagger/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

// @title Instance Manager User Service
// @version 0.1.0
func main() {
	c := config.ProvideConfig()

	db, err := storage.ProvideDatabase(c)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.ProvideRepository(db)
	userService := user.ProvideService(userRepository)

	redis := storage.ProvideRedis(c)
	tokenRepository := token.ProvideTokenRepository(redis)
	tokenService := token.ProvideTokenService(c, tokenRepository)

	userHandler := user.ProvideHandler(c, userService, tokenService)

	groupRepository := group.ProvideRepository(db)
	groupService := group.ProvideService(groupRepository, userRepository)
	groupHandler := group.ProvideHandler(groupService, userService)

	createAdminUser(c, userService, groupService)

	authenticationMiddleware := middleware.ProvideAuthentication(userService, tokenService)
	authorizationMiddleware := middleware.ProvideAuthorization(userService)

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.ErrorHandler())

	router := r.Group(c.BasePath)
	router.GET("/health", health.Health)
	router.POST("/signup", userHandler.Signup)
	router.POST("/refresh", userHandler.RefreshToken)
	router.GET("/findbyid/:id", userHandler.FindById)
	router.POST("/signin", authenticationMiddleware.BasicAuthentication, userHandler.SignIn)

	docs.SwaggerInfo.BasePath = c.BasePath
	router.GET("/swagger/*any", authenticationMiddleware.BasicAuthentication, ginSwagger.WrapHandler(swaggerFiles.Handler))

	tokenAuthenticationRouter := router.Group("")
	tokenAuthenticationRouter.Use(authenticationMiddleware.TokenAuthentication)
	tokenAuthenticationRouter.GET("/me", userHandler.Me)
	tokenAuthenticationRouter.GET("/signout", userHandler.SignOut)

	administratorRestrictedRouter := tokenAuthenticationRouter.Group("")
	administratorRestrictedRouter.Use(authorizationMiddleware.RequireAdministrator)
	administratorRestrictedRouter.POST("/groups", groupHandler.Create)
	administratorRestrictedRouter.POST("/groups/:groupId/users/:userId", groupHandler.AddUserToGroup)
	administratorRestrictedRouter.POST("/groups/:groupId/cluster-configuration", groupHandler.AddClusterConfiguration)

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func createAdminUser(config config.Config, userService user.Service, groupService group.Service) {
	adminUserEmail := config.AdminUser.Email
	adminUserPassword := config.AdminUser.Password

	u, _ := userService.FindByEmail(adminUserEmail)
	if u != nil && u.ID > 0 {
		log.Println("Admin user exists")
		return
	}

	adminUser, err := userService.Signup(adminUserEmail, adminUserPassword)
	if err != nil {
		log.Fatalf("Failed to create admin user: %s", err)
	}

	g, err := groupService.Create(model.AdministratorGroupName, "")
	if err != nil {
		log.Fatalf("Failed to create admin group: %s", err)
	}

	err = groupService.AddUser(g.ID, adminUser.ID)
	if err != nil {
		log.Fatalf("Failed to add user to admin group: %s", err)
	}
}
