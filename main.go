package main

import (
	"go_app/src/configuration/logger"
	"go_app/src/controller/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("About to start user application.")

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
