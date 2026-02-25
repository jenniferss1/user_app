package view

import (
	"go_app/src/controller/model/response"
	"go_app/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserReponse {
	return response.UserReponse{
		Id:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
