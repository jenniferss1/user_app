package main

import (
	"context"
	"go_app/src/configuration/database/mongodb"
	"go_app/src/configuration/logger"
	"go_app/src/controller/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("About to start user application.")

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
