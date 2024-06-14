package services

import (
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
	response, err := http.Post(os.Getenv("PROFILE_SERVICE")+"/login", "json", ctx.Request.Body)

	if err != nil {
		return "", http.StatusBadRequest, errors.New("unable to login")
	}

	if response.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(response.Body)

		type user_type struct {
			User_type string
		}

		var user_type_data user_type

		json.Unmarshal(body, &user_type_data)

		token, err := utils.GenerateToken(user_type_data.User_type)

		if err != nil {
			return "", http.StatusBadRequest, errors.New("unable to generate token")
		}

		return token, http.StatusOK, nil
	}

	return "", http.StatusBadRequest, errors.New("unable to login")
}
