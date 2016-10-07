package main

import (
	"trom/apigateway"
	"trom/auth"
	"trom/router"
)

func main() {

	router := router.CreateNew()

	educationRoute := router.AddRoute("/educacao")
	educationRoute.AddNode("localhost:8888")
	educationRoute.AddNode("localhost:8888")

	rootRoute := router.AddRoute("/")
	rootRoute.AddNode("localhost:7777")
	rootRoute.AddNode("localhost:7777")

	userAuth := auth.User()
	appAuth := auth.App()

	apigateway := apigateway.New(router)
	apigateway.AddPreHandler(userAuth)
	apigateway.AddPreHandler(appAuth)
	apigateway.Start()
}
