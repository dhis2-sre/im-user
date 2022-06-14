package server

import (
	"log"

	"github.com/dhis2-sre/im-user/internal/middleware"
	"github.com/dhis2-sre/im-user/pkg/config"
	"github.com/dhis2-sre/im-user/pkg/group"
	"github.com/dhis2-sre/im-user/pkg/health"
	"github.com/dhis2-sre/im-user/pkg/model"
	"github.com/dhis2-sre/im-user/pkg/token"
	"github.com/dhis2-sre/im-user/pkg/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	redocMiddleware "github.com/go-openapi/runtime/middleware"
)

func GetEngine(c config.Config, tokenHandler token.Handler, usrHandler user.Handler, groupHandler group.Handler, authenticationMiddleware middleware.AuthenticationMiddleware, authorizationMiddleware middleware.AuthorizationMiddleware, usrSvc user.Service, groupSvc group.Service) *gin.Engine {
	basePath := c.BasePath

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.ErrorHandler())

	router := r.Group(basePath)

	redoc(router, basePath)

	router.GET("/health", health.Health)

	router.GET("/jwks", tokenHandler.Jwks)

	router.POST("/users", usrHandler.SignUp)
	router.POST("/refresh", usrHandler.RefreshToken)

	basicAuthenticationRouter := router.Group("")
	basicAuthenticationRouter.Use(authenticationMiddleware.BasicAuthentication)
	basicAuthenticationRouter.POST("/tokens", usrHandler.SignIn)

	tokenAuthenticationRouter := router.Group("")
	tokenAuthenticationRouter.Use(authenticationMiddleware.TokenAuthentication)
	tokenAuthenticationRouter.GET("/me", usrHandler.Me)
	tokenAuthenticationRouter.DELETE("/users", usrHandler.SignOut)
	tokenAuthenticationRouter.GET("/users/:id", usrHandler.FindById)

	tokenAuthenticationRouter.GET("/groups/:name", groupHandler.Find)

	administratorRestrictedRouter := tokenAuthenticationRouter.Group("")
	administratorRestrictedRouter.Use(authorizationMiddleware.RequireAdministrator)
	administratorRestrictedRouter.POST("/groups", groupHandler.Create)
	administratorRestrictedRouter.POST("/groups/:groupName/users/:userId", groupHandler.AddUserToGroup)
	administratorRestrictedRouter.POST("/groups/:groupName/cluster-configuration", groupHandler.AddClusterConfiguration)

	createAdminUser(c, usrSvc, groupSvc)
	createGroups(c, groupSvc)
	createServiceUsers(c, usrSvc, groupSvc)

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

func createGroups(config config.Config, groupSvc group.Service) {
	log.Println("Creating groups...")
	groups := config.Groups
	for _, g := range groups {
		newGroup, err := groupSvc.FindOrCreate(g.Name, g.Hostname)
		if err != nil {
			log.Fatalln(err)
		}
		if newGroup != nil {
			log.Println("Created:", newGroup.Name)
		}
	}
}

func createAdminUser(config config.Config, usrSvc user.Service, groupSvc group.Service) {
	adminUserEmail := config.AdminUser.Email
	adminUserPassword := config.AdminUser.Password

	u, err := usrSvc.FindOrCreate(adminUserEmail, adminUserPassword)
	if err != nil {
		log.Fatalln(err)
	}

	g, err := groupSvc.FindOrCreate(model.AdministratorGroupName, "")
	if err != nil {
		log.Fatalf("Failed to create admin group: %s", err)
	}

	err = groupSvc.AddUser(g.Name, u.ID)
	if err != nil {
		log.Fatalf("Failed to add user to admin group: %s", err)
	}
}

func createServiceUsers(config config.Config, usrSvc user.Service, groupSvc group.Service) {
	log.Println("Creating service users...")
	g, err := groupSvc.FindOrCreate(model.AdministratorGroupName, "")
	if err != nil {
		log.Fatalf("Failed to create admin group: %s", err)
	}

	for _, serviceUser := range config.ServiceUsers {
		email := serviceUser.Email
		password := serviceUser.Password

		u, err := usrSvc.FindOrCreate(email, password)
		if err != nil {
			log.Fatalln(err)
		}

		err = groupSvc.AddUser(g.Name, u.ID)
		if err != nil {
			log.Fatalf("Failed to add user to admin group: %s", err)
		}

		log.Println("Created:", serviceUser.Email)
	}
}
