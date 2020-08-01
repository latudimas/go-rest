package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"example.com/go-rest/connectors"
	"example.com/go-rest/routes"
)

func main() {
	// Initiate Database
	connectors.ConnectDatabase()

	// Init Router
	router := gin.Default()

	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}
