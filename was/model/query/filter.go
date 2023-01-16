package query

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func GetDefaultIDFilter(obID primitive.ObjectID) bson.M {
	return bson.M{"_id": obID}
}

func GetAddressFilter(address string) bson.M {
	return bson.M{"address": address}
}

func GetUpdatePWDFilter(encryptPassword string) bson.M {
	return bson.M{"$set": bson.D{
		bson.E{Key: "password", Value: encryptPassword},
		bson.E{Key: "base_time.updated_at", Value: time.Now()},
	}}
}

func GetBlackIronIncreaseFilter() bson.D {
	return bson.D{
		bson.E{Key: "$inc", Value: bson.M{"black_iron": 1}},
		bson.E{Key: "$set", Value: bson.M{"base_time.updated_at": time.Now()}},
	}
}

func GetNameFilter(name string) bson.M {
	return bson.M{"name": name}
}
