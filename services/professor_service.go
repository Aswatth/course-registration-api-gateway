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

type ProfessorProfileService struct {
	client http.Client
}

func (obj *ProfessorProfileService) Init() {
	obj.client = http.Client{}
}

func (obj *ProfessorProfileService) GetProfessorProfile(context *gin.Context) {
	req, err := http.NewRequest("GET", os.Getenv("PROFILE_SERVICE")+"/professors/"+context.Param("email_id"), context.Request.Body)

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

			context.JSON(response.StatusCode, data)
		}
	}
}

func (obj *ProfessorProfileService) UpdateProfessorPassword(context *gin.Context) {
	req, err := http.NewRequest("PUT", os.Getenv("PROFILE_SERVICE")+"/professors/password/"+context.Param("email_id"), context.Request.Body)

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

func (obj *ProfessorProfileService) GetAvailableCourses(context *gin.Context) {
	req, err := http.NewRequest("GET", os.Getenv("COURSE_SERVICE")+"/courses", context.Request.Body)

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

func (obj *ProfessorProfileService) OfferCourse(context *gin.Context) {
	//Check if course exists
	request_data := make(map[string]any)

	request_body, _ := io.ReadAll(context.Request.Body)

	json.Unmarshal(request_body, &request_data)

	new_req, _ := http.NewRequest("GET", os.Getenv("COURSE_SERVICE")+"/courses?course_id="+fmt.Sprint(request_data["course_id"]), context.Request.Body)
	response, _ := obj.client.Do(new_req)

	resp_body, _ := io.ReadAll(response.Body)

	response_data := make(map[string]any)
	
	json.Unmarshal(resp_body, &response_data)

	if fmt.Sprint(request_data["course_id"]) != fmt.Sprint(response_data["course_id"]) {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "invalid course_id"})
	} else {
		context.Request.Body = io.NopCloser(bytes.NewReader(request_body))

		req, err := http.NewRequest("POST", os.Getenv("REGISTRATION_SERVICE")+"/offered_course", context.Request.Body)

		if err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error creating new request"})
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
}
func (obj *ProfessorProfileService) GetOfferedCoursesByCRN(context *gin.Context) {
	req, err := http.NewRequest("GET", os.Getenv("REGISTRATION_SERVICE")+"/offered_course?crn="+context.Query("crn"), context.Request.Body)

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

			var result_data interface{}
			json.Unmarshal(result, &result_data)
			context.JSON(response.StatusCode, result_data)
		}
	}
}

func (obj *ProfessorProfileService) GetOfferedCoursesByProfessor(context *gin.Context) {
	req, err := http.NewRequest("GET", os.Getenv("REGISTRATION_SERVICE")+"/offered_course?email_id="+context.Query("email_id"), context.Request.Body)

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

			var result_data interface{}
			json.Unmarshal(result, &result_data)
			context.JSON(response.StatusCode, result_data)
		}
	}
}

func (obj *ProfessorProfileService) UpdateOfferedCourse(context *gin.Context) {
	req, err := http.NewRequest("PUT", os.Getenv("REGISTRATION_SERVICE")+"/offered_course/"+context.Param("crn"), context.Request.Body)

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

			result_data := make(map[string]any)
			json.Unmarshal(result, &result_data)

			context.JSON(response.StatusCode, result_data)
		}
	}
}

func (obj *ProfessorProfileService) DeleteOfferedCourse(context *gin.Context) {
	//Check if it is registered by any student before deleting
	req, err := http.NewRequest("GET", os.Getenv("REGISTRATION_SERVICE")+"/register_course?crn="+context.Param("crn"), context.Request.Body)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error creating new request"})
	} else {
		response, err := obj.client.Do(req)

		if err != nil {
			log.Println(err.Error())
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"response": "error executing request"})
		} else {
			result, _ := io.ReadAll(response.Body)

			var result_data interface{}

			json.Unmarshal(result, &result_data)

			if(result_data != nil) {
				context.JSON(http.StatusBadRequest, gin.H{"response": "cannot delete a registered course"})
			} else {
				req, err := http.NewRequest("DELETE", os.Getenv("REGISTRATION_SERVICE")+"/offered_course/"+context.Param("crn"), context.Request.Body)

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

						result_data := make(map[string]any)
						json.Unmarshal(result, &result_data)

						context.JSON(response.StatusCode, result_data)
					}
				}
			}
		}
	}
}
