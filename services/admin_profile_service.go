package services

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AdminProfileService struct {
	client http.Client
}

func (obj *AdminProfileService) Init() {
	obj.client = http.Client{}
}

func (obj *AdminProfileService) CreateStudentProfile(context *gin.Context) {
	req, _ := http.NewRequest("POST", os.Getenv("PROFILE_SERVICE")+"/admin/students", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminProfileService) GetStudentProfile(context *gin.Context) ([]byte, error) {
	req, _ := http.NewRequest("GET", os.Getenv("PROFILE_SERVICE")+"/admin/students?email_id="+context.Query("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	response, _ := obj.client.Do(req)

	return io.ReadAll(response.Body)
}

func (obj *AdminProfileService) GetAllStudentProfiles(context *gin.Context) ([]byte, error) {
	req, _ := http.NewRequest("GET", os.Getenv("PROFILE_SERVICE")+"/admin/students", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	response, _ := obj.client.Do(req)

	return io.ReadAll(response.Body)
}

func (obj *AdminProfileService) UpdateStudentProfile(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/admin/students/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminProfileService) DeleteStudentProfile(context *gin.Context) {
	req, _ := http.NewRequest("DELETE", os.Getenv("PROFILE_SERVICE")+"/admin/students/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminProfileService) CreateProfessorProfile(context *gin.Context) {
	req, _ := http.NewRequest("POST", os.Getenv("PROFILE_SERVICE")+"/admin/professors", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminProfileService) GetProfessorProfile(context *gin.Context) ([]byte, error) {
	req, _ := http.NewRequest("GET", os.Getenv("PROFILE_SERVICE")+"/admin/professors?email_id="+context.Query("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	response, _ := obj.client.Do(req)

	return io.ReadAll(response.Body)
}

func (obj *AdminProfileService) GetAllProfessorProfiles(context *gin.Context) ([]byte, error) {
	req, _ := http.NewRequest("GET", os.Getenv("PROFILE_SERVICE")+"/admin/professors", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	response, _ := obj.client.Do(req)

	return io.ReadAll(response.Body)
}

func (obj *AdminProfileService) UpdateProfessorProfile(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/admin/professors/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminProfileService) DeleteProfessorProfile(context *gin.Context) {
	req, _ := http.NewRequest("DELETE", os.Getenv("PROFILE_SERVICE")+"/admin/professors/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *AdminProfileService) UpdatePassword(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/admin/password/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}
