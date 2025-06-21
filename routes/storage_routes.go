// To define / register the REST routes and endpoints
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/handlers"
)

func RegisterStoragesRoutes(server *gin.Engine) {

	storagesGroup := server.Group("/storages")
	storagesGroup.GET("", handlers.ListStorages)
	storagesGroup.POST("", handlers.CreateStorage)
	storagesGroup.GET("/:id", handlers.GetStorage)
	storagesGroup.DELETE("/:id", handlers.DeleteStorage)
	storagesGroup.PUT("/:id", handlers.UpdateStorageSize)
	storagesGroup.POST("/:id/attach/:vmid", handlers.AttachStorageToVM)
	storagesGroup.POST("/:id/detach/:vmid", handlers.DetachStorageFromVM)
}
