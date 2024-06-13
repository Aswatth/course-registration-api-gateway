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

func (obj *ProfessorProfileController) CreateProfessor(context *gin.Context) {
	obj.service.CreateProfessorProfile(context)
}

func (obj *ProfessorProfileController) UpdateProfessor(context *gin.Context) {
	obj.service.UpdateProfessorProfile(context)
}

func (obj *ProfessorProfileController) DeleteProfessor(context *gin.Context) {
	obj.service.DeleteProfessorProfile(context)
}

func (obj *ProfessorProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	profile_routes := rg.Group("")

	profile_routes.Use(middlewares.ValidateAuthorization([]string{"ADMIN"})).POST("/professors", obj.CreateProfessor)
	profile_routes.Use(middlewares.ValidateAuthorization([]string{"ADMIN"})).DELETE("/professors/:email_id", obj.DeleteProfessor)
	profile_routes.Use(middlewares.ValidateAuthorization([]string{"ADMIN"})).PUT("/professors/:email_id", obj.UpdateProfessor)
}
