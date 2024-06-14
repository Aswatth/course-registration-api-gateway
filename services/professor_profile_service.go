package services

import (
	"net/http"
)

type ProfessorProfileService struct {
	client http.Client
}

func (obj *ProfessorProfileService) Init() {
	obj.client = http.Client{}
}
