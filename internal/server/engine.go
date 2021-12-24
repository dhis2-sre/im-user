package server

import (
	"github.com/dhis2-sre/im-users/internal/di"
	"github.com/dhis2-sre/im-users/internal/middleware"
	"github.com/dhis2-sre/im-users/pkg/config"
	"github.com/dhis2-sre/im-users/pkg/group"
	"github.com/dhis2-sre/im-users/pkg/health"
	"github.com/dhis2-sre/im-users/pkg/model"
	"github.com/dhis2-sre/im-users/pkg/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	redocMiddleware "github.com/go-openapi/runtime/middleware"
	"log"
	"strings"
)

func GetEngine(environment di.Environment) *gin.Engine {
	basePath := environment.Config.BasePath

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.ErrorHandler())

	router := r.Group(basePath)

	redoc(router, basePath)

	router.GET("/health", health.Health)

	router.GET("/jwks", environment.TokenHandler.Jwks)

	router.POST("/signup", environment.UserHandler.Signup)
	router.POST("/refresh", environment.UserHandler.RefreshToken)

	basicAuthenticationRouter := router.Group("")
	basicAuthenticationRouter.Use(environment.AuthenticationMiddleware.BasicAuthentication)
	basicAuthenticationRouter.POST("/signin", environment.UserHandler.SignIn)

	tokenAuthenticationRouter := router.Group("")
	tokenAuthenticationRouter.Use(environment.AuthenticationMiddleware.TokenAuthentication)
	tokenAuthenticationRouter.GET("/me", environment.UserHandler.Me)
	tokenAuthenticationRouter.GET("/signout", environment.UserHandler.SignOut)
	tokenAuthenticationRouter.GET("/groups-name-to-id/:name", environment.GroupHandler.NameToId)
	tokenAuthenticationRouter.GET("/findbyid/:id", environment.UserHandler.FindById)
	tokenAuthenticationRouter.GET("/groups/:id", environment.GroupHandler.FindById)

	administratorRestrictedRouter := tokenAuthenticationRouter.Group("")
	administratorRestrictedRouter.Use(environment.AuthorizationMiddleware.RequireAdministrator)
	administratorRestrictedRouter.POST("/groups", environment.GroupHandler.Create)
	administratorRestrictedRouter.POST("/groups/:groupId/users/:userId", environment.GroupHandler.AddUserToGroup)
	administratorRestrictedRouter.POST("/groups/:groupId/cluster-configuration", environment.GroupHandler.AddClusterConfiguration)

	createAdminUser(environment.Config, environment.UserService, environment.GroupService)

	createGroups(environment.Config, environment.GroupService)

	return r
}

func redoc(router *gin.RouterGroup, basePath string) {
	router.StaticFile("/swagger.yaml", "./swagger/swagger.yaml")

	redocOpts := redocMiddleware.RedocOpts{
		BasePath: basePath,
		SpecURL:  basePath + "/swagger.yaml",
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
			if strings.HasPrefix(err.Error(), "ERROR: duplicate key value violates unique constraint \"groups_name_key\" (SQLSTATE 23505)") {
				log.Println("Group exists:", g.Name)
			} else {
				log.Fatalln(err)
			}
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

	err = groupService.AddUser(g.ID, u.ID)
	if err != nil {
		log.Fatalf("Failed to add user to admin group: %s", err)
	}
}
