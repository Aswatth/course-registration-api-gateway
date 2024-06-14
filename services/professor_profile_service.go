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

func (obj *ProfessorProfileService) UpdateProfessorPassword(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/professors/password/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}
