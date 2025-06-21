// To define / register the REST routes and endpoints
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/handlers"
	"github.com/odeeka/go-minicloud-rest-api/middlewares"
)

func RegisterVmsRoutes(server *gin.Engine) {

	// without authentication
	// vmsGroup := server.Group("/vms")
	// vmsGroup.GET("", handlers.ListVMs)
	// vmsGroup.GET("/:id", handlers.GetVM)
	// vmsGroup.POST("", handlers.CreateVM)
	// vmsGroup.DELETE("/:id", handlers.DeleteVM)
	// vmsGroup.PUT("/:id", handlers.UpdateVM)

	// with authentication through middleware
	authVmsGroup := server.Group("/vms")
	authVmsGroup.Use(middlewares.Authenticate)
	authVmsGroup.GET("", handlers.ListVMs)
	authVmsGroup.GET("/:id", handlers.GetVM)
	authVmsGroup.POST("", handlers.CreateVM)
	authVmsGroup.DELETE("/:id", handlers.DeleteVM)
	authVmsGroup.PUT("/:id", handlers.UpdateVM)
}
