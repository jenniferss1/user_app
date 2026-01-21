package controller

import (
	"fmt"
	"go_app/src/configuration/logger"
	"go_app/src/configuration/validation"
	"go_app/src/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller")
	var userReuqest request.UserRequest

	if err := c.ShouldBindJSON(&userReuqest); err != nil {
		logger.Error("Error trying to validate user info", err)
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userReuqest)
}
