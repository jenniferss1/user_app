package service

import (
	"fmt"
	"go_app/src/configuration/logger"
	"go_app/src/configuration/rest_err"
	"go_app/src/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.Resterr {
	logger.Info("Init createUser model")
	userDomain.EncryptPassword()

	fmt.Println(ud)
	return nil
}
