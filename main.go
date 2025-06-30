// @title MiniCloud REST API
// @version 1.0
// @description Simulated VM management API built with Go and Docker
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/db"
	"github.com/odeeka/go-minicloud-rest-api/routes"

	_ "github.com/odeeka/go-minicloud-rest-api/docs" // Generated docs by Swagger init
	swaggerFiles "github.com/swaggo/files"           // Embedded Swagger UI files
	ginSwagger "github.com/swaggo/gin-swagger"       // Swagger endpoint handler
)

func main() {
	fmt.Println("MiniCloud Rest API...")

	db.InitDB()

	server := gin.Default()

	// Swagger endpoint
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterAccountRoutes(server)
	routes.RegisterVmsRoutes(server)
	routes.RegisterStoragesRoutes(server)
	routes.RegisterPublicRoutes(server)

	server.Run(":8080")
}
