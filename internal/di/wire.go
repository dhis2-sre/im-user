//+build wireinject

package di

import (
	"github.com/dhis2-sre/im-user/internal/middleware"
	"github.com/dhis2-sre/im-user/pkg/config"
	"github.com/dhis2-sre/im-user/pkg/group"
	"github.com/dhis2-sre/im-user/pkg/storage"
	"github.com/dhis2-sre/im-user/pkg/token"
	"github.com/dhis2-sre/im-user/pkg/user"
	"github.com/google/wire"
	"gorm.io/gorm"
	"log"
)

type Environment struct {
	Config                   config.Config
	TokenService             token.Service
	TokenHandler             token.Handler
	UserService              user.Service
	UserHandler              user.Handler
	GroupService             group.Service
	GroupHandler             group.Handler
	AuthenticationMiddleware middleware.AuthenticationMiddleware
	AuthorizationMiddleware  middleware.AuthorizationMiddleware
}

func ProvideEnvironment(
	config config.Config,
	tokenService token.Service,
	tokenHandler token.Handler,
	userService user.Service,
	userHandler user.Handler,
	groupService group.Service,
	groupHandler group.Handler,
	authenticationMiddleware middleware.AuthenticationMiddleware,
	authorizationMiddleware middleware.AuthorizationMiddleware,
) Environment {
	return Environment{
		config,
		tokenService,
		tokenHandler,
		userService,
		userHandler,
		groupService,
		groupHandler,
		authenticationMiddleware,
		authorizationMiddleware,
	}
}

func GetEnvironment() Environment {
	wire.Build(
		config.ProvideConfig,

		provideDatabase,
		storage.ProvideRedis,

		token.ProvideTokenRepository,
		token.ProvideTokenService,

		user.ProvideRepository,
		user.ProvideService,
		user.ProvideHandler,

		group.ProvideRepository,
		group.ProvideService,
		group.ProvideHandler,

		token.ProvideHandler,

		middleware.ProvideAuthentication,
		middleware.ProvideAuthorization,

		ProvideEnvironment,
	)
	return Environment{}
}

func provideDatabase(c config.Config) *gorm.DB {
	database, err := storage.ProvideDatabase(c)
	if err != nil {
		log.Fatalln(err)
	}
	return database
}
