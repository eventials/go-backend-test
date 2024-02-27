package main

import (
	"github.com/gin-gonic/gin"
	"transmissions-pub/database"
	"transmissions-pub/src/middlewares"
	"transmissions-pub/src/models"
	"transmissions-pub/src/routes"
)

func main() {
	database.Connect()

	err := database.DB.AutoMigrate(&models.AuthUser{}, &models.Transmission{})
	if err != nil {
		return
	}

	r := gin.Default()

	r.GET("/", routes.GetPage)
	r.GET("/wb", routes.WebSocket)

	api := r.Group("/api")
	{
		login := api.Group("/login")
		login.POST("", routes.Login)

		user := api.Group("/user")
		user.GET("/current", routes.CurrentUser)
		user.POST("", routes.CreateUser)

		transmision := api.Group("/transmission", middlewares.JwtAuthMiddleware())
		transmision.GET("", routes.GetTransmissions)
		transmision.POST("", routes.CreateTransmission)
		transmision.GET("/next", routes.GetNextTransmission)
	}

	err = r.Run("localhost:8080")

	if err != nil {
		return
	}
}
