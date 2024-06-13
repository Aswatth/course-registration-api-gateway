package main

import (
	"course-registration-system/api-gateway/controllers"
	"course-registration-system/api-gateway/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	login_service := new(services.LoginService)
	login_controller := new(controllers.LoginController)
	login_controller.Init(login_service)

	server := gin.Default()

	base_path := server.Group("/api")

	login_controller.RegisterRoutes(base_path)

	server.GET("/isAlive", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Yes",
		})
	})

	server.Run(":" + os.Getenv("PORT"))
}
