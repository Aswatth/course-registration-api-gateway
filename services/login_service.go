package services

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type LoginService struct {
}

func (obj *LoginService) Login(ctx *gin.Context) (string, int) {
	response, err := http.Post(os.Getenv("PROFILE_SERVICE")+"/login", "json", ctx.Request.Body)

	if err != nil {
		return "", http.StatusBadRequest
	}

	if response.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(response.Body)

		type token struct {
			Token string
		}

		var token_data token

		json.Unmarshal(body, &token_data)

		return token_data.Token, http.StatusOK
	}

	return "", http.StatusBadRequest
}
