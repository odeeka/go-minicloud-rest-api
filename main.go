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

	routes.RegisterVmsRoutes(server)

	server.Run(":8080")
}
