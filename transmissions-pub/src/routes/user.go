package routes

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"transmissions-pub/database"
	"transmissions-pub/lib"
	"transmissions-pub/src/models"
)

func CreateUser(context *gin.Context) {
	var authUser *models.AuthUser

	if err := context.ShouldBindJSON(&authUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "houve um erro inesperado"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authUser.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	(*authUser).Password = string(passwordHash)

	database.DB.Create(&authUser)

	context.JSON(http.StatusOK, authUser.ID)
}

func CurrentUser(context *gin.Context) {

	userId, err := lib.ExtractTokenID(context)

	authUser := models.AuthUser{}

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Find(&authUser, userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "success", "data": authUser})
}
