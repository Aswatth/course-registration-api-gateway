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

func (obj *ProfessorProfileController) GetProfessorProfile(context *gin.Context) {
	obj.service.GetProfessorProfile(context)
}

func (obj *ProfessorProfileController) UpdateProfessorPassword(context *gin.Context) {
	obj.service.UpdateProfessorPassword(context)
}

func (obj *ProfessorProfileController) OfferCourse(context *gin.Context) {
	obj.service.OfferCourse(context)
}

func (obj *ProfessorProfileController) GetOfferedCourse(context *gin.Context) {
	obj.service.GetOfferedCourse(context)
}

func (obj *ProfessorProfileController) UpdateOfferedCourse(context *gin.Context) {
	obj.service.UpdateOfferedCourse(context)
}

func (obj *ProfessorProfileController) DeleteOfferedCourse(context *gin.Context) {
	obj.service.DeleteOfferedCourse(context)
}

func (obj *ProfessorProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	professor_routes := rg.Group("/professors")

	professor_routes.Use(middlewares.ValidateAuthorization([]string{"PROFESSOR"}))

	professor_routes.GET("/:email_id", obj.GetProfessorProfile)
	professor_routes.PUT("/password/:email_id", obj.UpdateProfessorPassword)

	professor_routes.POST("/offered_course", obj.OfferCourse)
	professor_routes.GET("/offered_course/:crn", obj.GetOfferedCourse)
	professor_routes.PUT("/offered_course/:crn", obj.UpdateOfferedCourse)
	professor_routes.DELETE("/offered_course/:crn", obj.DeleteOfferedCourse)
}
