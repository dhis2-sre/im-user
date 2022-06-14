// Package classification Instance Manager User Service.
//
// User Service as part of the Instance Manager environment
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//    Version: 0.1.0
//    License: TODO
//    Contact: <info@dhis2.org> https://github.com/dhis2-sre/im-user
//
//    Consumes:
//      - application/json
//      - multipart/form-data
//
//    SecurityDefinitions:
//      basicAuth:
//        type: basic
//      oauth2:
//        type: oauth2
//        tokenUrl: /tokens
//        refreshUrl: /refresh
//        flow: password
// swagger:meta
package main

import (
	"fmt"
	"os"

	"github.com/dhis2-sre/im-user/internal/middleware"
	"github.com/dhis2-sre/im-user/internal/server"
	"github.com/dhis2-sre/im-user/pkg/config"
	"github.com/dhis2-sre/im-user/pkg/group"
	"github.com/dhis2-sre/im-user/pkg/storage"
	"github.com/dhis2-sre/im-user/pkg/token"
	"github.com/dhis2-sre/im-user/pkg/user"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err) // nolint:errcheck
		os.Exit(1)
	}
}

func run() error {
	cfg := config.New()

	client := storage.NewRedis(cfg)
	repository := token.NewRepository(client)
	tokenSvc := token.NewService(cfg, repository)
	tokenHandler := token.NewHandler(cfg)

	db, err := storage.NewDatabase(cfg)
	if err != nil {
		return err
	}
	usrRepository := user.NewRepository(db)
	usrSvc := user.NewService(usrRepository)
	usrHandler := user.NewHandler(cfg, usrSvc, tokenSvc)

	groupRepository := group.NewRepository(db)
	groupSvc := group.NewService(groupRepository, usrRepository)
	groupHandler := group.NewHandler(groupSvc, usrSvc)

	authenticationMiddleware := middleware.NewAuthentication(usrSvc, tokenSvc)
	authorizationMiddleware := middleware.NewAuthorization(usrSvc)

	r := server.GetEngine(cfg, tokenHandler, usrHandler, groupHandler, authenticationMiddleware, authorizationMiddleware, usrSvc, groupSvc)
	return r.Run()
}
