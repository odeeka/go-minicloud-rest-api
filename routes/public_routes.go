package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/handlers"
)

func RegisterPublicRoutes(server *gin.Engine) {

	publicGroup := server.Group("/public")
	publicGroup.GET("", handlers.Ping)
	publicGroup.GET("/ping", handlers.Ping)
}
