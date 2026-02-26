package service

import (
	"go_app/src/configuration/rest_err"
	"go_app/src/model"
	"go_app/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.Resterr
	UpdateUser(string, model.UserDomainInterface) *rest_err.Resterr
	FindUser(string) (*model.UserDomainInterface, *rest_err.Resterr)
	DeleteUser(string) *rest_err.Resterr
}
