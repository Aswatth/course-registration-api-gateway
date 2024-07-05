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

func (obj *StudentProfileService) GetAllOfferedCourses(context *gin.Context) {
	req, err := http.NewRequest("GET", os.Getenv("REGISTRATION_SERVICE")+"/offered_course", context.Request.Body)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error creating new request"})
	} else {
		req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

		response, err := obj.client.Do(req)

		if err != nil {
			log.Println(err.Error())
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error executing request"})
		} else {
			result, _ := io.ReadAll(response.Body)

			var result_data []interface{}

			json.Unmarshal(result, &result_data)

			//Get course details for each offered course
			for _, data := range result_data {
				course_id := fmt.Sprint(data.(map[string]interface{})["course_id"])
				course_req, err := http.NewRequest("GET", os.Getenv("COURSE_SERVICE")+"/courses?course_id="+course_id, context.Request.Body)

				if err != nil {
					context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error creating new request"})
				} else {
					course_req.Header.Set("Authorization", context.Request.Header.Get("Authorization"))

					course_response, err := obj.client.Do(course_req)

					if err != nil {
						log.Println(err.Error())
						context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error executing request"})
					} else {
						course_result, _ := io.ReadAll(course_response.Body)

						var course_result_data interface{}

						json.Unmarshal(course_result, &course_result_data)

						data.(map[string]interface{})["course_info"] = course_result_data
					}
				}
			}

			context.JSON(response.StatusCode, result_data)
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
