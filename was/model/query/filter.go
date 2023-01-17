package query

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func GetDefaultIDFilter(obID primitive.ObjectID) bson.M {
	return bson.M{enum.ID: obID}
}

func GetAddressFilter(address string) bson.M {
	return bson.M{enum.Address: address}
}

func GetContractAddressFilter(contractAddr string) bson.M {
	return bson.M{enum.ContractAddress: contractAddr}
}

func GetUpdatePWDFilter(encryptPassword string) bson.M {
	return bson.M{enum.QuerySet: bson.D{
		bson.E{Key: enum.Password, Value: encryptPassword},
		bson.E{Key: enum.UpdateAt, Value: time.Now()},
	}}
}

func GetBlackIronIncreaseFilter() bson.D {
	return bson.D{
		bson.E{Key: enum.QueryInc, Value: bson.M{enum.BlackIron: 1}},
		bson.E{Key: enum.QuerySet, Value: bson.M{enum.UpdateAt: time.Now()}},
	}
}

func GetBlackIronSetFilter(blackIron int) bson.D {
	return bson.D{
		{enum.QuerySet,
			bson.D{
				{enum.BlackIron, blackIron},
				{enum.UpdateAt, time.Now()},
			},
		},
	}
}

func GetNameFilter(name string) bson.M {
	return bson.M{enum.Name: name}
}
