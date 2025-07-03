// To define / register the REST routes and endpoints for 'account'
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/handlers"
)

func RegisterAccountRoutes(server *gin.Engine) {

	accountsGroup := server.Group("/accounts")
	accountsGroup.GET("", handlers.GetAccounts)
	accountsGroup.POST("", handlers.RegisterAccount)
	accountsGroup.POST("/login", handlers.LoginAccount)
}
