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
//    Contact: <info@dhis2.org> https://github.com/dhis2-sre/im-users
//
//    Consumes:
//    - application/json
//    - multipart/form-data
//
// swagger:meta
package main

import (
	"github.com/dhis2-sre/im-users/internal/di"
	"github.com/dhis2-sre/im-users/internal/server"
	"log"
)

func main() {
	environment := di.GetEnvironment()
	r := server.GetEngine(environment)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
