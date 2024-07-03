package services

import (
	"bytes"
	"course-registration-system/api-gateway/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type LoginService struct {
}

func (obj *LoginService) Login(ctx *gin.Context) (string, int, error) {
	//Extract email_id from request body
	req_body, _ := io.ReadAll(ctx.Request.Body)
	req_body_data := make(map[string]string)
	json.Unmarshal(req_body, &req_body_data)
	email_id := req_body_data["email_id"]

	ctx.Request.Body = io.NopCloser(bytes.NewReader(req_body))

	response, err := http.Post(os.Getenv("PROFILE_SERVICE")+"/login", "json", ctx.Request.Body)

	if err != nil {
		return "", http.StatusInternalServerError, errors.New("unable to access login service")
	}

	if response.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(response.Body)

		data := make(map[string]string)

		json.Unmarshal(body, &data)

		
		token, err := utils.GenerateToken(data["user_type"], email_id)

		if err != nil {
			return "", http.StatusBadRequest, errors.New("unable to generate token")
		}

		return token, http.StatusOK, nil
	} else {
		response_body, _ := io.ReadAll(response.Body)

		data := make(map[string]string)
		json.Unmarshal(response_body, &data)

		return "", response.StatusCode, errors.New(data["response"])
	}
}
