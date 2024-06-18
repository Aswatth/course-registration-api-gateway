package controllers

import (
	"course-registration-system/api-gateway/middlewares"
	"course-registration-system/api-gateway/services"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminCourseController struct {
	service services.AdminCourseService
}

func (obj *AdminCourseController) Init(service services.AdminCourseService) {
	obj.service = service
}

func (obj *AdminCourseController) CreateCourse(context *gin.Context) {
	obj.service.CreateCourse(context)
}

func (obj *AdminCourseController) GetCourse(context *gin.Context) {
	result, err := obj.service.GetCourse(context)

	if err != nil {
		context.AbortWithError(http.StatusBadGateway, err)
	}

	var data interface{}

	json.Unmarshal(result, &data)

	context.JSON(http.StatusOK, data)
}

func (obj *AdminCourseController) UpdateCourse(context *gin.Context) {
	obj.service.UpdateCourse(context)
}

func (obj *AdminCourseController) DeleteCourse(context *gin.Context) {
	obj.service.DeleteCourse(context)
}

func (obj *AdminCourseController) RegisterRoutes(rg *gin.RouterGroup) {
	admin_routes := rg.Group("/admin/courses")

	admin_routes.Use(middlewares.ValidateAuthorization([]string{"ADMIN"}))

	admin_routes.POST("", obj.CreateCourse)
	admin_routes.GET("/:course_id", obj.GetCourse)
	admin_routes.PUT("/:course_id", obj.UpdateCourse)
	admin_routes.DELETE("/:course_id", obj.DeleteCourse)
}
