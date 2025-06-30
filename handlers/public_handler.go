package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary Show public API status
// @Description Responds with a simple ping response to verify that the public endpoint is reachable.
// @Tags public
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /ping [get]
// @Router /public/ping [get]
func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Public endpoint of Minicloud REST API", "ping": "pong"})
}
