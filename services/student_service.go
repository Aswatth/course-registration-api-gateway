package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func (obj *StudentProfileService) GetStudentProfile(context *gin.Context) {
	req, err := http.NewRequest("GET", os.Getenv("PROFILE_SERVICE")+"/students/"+context.Param("email_id"), context.Request.Body)

	if err != nil {
		log.Println(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error creating a new request"})
	} else {
		req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

		response, err := obj.client.Do(req)

		if err != nil {
			log.Println(err.Error())
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error executing request"})
		} else {
			var data interface{}

			body, _ := io.ReadAll(response.Body)

			json.Unmarshal(body, &data)

			if data == nil {
				context.Status(response.StatusCode)
			} else {
				context.JSON(response.StatusCode, data)
			}
		}
	}
}

func (obj *StudentProfileService) UpdateStudentPassword(context *gin.Context) {
	req, err := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/students/password/"+context.Param("email_id"), context.Request.Body)

	if err != nil {
		log.Println(err.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error creating a new request"})
	} else {
		req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

		response, err := obj.client.Do(req)

		if err != nil {
			log.Println(err.Error())
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error executing request"})
		} else {
			var data interface{}

			body, _ := io.ReadAll(response.Body)

			json.Unmarshal(body, &data)

			if data == nil {
				context.Status(response.StatusCode)
			} else {
				context.JSON(response.StatusCode, data)
			}
		}
	}
}

func (obj *StudentProfileService) RegisterCourse(context *gin.Context) {
	//check if crns are valid
	request_data := make(map[string]any)

	request_body, _ := io.ReadAll(context.Request.Body)

	json.Unmarshal(request_body, &request_data)

	for _, crn := range request_data["registered_course_crns"].([]interface{}) {
		new_req, _ := http.NewRequest("GET", os.Getenv("REGISTRATION_SERVICE")+"/offered_course/"+fmt.Sprint(crn), context.Request.Body)
		response, _ := obj.client.Do(new_req)

		if response.StatusCode != http.StatusOK {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": " invalid crn " + fmt.Sprint(crn)})
			return
		}
	}

	context.Request.Body = io.NopCloser(bytes.NewReader(request_body))

	req, _ := http.NewRequest("POST", os.Getenv("REGISTRATION_SERVICE")+"/register_course", context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	response, err := obj.client.Do(req)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		var data interface{}

		body, _ := io.ReadAll(response.Body)

		json.Unmarshal(body, &data)

		if data == nil {
			context.Status(response.StatusCode)
		} else {
			context.JSON(response.StatusCode, data)
		}
	}
}

func (obj *StudentProfileService) GetRegisteredCourse(context *gin.Context) {
	req, err := http.NewRequest("GET", os.Getenv("REGISTRATION_SERVICE")+"/register_course/"+context.Param("student_email_id"), context.Request.Body)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

		response, err := obj.client.Do(req)

		if err != nil {
			context.AbortWithStatusJSON(response.StatusCode, gin.H{"response": err.Error()})
		} else {
			var data interface{}

			body, _ := io.ReadAll(response.Body)

			json.Unmarshal(body, &data)

			if data == nil {
				context.Status(response.StatusCode)
			} else {
				context.JSON(response.StatusCode, data)
			}
		}
	}
}

func (obj *StudentProfileService) UpdateRegisteredCourses(context *gin.Context) {
	//check if crns are valid
	request_data := make(map[string]any)

	request_body, _ := io.ReadAll(context.Request.Body)

	json.Unmarshal(request_body, &request_data)

	fmt.Println(request_data["registered_course_crns"].([]interface{}))

	for _, crn := range request_data["registered_course_crns"].([]interface{}) {
		fmt.Println(crn)

		new_req, _ := http.NewRequest("GET", os.Getenv("REGISTRATION_SERVICE")+"/offered_course/"+fmt.Sprint(crn), context.Request.Body)
		response, _ := obj.client.Do(new_req)

		if response.StatusCode != http.StatusOK {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "invalid crn " + fmt.Sprint(crn)})
			return
		}

		context.Request.Body = io.NopCloser(bytes.NewReader(request_body))
	}

	context.Request.Body = io.NopCloser(bytes.NewReader(request_body))

	req, _ := http.NewRequest("PUT", os.Getenv("REGISTRATION_SERVICE")+"/register_course/"+context.Param("student_email_id"), context.Request.Body)

	req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

	response, err := obj.client.Do(req)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		var data interface{}

		body, _ := io.ReadAll(response.Body)

		json.Unmarshal(body, &data)

		if data == nil {
			context.Status(response.StatusCode)
		} else {
			context.JSON(response.StatusCode, data)
		}
	}
}

func (obj *StudentProfileService) DeleteRegisteredCourses(context *gin.Context) {
	req, err := http.NewRequest("DELETE", os.Getenv("REGISTRATION_SERVICE")+"/register_course/"+context.Param("student_email_id"), context.Request.Body)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": err.Error()})
	} else {
		req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

		response, err := obj.client.Do(req)

		if err != nil {
			context.AbortWithStatusJSON(response.StatusCode, gin.H{"response": err.Error()})
		} else {
			var data interface{}

			body, _ := io.ReadAll(response.Body)

			json.Unmarshal(body, &data)

			if data == nil {
				context.Status(response.StatusCode)
			} else {
				context.JSON(response.StatusCode, data)
			}
		}
	}
}
