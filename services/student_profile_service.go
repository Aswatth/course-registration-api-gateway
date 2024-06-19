package services

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type StudentProfileService struct {
	client http.Client
}

func (obj *StudentProfileService) Init() {
	obj.client = http.Client{}
}

func (obj *StudentProfileService) UpdateStudentPassword(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/students/password/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *StudentProfileService) RegisterCourse(context *gin.Context) {
	req, _ := http.NewRequest("POST", os.Getenv("REGISTRATION_SERVICE")+"/register_course", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *StudentProfileService) GetRegisteredCourse(context *gin.Context) ([]byte, error) {
	req, _ := http.NewRequest("GET", os.Getenv("REGISTRATION_SERVICE")+"/register_course/"+context.Param("student_email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	response, _ := obj.client.Do(req)

	return io.ReadAll(response.Body)
}

func (obj *StudentProfileService) UpdateRegisteredCourses(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("REGISTRATION_SERVICE")+"/register_course/"+context.Param("student_email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *StudentProfileService) DeleteRegisteredCourses(context *gin.Context) {
	req, _ := http.NewRequest("DELETE", os.Getenv("REGISTRATION_SERVICE")+"/register_course/"+context.Param("student_email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}
