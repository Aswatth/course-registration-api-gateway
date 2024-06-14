package controllers

import (
	"course-registration-system/api-gateway/middlewares"
	"course-registration-system/api-gateway/services"

	"github.com/gin-gonic/gin"
)

type ProfessorProfileController struct {
	service services.ProfessorProfileService
}

func (obj *ProfessorProfileController) Init(service services.ProfessorProfileService) {
	obj.service = service
}

func (obj *ProfessorProfileController) UpdateProfessorPassword(context *gin.Context) {
	// obj.service.UpdateStudentPassword(context)
}

func (obj *ProfessorProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	professor_routes := rg.Group("")

	professor_routes.Use(middlewares.ValidateAuthorization([]string{"PROFESSOR"}))

	professor_routes.PUT("/professors/password/:email_id", obj.UpdateProfessorPassword)
}
