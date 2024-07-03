package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type AdminProfileService struct {
	client http.Client
}

func (obj *AdminProfileService) Init() {
	obj.client = http.Client{}
}

func (obj *AdminProfileService) studentGetAction(fetch_all bool, context *gin.Context) {
	url := ""
	if(fetch_all) {
		url = os.Getenv("PROFILE_SERVICE")+"/admin/students"
	} else {
		url = os.Getenv("PROFILE_SERVICE")+"/admin/students?email_id="+context.Query("email_id")
	}

	req, err := http.NewRequest("GET", url, context.Request.Body)

	if err != nil {
		log.Println(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error creating a new request"})
	} else {
		req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

		response, err := obj.client.Do(req)

		if err != nil {
			log.Println(err.Error())
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error executing a new request"})
		} else {
			var data interface{}

			body, _ := io.ReadAll(response.Body)

			json.Unmarshal(body, &data)

			context.JSON(response.StatusCode, data)
		}
	}
}
func (obj *AdminProfileService) StudentProfileActions(action string, context *gin.Context) {
	action = strings.ToUpper(action)
	
	var url = os.Getenv("PROFILE_SERVICE")+"/admin/students"

	switch action {
	case "GET" : 
		if(context.Query("email_id") != "") {
			obj.studentGetAction(false, context)
		} else {
			obj.studentGetAction(true, context)
		}
	case "POST": 
	case "PUT": fallthrough
	case "DELETE": {url += "/"+context.Param("email_id")}
	}

	req, err := http.NewRequest(action, url, context.Request.Body)

	if err != nil {
		log.Println(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error creating a new request"})
	} else {
		req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

		response, err := obj.client.Do(req)

		if (err != nil) {
			log.Println(err.Error())
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error executing a new request"})
		} else {
			if(response.StatusCode == 200) {
				context.Status(http.StatusOK)
			} else {
				var data interface{}

				body, _ := io.ReadAll(response.Body)

				json.Unmarshal(body, &data)

				context.JSON(response.StatusCode, data)
			}
		}
	}
}

func (obj *AdminProfileService) professorGetAction(fetch_all bool, context *gin.Context) {
	url := ""
	if(fetch_all) {
		url = os.Getenv("PROFILE_SERVICE")+"/admin/professors"
	} else {
		url = os.Getenv("PROFILE_SERVICE")+"/admin/professors?email_id="+context.Query("email_id")
	}

	req, err := http.NewRequest("GET", url, context.Request.Body)

	if err != nil {
		log.Println(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error creating a new request"})
	} else {
		req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

		response, err := obj.client.Do(req)

		if err != nil {
			log.Println(err.Error())
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error executing a new request"})
		} else {
			var data interface{}

			body, _ := io.ReadAll(response.Body)

			json.Unmarshal(body, &data)

			context.JSON(response.StatusCode, data)
		}
	}
}
func (obj *AdminProfileService) ProfessorProfileActions(action string, context *gin.Context) {
	action = strings.ToUpper(action)
	
	var url = os.Getenv("PROFILE_SERVICE")+"/admin/professors"

	switch action {
	case "GET" : 
		if(context.Query("email_id") != "") {
			obj.professorGetAction(false, context)
		} else {
			obj.professorGetAction(true, context)
		}
	case "POST": 
	case "PUT": fallthrough
	case "DELETE": {url += "/"+context.Param("email_id")}
	}

	req, err := http.NewRequest(action, url, context.Request.Body)

	if err != nil {
		log.Println(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error creating a new request"})
	} else {
		req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

		response, err := obj.client.Do(req)

		if (err != nil) {
			log.Println(err.Error())
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error executing a new request"})
		} else {
			if(response.StatusCode == 200) {
				context.Status(http.StatusOK)
			} else {
				var data interface{}

				body, _ := io.ReadAll(response.Body)

				json.Unmarshal(body, &data)

				context.JSON(response.StatusCode, data)
			}
		}
	}
}

func (obj *AdminProfileService) UpdatePassword(context *gin.Context) {
	req, err := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/admin/password/"+context.Param("email_id"), context.Request.Body)

	if err != nil {
		log.Println(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error creating a new request"})
	} else {
		req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

		response, err := obj.client.Do(req)

		if (err != nil) {
			log.Println(err.Error())
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error executing a new request"})
		} else {
			if(response.StatusCode == 200) {
				context.Status(http.StatusOK)
			} else {
				var data interface{}

				body, _ := io.ReadAll(response.Body)

				json.Unmarshal(body, &data)

				context.JSON(response.StatusCode, data)
			}
		}
	}
}