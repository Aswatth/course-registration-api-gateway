package services

import (
	"io"
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

func (obj *ProfessorProfileService) OfferCourse(context *gin.Context) {
	req, _ := http.NewRequest("POST", os.Getenv("REGISTRATION_SERVICE")+"/offered_course", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *ProfessorProfileService) GetOfferedCourse(context *gin.Context) ([]byte, error) {
	req, _ := http.NewRequest("GET", os.Getenv("REGISTRATION_SERVICE")+"/offered_course/"+context.Param("crn"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	response, _ := obj.client.Do(req)

	result, err := io.ReadAll(response.Body)

	return result, err
}

func (obj *ProfessorProfileService) UpdateOfferedCourse(context *gin.Context) {
	req, _ := http.NewRequest("PUT", os.Getenv("REGISTRATION_SERVICE")+"/offered_course/"+context.Param("crn"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}

func (obj *ProfessorProfileService) DeleteOfferedCourse(context *gin.Context) {
	req, _ := http.NewRequest("DELETE", os.Getenv("REGISTRATION_SERVICE")+"/offered_course/"+context.Param("crn"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	obj.client.Do(req)
}
