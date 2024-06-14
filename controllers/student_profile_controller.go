package controllers

import (
	"course-registration-system/api-gateway/middlewares"
	"course-registration-system/api-gateway/services"

	"github.com/gin-gonic/gin"
)

type StudentProfileController struct {
	service services.StudentProfileService
}

func (obj *StudentProfileController) Init(service services.StudentProfileService) {
	obj.service = service
}

func (obj *StudentProfileController) UpdateStudentPassword(context *gin.Context) {
	obj.service.UpdateStudentPassword(context)
}

func (obj *StudentProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	student_routes := rg.Group("/students")

	student_routes.Use(middlewares.ValidateAuthorization([]string{"STUDENT"}))

	student_routes.PUT("/password/:email_id", obj.UpdateStudentPassword)
}
