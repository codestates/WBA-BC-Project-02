package user

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var instance *userModel

type userModel struct {
	collection *mongo.Collection
}

func NewUserModel(col *mongo.Collection) *userModel {
	if instance != nil {
		return instance
	}
	instance = &userModel{
		collection: col,
	}
	return instance
}
