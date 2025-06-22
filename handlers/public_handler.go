package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Public endpoint of Minicloud REST API", "ping": "pong"})
}
