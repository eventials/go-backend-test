package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"transmissions-pub/database"
	"transmissions-pub/src/models"
)

func CreateTransmission(context *gin.Context) {
	var transmission *models.Transmission

	err := context.BindJSON(&transmission)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "houve um erro inesperado"})
		return
	}

	database.DB.Create(transmission)

	context.JSON(http.StatusOK, transmission.ID)
}

func GetTransmissions(context *gin.Context) {
	var transmissions []*models.Transmission

	database.DB.Find(&transmissions)

	context.JSON(http.StatusOK, transmissions)
}

func GetNextTransmission(context *gin.Context) {
	var transmission *models.Transmission

	database.DB.Order("date").Where("date >= ?", time.Now()).First(&transmission)

	context.JSON(http.StatusOK, transmission)
}
