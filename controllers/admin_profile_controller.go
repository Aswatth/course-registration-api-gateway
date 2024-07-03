package controllers

import (
	"course-registration-system/api-gateway/middlewares"
	"course-registration-system/api-gateway/services"

	"github.com/gin-gonic/gin"
)

type AdminProfileController struct {
	service services.AdminProfileService
}

func (obj *AdminProfileController) Init(service services.AdminProfileService) {
	obj.service = service
}

func (obj *AdminProfileController) CreateStudent(context *gin.Context) {
	obj.service.StudentProfileActions("POST", context)
}

func (obj *AdminProfileController) GetStudentProfile(context *gin.Context) {
	obj.service.StudentProfileActions("GET", context)
}

func (obj *AdminProfileController) UpdateStudent(context *gin.Context) {
	obj.service.StudentProfileActions("PUT", context)
}

func (obj *AdminProfileController) DeleteStudent(context *gin.Context) {
	obj.service.StudentProfileActions("DELETE", context)
}

func (obj *AdminProfileController) CreateProfessor(context *gin.Context) {
	obj.service.ProfessorProfileActions("POST", context)
}

func (obj *AdminProfileController) GetProfessorProfile(context *gin.Context) {
	obj.service.ProfessorProfileActions("GET", context)
}

func (obj *AdminProfileController) UpdateProfessor(context *gin.Context) {
	obj.service.ProfessorProfileActions("PUT", context)
}

func (obj *AdminProfileController) DeleteProfessor(context *gin.Context) {
	obj.service.ProfessorProfileActions("DELETE", context)
}

func (obj *AdminProfileController) UpdatePassword(context *gin.Context) {
	obj.service.UpdatePassword(context)
}

func (obj *AdminProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	admin_routes := rg.Group("/admin")

	admin_routes.Use(middlewares.ValidateAuthorization([]string{"ADMIN"}))

	admin_routes.POST("/students", obj.CreateStudent)
	admin_routes.GET("/students", obj.GetStudentProfile)
	admin_routes.DELETE("/students/:email_id", obj.DeleteStudent)
	admin_routes.PUT("/students/:email_id", obj.UpdateStudent)

	admin_routes.POST("/professors", obj.CreateProfessor)
	admin_routes.GET("/professors", obj.GetProfessorProfile)
	admin_routes.DELETE("/professors/:email_id", obj.DeleteProfessor)
	admin_routes.PUT("/professors/:email_id", obj.UpdateProfessor)

	admin_routes.PUT("/password/:email_id", obj.UpdatePassword)
}
