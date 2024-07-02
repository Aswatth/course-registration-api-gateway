package services

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AdminCourseService struct {
	client http.Client
}

func (obj *AdminCourseService) Init() {
	obj.client = http.Client{}
}

func (obj *AdminCourseService) CreateCourse(context *gin.Context) {
	req, _ := http.NewRequest("POST", os.Getenv("COURSE_SERVICE")+"/courses", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminCourseService) GetCourse(context *gin.Context) ([]byte, error) {
	req, _ := http.NewRequest("GET", os.Getenv("COURSE_SERVICE")+"/courses?course_id="+context.Query("course_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	response, _ := obj.client.Do(req)

	return io.ReadAll(response.Body)
}

func (obj *AdminCourseService) GetAllCourses(context *gin.Context) ([]byte, error) {
	req, _ := http.NewRequest("GET", os.Getenv("COURSE_SERVICE")+"/courses", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	response, _ := obj.client.Do(req)

	return io.ReadAll(response.Body)
}

func (obj *AdminCourseService) DeleteCourse(context *gin.Context) {
	req, _ := http.NewRequest("DELETE", os.Getenv("COURSE_SERVICE")+"/courses/"+context.Param("course_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminCourseService) UpdateCourse(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("COURSE_SERVICE")+"/courses/"+context.Param("course_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}
