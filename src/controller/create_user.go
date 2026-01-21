package controller

import (
	"fmt"
	"go_app/src/configuration/validation"
	"go_app/src/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userReuqest request.UserRequest

	if err := c.ShouldBindJSON(&userReuqest); err != nil {
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userReuqest)
}
