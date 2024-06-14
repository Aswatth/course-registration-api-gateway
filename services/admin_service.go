package services

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AdminService struct {
	client http.Client
}

func (obj *AdminService) Init() {
	obj.client = http.Client{}
}

func (obj *AdminService) CreateStudentProfile(context *gin.Context) {
	req, _ := http.NewRequest("POST", os.Getenv("PROFILE_SERVICE")+"/admin/students", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminService) UpdateStudentProfile(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/admin/students/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminService) DeleteStudentProfile(context *gin.Context) {
	req, _ := http.NewRequest("DELETE", os.Getenv("PROFILE_SERVICE")+"/admin/students/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminService) CreateProfessorProfile(context *gin.Context) {
	req, _ := http.NewRequest("POST", os.Getenv("PROFILE_SERVICE")+"/admin/professors", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminService) UpdateProfessorProfile(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/admin/professors/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminService) DeleteProfessorProfile(context *gin.Context) {
	req, _ := http.NewRequest("DELETE", os.Getenv("PROFILE_SERVICE")+"/admin/professors/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}
