package services

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type ProfessorProfileService struct {
	client http.Client
}

func (obj *ProfessorProfileService) Init() {
	obj.client = http.Client{}
}

func (obj *ProfessorProfileService) CreateProfessorProfile(context *gin.Context) {
	req, _ := http.NewRequest("POST", os.Getenv("PROFILE_SERVICE")+"admin/professors", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *ProfessorProfileService) UpdateProfessorProfile(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/admin/professors/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *ProfessorProfileService) DeleteProfessorProfile(context *gin.Context) {
	req, _ := http.NewRequest("DELETE", os.Getenv("PROFILE_SERVICE")+"/admin/professors/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}
