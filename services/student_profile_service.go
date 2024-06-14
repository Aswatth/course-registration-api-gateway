package services

import (
	"net/http"
)

type StudentProfileService struct {
	client http.Client
}

func (obj *StudentProfileService) Init() {
	obj.client = http.Client{}
}
