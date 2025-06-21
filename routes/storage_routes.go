// To define / register the REST routes and endpoints
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/handlers"
	"github.com/odeeka/go-minicloud-rest-api/middlewares"
)

func RegisterStoragesRoutes(server *gin.Engine) {

	// without authentication
	// storagesGroup := server.Group("/storages")
	// storagesGroup.GET("", handlers.ListStorages)
	// storagesGroup.POST("", handlers.CreateStorage)
	// storagesGroup.GET("/:id", handlers.GetStorage)
	// storagesGroup.DELETE("/:id", handlers.DeleteStorage)
	// storagesGroup.PUT("/:id", handlers.UpdateStorageSize)
	// storagesGroup.POST("/:id/attach/:vmid", handlers.AttachStorageToVM)
	// storagesGroup.POST("/:id/detach/:vmid", handlers.DetachStorageFromVM)

	// with authentication through middleware
	authStoragesGroup := server.Group("/storages")
	authStoragesGroup.Use(middlewares.Authenticate)
	authStoragesGroup.GET("", handlers.ListStorages)
	authStoragesGroup.POST("", handlers.CreateStorage)
	authStoragesGroup.GET("/:id", handlers.GetStorage)
	authStoragesGroup.DELETE("/:id", handlers.DeleteStorage)
	authStoragesGroup.PUT("/:id", handlers.UpdateStorageSize)
	authStoragesGroup.POST("/:id/attach/:vmid", handlers.AttachStorageToVM)
	authStoragesGroup.POST("/:id/detach/:vmid", handlers.DetachStorageFromVM)
}
