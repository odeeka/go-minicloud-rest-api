package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/db"
	"github.com/odeeka/go-minicloud-rest-api/routes"
)

func main() {
	fmt.Println("MiniCloud Rest API...")

	db.InitDB()

	server := gin.Default()

	routes.RegisterAccountRoutes(server)
	routes.RegisterVmsRoutes(server)
	routes.RegisterStoragesRoutes(server)
	routes.RegisterPublicRoutes(server)

	server.Run(":8080")
}
