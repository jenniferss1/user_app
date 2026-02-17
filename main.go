package main

import (
	"go_app/src/configuration/database/mongodb"
	"go_app/src/configuration/logger"
	"go_app/src/controller"
	"go_app/src/controller/routes"
	"go_app/src/model/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("About to start user application.")

	mongodb.InitConnection()

	// Innit dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
