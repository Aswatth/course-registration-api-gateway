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

func (obj *StudentProfileController) CreateStudent(context *gin.Context) {
	obj.service.CreateStudentProfile(context)
}

func (obj *StudentProfileController) UpdateStudent(context *gin.Context) {
	obj.service.UpdateStudentProfile(context)
}

func (obj *StudentProfileController) DeleteStudent(context *gin.Context) {
	obj.service.DeleteStudentProfile(context)
}

func (obj *StudentProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	profile_routes := rg.Group("")

	profile_routes.Use(middlewares.ValidateAuthorization([]string{"ADMIN"})).POST("/students", obj.CreateStudent)
	profile_routes.Use(middlewares.ValidateAuthorization([]string{"ADMIN"})).DELETE("/students/:email_id", obj.DeleteStudent)
	profile_routes.Use(middlewares.ValidateAuthorization([]string{"ADMIN"})).PUT("/students/:email_id", obj.UpdateStudent)
}
