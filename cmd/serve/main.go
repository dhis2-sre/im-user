package main

import (
	"github.com/dhis2-sre/im-users/internal/di"
	"github.com/dhis2-sre/im-users/internal/server"
	"log"
)

// @title Instance Manager User Service
// @version 0.1.0
func main() {
	environment := di.GetEnvironment()
	r := server.GetEngine(environment)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
