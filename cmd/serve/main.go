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
	"expvar"
	"fmt"
	"net/http"
	"net/http/pprof"
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
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("error getting config: %v", err)
	}

	client := storage.NewRedis(cfg)
	repository := token.NewRepository(client)
	tokenSvc, err := token.NewService(cfg, repository)
	if err != nil {
		return err
	}
	tokenHandler, err := token.NewHandler(cfg)
	if err != nil {
		return err
	}

	db, err := storage.NewDatabase(cfg)
	if err != nil {
		return err
	}
	usrRepository := user.NewRepository(db)
	usrSvc := user.NewService(usrRepository)
	usrHandler := user.NewHandler(cfg, usrSvc, tokenSvc)

	groupRepository := group.NewRepository(db)
	groupSvc := group.NewService(groupRepository, usrSvc)
	groupHandler := group.NewHandler(groupSvc)

	authenticationMiddleware := middleware.NewAuthentication(usrSvc, tokenSvc)
	authorizationMiddleware := middleware.NewAuthorization(usrSvc)

	r, err := server.GetEngine(cfg, tokenHandler, usrHandler, groupHandler, authenticationMiddleware, authorizationMiddleware, usrSvc, groupSvc)
	if err != nil {
		return err
	}

	addr := ":4000"
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
		mux.Handle("/debug/vars", expvar.Handler())
		if err := http.ListenAndServe(addr, mux); err != nil {
			fmt.Printf("error serving debug endpoints: %v\n", err)
		}
	}()

	return r.Run()
}
