package controllers

import (
	"course-registration-system/api-gateway/middlewares"
	"course-registration-system/api-gateway/services"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminProfileController struct {
	service services.AdminProfileService
}

func (obj *AdminProfileController) Init(service services.AdminProfileService) {
	obj.service = service
}

func (obj *AdminProfileController) CreateStudent(context *gin.Context) {
	obj.service.CreateStudentProfile(context)
}

func (obj *AdminProfileController) GetStudentProfile(context *gin.Context) {
	result, err := obj.service.GetStudentProfile(context)

	if err != nil {
		context.AbortWithError(http.StatusBadGateway, err)
	}

	var data interface{}
	json.Unmarshal(result, &data)

	context.JSON(http.StatusOK, data)
}

func (obj *AdminProfileController) GetAllStudentProfiles(context *gin.Context) {
	if(context.Query("email_id") != "") {
		obj.GetStudentProfile(context)
		return
	}

	result, err := obj.service.GetStudentProfile(context)

	if err != nil {
		context.AbortWithError(http.StatusBadGateway, err)
	}

	var data interface{}
	json.Unmarshal(result, &data)

	context.JSON(http.StatusOK, data)
}

func (obj *AdminProfileController) UpdateStudent(context *gin.Context) {
	obj.service.UpdateStudentProfile(context)
}

func (obj *AdminProfileController) DeleteStudent(context *gin.Context) {
	obj.service.DeleteStudentProfile(context)
}

func (obj *AdminProfileController) CreateProfessor(context *gin.Context) {
	obj.service.CreateProfessorProfile(context)
}

func (obj *AdminProfileController) GetProfessorProfile(context *gin.Context) {
	result, err := obj.service.GetProfessorProfile(context)

	if err != nil {
		context.AbortWithError(http.StatusBadGateway, err)
	}

	m := make(map[string]string)
	json.Unmarshal(result, &m)

	context.JSON(http.StatusOK, m)
}

func (obj *AdminProfileController) GetAllProfessorProfiles(context *gin.Context) {
	if(context.Query("email_id") != "") {
		obj.GetProfessorProfile(context)
		return
	}

	result, err := obj.service.GetAllProfessorProfiles(context)

	if err != nil {
		context.AbortWithError(http.StatusBadGateway, err)
	}

	var data interface{}
	json.Unmarshal(result, &data)

	context.JSON(http.StatusOK, data)
}

func (obj *AdminProfileController) UpdateProfessor(context *gin.Context) {
	obj.service.UpdateProfessorProfile(context)
}

func (obj *AdminProfileController) DeleteProfessor(context *gin.Context) {
	obj.service.DeleteProfessorProfile(context)
}

func (obj *AdminProfileController) UpdatePassword(context *gin.Context) {
	obj.service.UpdatePassword(context)
}

func (obj *AdminProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	admin_routes := rg.Group("/admin")

	admin_routes.Use(middlewares.ValidateAuthorization([]string{"ADMIN"}))

	admin_routes.POST("/students", obj.CreateStudent)
	admin_routes.GET("/students", obj.GetAllStudentProfiles)
	admin_routes.DELETE("/students/:email_id", obj.DeleteStudent)
	admin_routes.PUT("/students/:email_id", obj.UpdateStudent)

	admin_routes.POST("/professors", obj.CreateProfessor)
	admin_routes.GET("/professors", obj.GetAllProfessorProfiles)
	admin_routes.DELETE("/professors/:email_id", obj.DeleteProfessor)
	admin_routes.PUT("/professors/:email_id", obj.UpdateProfessor)

	admin_routes.PUT("/password/:email_id", obj.UpdatePassword)
}
