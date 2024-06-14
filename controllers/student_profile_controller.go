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

func (obj *StudentProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	profile_routes := rg.Group("")

	profile_routes.Use(middlewares.ValidateAuthorization([]string{"STUDENT"}))
}
