// To define / register the REST routes and endpoints
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/handlers"
)

func RegisterAccountRoutes(server *gin.Engine) {

	accountsGroup := server.Group("/accounts")
	accountsGroup.GET("", handlers.GetAccounts)
	//accountsGroup.GET("/:id", handlers.GetAccountByUsername)
	accountsGroup.POST("", handlers.RegisterAccount)
	accountsGroup.POST("/login", handlers.LoginAccount)
	//vmsGroup.DELETE("/:id", handlers.DeleteAccount)
	//vmsGroup.PUT("/:id", handlers.UpdateAccount)
}
