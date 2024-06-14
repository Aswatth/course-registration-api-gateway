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

func (obj *StudentProfileService) UpdateStudentPassword(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/students/password/"+context.Param("email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}
