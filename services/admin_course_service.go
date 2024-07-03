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

type AdminCourseService struct {
	client http.Client
}

func (obj *AdminCourseService) Init() {
	obj.client = http.Client{}
}

func (obj *AdminCourseService) courseGetAction(fetch_all bool, context *gin.Context) {
	url := os.Getenv("COURSE_SERVICE")+"/courses"
	if(!fetch_all) {
		url += "?course_id="+context.Query("course_id")
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

func (obj *AdminCourseService) CourseActions(action string, context *gin.Context) {
	action = strings.ToUpper(action)
	
	var url = os.Getenv("COURSE_SERVICE")+"/courses"

	switch action {
	case "GET" : 
		if(context.Query("course_id") != "") {
			obj.courseGetAction(false, context)
		} else {
			obj.courseGetAction(true, context)
		}
	case "POST": 
	case "PUT": fallthrough
	case "DELETE": {url += "/"+context.Param("course_id")}
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