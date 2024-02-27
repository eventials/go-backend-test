package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"transmissions-pub/lib"
	"transmissions-pub/src/models"
)

func Login(context *gin.Context) {
	var inputUser models.AuthUser

	if err := context.ShouldBindJSON(&inputUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := lib.LoginCheck(inputUser.Username, inputUser.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})

}
