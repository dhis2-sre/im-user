package server

import (
	"fmt"
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

type userService interface {
	FindOrCreate(email string, password string) (*model.User, error)
}

type groupService interface {
	AddUser(groupName string, userId uint) error
	FindOrCreate(name string, hostname string) (*model.Group, error)
}

func GetEngine(c config.Config, tokenHandler token.Handler, usrHandler user.Handler, groupHandler group.Handler, authenticationMiddleware middleware.AuthenticationMiddleware, authorizationMiddleware middleware.AuthorizationMiddleware, usrSvc userService, groupSvc groupService) (*gin.Engine, error) {
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
	administratorRestrictedRouter.POST("/groups/:group/users/:userId", groupHandler.AddUserToGroup)
	administratorRestrictedRouter.POST("/groups/:group/cluster-configuration", groupHandler.AddClusterConfiguration)

	err := createAdminUser(c, usrSvc, groupSvc)
	if err != nil {
		return nil, err
	}
	err = createGroups(c, groupSvc)
	if err != nil {
		return nil, err
	}
	err = createServiceUsers(c, usrSvc, groupSvc)
	if err != nil {
		return nil, err
	}

	return r, nil
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

func createGroups(config config.Config, groupSvc groupService) error {
	log.Println("Creating groups...")
	groups := config.Groups
	for _, g := range groups {
		newGroup, err := groupSvc.FindOrCreate(g.Name, g.Hostname)
		if err != nil {
			return fmt.Errorf("error creating group: %v", err)
		}
		if newGroup != nil {
			log.Println("Created:", newGroup.Name)
		}
	}

	return nil
}

func createAdminUser(config config.Config, usrSvc userService, groupSvc groupService) error {
	adminUserEmail := config.AdminUser.Email
	adminUserPassword := config.AdminUser.Password

	u, err := usrSvc.FindOrCreate(adminUserEmail, adminUserPassword)
	if err != nil {
		return fmt.Errorf("error creating admin user: %v", err)
	}

	g, err := groupSvc.FindOrCreate(model.AdministratorGroupName, "")
	if err != nil {
		return fmt.Errorf("error creating admin group: %v", err)
	}

	err = groupSvc.AddUser(g.Name, u.ID)
	if err != nil {
		return fmt.Errorf("error adding admin user to admin group: %v", err)
	}

	return nil
}

func createServiceUsers(config config.Config, usrSvc userService, groupSvc groupService) error {
	log.Println("Creating service users...")
	g, err := groupSvc.FindOrCreate(model.AdministratorGroupName, "")
	if err != nil {
		return fmt.Errorf("error creating admin group: %v", err)
	}

	for _, serviceUser := range config.ServiceUsers {
		email := serviceUser.Email
		password := serviceUser.Password

		u, err := usrSvc.FindOrCreate(email, password)
		if err != nil {
			return fmt.Errorf("error creating service user: %v", err)
		}

		err = groupSvc.AddUser(g.Name, u.ID)
		if err != nil {
			return fmt.Errorf("error adding service user to admin group: %v", err)
		}

		log.Println("Created:", serviceUser.Email)
	}

	return nil
}
