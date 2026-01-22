package controller

import (
	"go_app/src/configuration/logger"
	"go_app/src/configuration/validation"
	"go_app/src/controller/model/request"
	"go_app/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
		userRequest.Weight,
		userRequest.Height,
	)
	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
	}

	c.String(http.StatusOK, "")
}
