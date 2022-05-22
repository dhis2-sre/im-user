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
	"log"

	"github.com/dhis2-sre/im-user/internal/di"
	"github.com/dhis2-sre/im-user/internal/server"
)

func main() {
	environment := di.GetEnvironment()
	r := server.GetEngine(environment)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
