package service

import (
	"go_app/src/configuration/rest_err"
	"go_app/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.Resterr
	UpdateUser(string, model.UserDomainInterface) *rest_err.Resterr
	FindUser(string) (*model.UserDomainInterface, *rest_err.Resterr)
	DeleteUser(string) *rest_err.Resterr
}
