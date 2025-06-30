package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/odeeka/go-minicloud-rest-api/models"
	"github.com/odeeka/go-minicloud-rest-api/utils"
)

// GetAccounts godoc
// @Summary List all accounts
// @Description Retrieves all registered user accounts
// @Tags account
// @Produce json
// @Success 200 {array} models.Account
// @Failure 500 {object} map[string]string
// @Router /account/all [get]
func GetAccounts(context *gin.Context) {
	accounts, err := models.GetAllAccount()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch accounts.", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, accounts)
}

// RegisterAccount godoc
// @Summary Register a new account
// @Description Create a new user account
// @Tags account
// @Accept json
// @Produce json
// @Param account body models.Account true "Account registration payload"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /accounts/register [post]
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

// LoginAccount godoc
// @Summary Authenticate an account
// @Description Login with username and password to receive JWT token
// @Tags account
// @Accept json
// @Produce json
// @Param credentials body models.Account true "Account login credentials"
// @Success 200 {object} map[string]interface{} "Returns JWT token and account info"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /accounts/login [post]
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
