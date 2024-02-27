package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func GetPage(context *gin.Context) {
	data, err := os.ReadFile("src/static/page.html")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "houve um erro inesperado"})
	}

	context.Data(http.StatusOK, "text/html; charset=utf-8", data)
}
