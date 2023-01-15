package query

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetDefaultIDFilter(obID primitive.ObjectID) bson.M {
	return bson.M{"_id": obID}
}

func GetAddressFilter(address string) bson.M {
	return bson.M{"address": address}
}

func GetUpdatePWDFilter(encryptPassword string) bson.M {
	return bson.M{"$set": bson.M{"password": encryptPassword}}
}
