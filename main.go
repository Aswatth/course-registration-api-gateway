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
	login_controller.Init(*login_service)

	admin_profile_service := new(services.AdminProfileService)
	admin_profile_service.Init()

	admin_profile_controller := new(controllers.AdminProfileController)
	admin_profile_controller.Init(*admin_profile_service)

	student_profile_service := new(services.StudentProfileService)
	student_profile_service.Init()

	student_profile_controller := new(controllers.StudentProfileController)
	student_profile_controller.Init(*student_profile_service)

	professor_profile_service := new(services.ProfessorProfileService)
	professor_profile_service.Init()

	professor_profile_controller := new(controllers.ProfessorProfileController)
	professor_profile_controller.Init(*professor_profile_service)

	server := gin.Default()

	base_path := server.Group("/api")

	login_controller.RegisterRoutes(base_path)
	admin_profile_controller.RegisterRoutes(base_path)
	student_profile_controller.RegisterRoutes(base_path)
	professor_profile_controller.RegisterRoutes(base_path)

	server.GET("/isAlive", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Yes",
		})
	})

	server.Run(":" + os.Getenv("PORT"))
}
