package model

import (
	"fmt"
	"go_app/src/configuration/logger"
	"go_app/src/configuration/rest_err"
)

func (ud *UserDomain) CreateUser() *rest_err.Resterr {
	logger.Info("Init createUser model")
	ud.EncryptPassword()

	fmt.Println(ud)
	return nil
}
