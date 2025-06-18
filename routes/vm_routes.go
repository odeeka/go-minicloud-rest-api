// To define / register the REST routes and endpoints
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/handlers"
)

func RegisterVmsRoutes(server *gin.Engine) {

	vmsGroup := server.Group("/vms")
	vmsGroup.GET("", handlers.ListVMs)
	vmsGroup.GET("/:id", handlers.GetVM)
	vmsGroup.POST("", handlers.CreateVM)
	vmsGroup.DELETE("/:id", handlers.DeleteVM)
	vmsGroup.PUT("/:id", handlers.UpdateVM)
}
