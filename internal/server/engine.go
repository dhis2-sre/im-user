package server

import (
	"github.com/dhis2-sre/im-users/internal/di"
	"github.com/dhis2-sre/im-users/internal/middleware"
	"github.com/dhis2-sre/im-users/pgk/config"
	"github.com/dhis2-sre/im-users/pgk/group"
	"github.com/dhis2-sre/im-users/pgk/health"
	"github.com/dhis2-sre/im-users/pgk/model"
	"github.com/dhis2-sre/im-users/pgk/user"
	"github.com/dhis2-sre/im-users/swagger/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

func GetEngine(environment di.Environment) *gin.Engine {
	basePath := environment.Config.BasePath

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(middleware.ErrorHandler())

	router := r.Group(basePath)
	router.GET("/health", health.Health)
	router.GET("/jwks", environment.TokenHandler.Jwks)
	router.POST("/signup", environment.UserHandler.Signup)
	router.POST("/refresh", environment.UserHandler.RefreshToken)
	router.GET("/findbyid/:id", environment.UserHandler.FindById)
	router.POST("/signin", environment.AuthenticationMiddleware.BasicAuthentication, environment.UserHandler.SignIn)

	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", environment.AuthenticationMiddleware.BasicAuthentication, ginSwagger.WrapHandler(swaggerFiles.Handler))

	tokenAuthenticationRouter := router.Group("")
	tokenAuthenticationRouter.Use(environment.AuthenticationMiddleware.TokenAuthentication)
	tokenAuthenticationRouter.GET("/me", environment.UserHandler.Me)
	tokenAuthenticationRouter.GET("/signout", environment.UserHandler.SignOut)

	administratorRestrictedRouter := tokenAuthenticationRouter.Group("")
	administratorRestrictedRouter.Use(environment.AuthorizationMiddleware.RequireAdministrator)
	administratorRestrictedRouter.POST("/groups", environment.GroupHandler.Create)
	administratorRestrictedRouter.POST("/groups/:groupId/users/:userId", environment.GroupHandler.AddUserToGroup)
	administratorRestrictedRouter.POST("/groups/:groupId/cluster-configuration", environment.GroupHandler.AddClusterConfiguration)

	createAdminUser(environment.Config, environment.UserService, environment.GroupService)

	return r
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
