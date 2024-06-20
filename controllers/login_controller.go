package controllers

import (
	"course-registration-system/api-gateway/services"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	service services.LoginService
}

func (obj *LoginController) Init(service services.LoginService) {
	obj.service = service
}

func (obj *LoginController) Login(context *gin.Context) {
	token, status_code, err := obj.service.Login(context)

	if err != nil {
		context.AbortWithStatusJSON(status_code, gin.H{"response": err.Error()})
	} else {
		context.JSON(status_code, gin.H{"token": token})
	}
}

func (obj *LoginController) RegisterRoutes(rg *gin.RouterGroup) {
	login_routes := rg.Group("")

	login_routes.POST("/login", obj.Login)
}
