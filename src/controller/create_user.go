package controller

import (
	"fmt"
	"go_app/src/configuration/rest_err"
	"go_app/src/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userReuqest request.UserRequest

	if err := c.ShouldBindJSON(&userReuqest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect field, error = %s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userReuqest)
}
