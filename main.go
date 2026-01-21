package main

import (
	"go_app/src/configuration/logger"
	"go_app/src/controller/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application.")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
