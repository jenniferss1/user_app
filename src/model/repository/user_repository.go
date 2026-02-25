package repository

import (
	"go_app/src/configuration/rest_err"
	"go_app/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.Resterr)
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}
