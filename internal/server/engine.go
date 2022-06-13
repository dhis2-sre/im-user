package server

import (
	"log"

	"github.com/dhis2-sre/im-user/internal/di"
	"github.com/dhis2-sre/im-user/internal/middleware"
	"github.com/dhis2-sre/im-user/pkg/config"
	"github.com/dhis2-sre/im-user/pkg/group"
	"github.com/dhis2-sre/im-user/pkg/health"
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/dhis2-sre/im-user/pkg/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	redocMiddleware "github.com/go-openapi/runtime/middleware"
)

func GetEngine(environment di.Environment) *gin.Engine {
	c := environment.Config
	basePath := c.BasePath

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.ErrorHandler())

	router := r.Group(basePath)

	redoc(router, basePath)

	router.GET("/health", health.Health)

	router.GET("/jwks", environment.TokenHandler.Jwks)

	router.POST("/users", environment.UserHandler.SignUp)
	router.POST("/refresh", environment.UserHandler.RefreshToken)

	basicAuthenticationRouter := router.Group("")
	basicAuthenticationRouter.Use(environment.AuthenticationMiddleware.BasicAuthentication)
	basicAuthenticationRouter.POST("/tokens", environment.UserHandler.SignIn)

	tokenAuthenticationRouter := router.Group("")
	tokenAuthenticationRouter.Use(environment.AuthenticationMiddleware.TokenAuthentication)
	tokenAuthenticationRouter.GET("/me", environment.UserHandler.Me)
	tokenAuthenticationRouter.DELETE("/users", environment.UserHandler.SignOut)
	tokenAuthenticationRouter.GET("/users/:id", environment.UserHandler.FindById)

	tokenAuthenticationRouter.GET("/groups/:name", environment.GroupHandler.Find)

	administratorRestrictedRouter := tokenAuthenticationRouter.Group("")
	administratorRestrictedRouter.Use(environment.AuthorizationMiddleware.RequireAdministrator)
	administratorRestrictedRouter.POST("/groups", environment.GroupHandler.Create)
	administratorRestrictedRouter.POST("/groups/:groupName/users/:userId", environment.GroupHandler.AddUserToGroup)
	administratorRestrictedRouter.POST("/groups/:groupName/cluster-configuration", environment.GroupHandler.AddClusterConfiguration)

	groupService := environment.GroupService
	userService := environment.UserService
	createAdminUser(c, userService, groupService)
	createGroups(c, groupService)
	createServiceUsers(c, userService, groupService)

	return r
}

func redoc(router *gin.RouterGroup, basePath string) {
	router.StaticFile("/swagger.yaml", "./swagger/swagger.yaml")

	redocOpts := redocMiddleware.RedocOpts{
		BasePath: basePath,
		SpecURL:  "./swagger.yaml",
	}
	router.GET("/docs", func(c *gin.Context) {
		redocHandler := redocMiddleware.Redoc(redocOpts, nil)
		redocHandler.ServeHTTP(c.Writer, c.Request)
	})
}

func createGroups(config config.Config, groupService group.Service) {
	log.Println("Creating groups...")
	groups := config.Groups
	for _, g := range groups {
		newGroup, err := groupService.FindOrCreate(g.Name, g.Hostname)
		if err != nil {
			log.Fatalln(err)
		}
		if newGroup != nil {
			log.Println("Created:", newGroup.Name)
		}
	}
}

func createAdminUser(config config.Config, userService user.Service, groupService group.Service) {
	adminUserEmail := config.AdminUser.Email
	adminUserPassword := config.AdminUser.Password

	u, err := userService.FindOrCreate(adminUserEmail, adminUserPassword)
	if err != nil {
		log.Fatalln(err)
	}

	g, err := groupService.FindOrCreate(model.AdministratorGroupName, "")
	if err != nil {
		log.Fatalf("Failed to create admin group: %s", err)
	}

	err = groupService.AddUser(g.Name, u.ID)
	if err != nil {
		log.Fatalf("Failed to add user to admin group: %s", err)
	}
}

func createServiceUsers(config config.Config, userService user.Service, groupService group.Service) {
	log.Println("Creating service users...")
	g, err := groupService.FindOrCreate(model.AdministratorGroupName, "")
	if err != nil {
		log.Fatalf("Failed to create admin group: %s", err)
	}

	for _, serviceUser := range config.ServiceUsers {
		email := serviceUser.Email
		password := serviceUser.Password

		u, err := userService.FindOrCreate(email, password)
		if err != nil {
			log.Fatalln(err)
		}

		err = groupService.AddUser(g.Name, u.ID)
		if err != nil {
			log.Fatalf("Failed to add user to admin group: %s", err)
		}

		log.Println("Created:", serviceUser.Email)
	}
}
