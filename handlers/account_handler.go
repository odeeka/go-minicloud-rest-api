package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/models"
	"github.com/odeeka/go-minicloud-rest-api/utils"
)

func GetAccounts(context *gin.Context) {
	accounts, err := models.GetAllAccount()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch accounts.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, accounts)
}

func RegisterAccount(context *gin.Context) {
	var acc models.Account

	err := context.ShouldBindJSON(&acc)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	err = acc.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save account.", "error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Account created successfully"})
}

func LoginAccount(context *gin.Context) {
	var acc models.Account

	err := context.ShouldBindJSON(&acc)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data.", "error": err.Error()})
		return
	}

	err = acc.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate the account.", "Error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(acc.Username, acc.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate the account", "Error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "Token": token, "Account ID": acc.ID, "Account username": acc.Username})
}
