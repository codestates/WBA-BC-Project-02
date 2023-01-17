package model

import (
	"github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/model/contract"
	"github.com/codestates/WBA-BC-Project-02/was/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var AppModel Modeler

var UserModel user.UserModeler
var ContractModel contract.ContractModeler

func LoadMongoModel(URI, DBName string, colNames []string) error {
	m, err := NewModel(URI, DBName, colNames)
	if err != nil {
		return err
	}
	AppModel = m
	return nil
}

func InjectModelsMongoDependency(m map[string]*mongo.Collection) {
	UserModel = user.NewUserModel(m[enum.UserCollectionName])
	ContractModel = contract.NewContractModel(m[enum.ContractCollectionName])
}

func CreateIndexesInModels() {
	AppModel.CreateIndexes(enum.UserCollectionName, true, "private_key", "address")
	AppModel.CreateIndexes(enum.ContractCollectionName, true, "name")
}
