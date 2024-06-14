package controllers

import (
	"course-registration-system/api-gateway/middlewares"
	"course-registration-system/api-gateway/services"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	service services.AdminService
}

func (obj *AdminController) Init(service services.AdminService) {
	obj.service = service
}

func (obj *AdminController) CreateStudent(context *gin.Context) {
	obj.service.CreateStudentProfile(context)
}

func (obj *AdminController) UpdateStudent(context *gin.Context) {
	obj.service.UpdateStudentProfile(context)
}

func (obj *AdminController) DeleteStudent(context *gin.Context) {
	obj.service.DeleteStudentProfile(context)
}

func (obj *AdminController) CreateProfessor(context *gin.Context) {
	obj.service.CreateProfessorProfile(context)
}

func (obj *AdminController) UpdateProfessor(context *gin.Context) {
	obj.service.UpdateProfessorProfile(context)
}

func (obj *AdminController) DeleteProfessor(context *gin.Context) {
	obj.service.DeleteProfessorProfile(context)
}

func (obj *AdminController) RegisterRoutes(rg *gin.RouterGroup) {
	admin_routes := rg.Group("admin")

	admin_routes.Use(middlewares.ValidateAuthorization([]string{"ADMIN"}))

	admin_routes.POST("/students", obj.CreateStudent)
	admin_routes.DELETE("/students/:email_id", obj.DeleteStudent)
	admin_routes.PUT("/students/:email_id", obj.UpdateStudent)

	admin_routes.POST("/professors", obj.CreateProfessor)
	admin_routes.DELETE("/professors/:email_id", obj.DeleteProfessor)
	admin_routes.PUT("/professors/:email_id", obj.UpdateProfessor)
}
