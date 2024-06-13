package services

import (
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

func (obj *StudentProfileService) CreateStudentProfile(context *gin.Context) {
	req, _ := http.NewRequest("POST", os.Getenv("PROFILE_SERVICE")+"/admin/students", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *StudentProfileService) UpdateStudentProfile(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/admin/students/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *StudentProfileService) DeleteStudentProfile(context *gin.Context) {
	req, _ := http.NewRequest("DELETE", os.Getenv("PROFILE_SERVICE")+"/admin/students/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}
