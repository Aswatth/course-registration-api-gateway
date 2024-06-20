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

func (obj *StudentProfileController) GetStudentProfile(context *gin.Context) {
	obj.service.GetStudentProfile(context)
}

func (obj *StudentProfileController) UpdateStudentPassword(context *gin.Context) {
	obj.service.UpdateStudentPassword(context)
}

func (obj *StudentProfileController) RegisterCourse(context *gin.Context) {
	obj.service.RegisterCourse(context)
}

func (obj *StudentProfileController) GetRegisteredCourses(context *gin.Context) {
	obj.service.GetRegisteredCourse(context)
}

func (obj *StudentProfileController) UpdateRegisteredCourses(context *gin.Context) {
	obj.service.UpdateRegisteredCourses(context)
}

func (obj *StudentProfileController) DeleteRegisteredCourses(context *gin.Context) {
	obj.service.DeleteRegisteredCourses(context)
}

func (obj *StudentProfileController) RegisterRoutes(rg *gin.RouterGroup) {
	student_routes := rg.Group("/students")

	student_routes.Use(middlewares.ValidateAuthorization([]string{"STUDENT"}))

	student_routes.GET("/:email_id", obj.GetStudentProfile)
	student_routes.PUT("/password/:email_id", obj.UpdateStudentPassword)

	student_routes.POST("/register_course", obj.RegisterCourse)
	student_routes.GET("/register_course/:student_email_id", obj.GetRegisteredCourses)
	student_routes.DELETE("/register_course/:student_email_id", obj.DeleteRegisteredCourses)
	student_routes.PUT("/register_course/:student_email_id", obj.UpdateRegisteredCourses)
}
