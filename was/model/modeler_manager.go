package model

import (
	"github.com/codestates/WBA-BC-Project-02/common/model"
	"github.com/codestates/WBA-BC-Project-02/was/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var AppModel model.Modeler

var UserModel user.UserModeler

func LoadMongoModel(URI, DBName string, colNames []string) error {
	m, err := model.NewModel(URI, DBName, colNames)
	if err != nil {
		return err
	}
	AppModel = m
	return nil
}

func InjectModelsMongoDependency(m map[string]*mongo.Collection) {

}

func CreateIndexesInModels() {

}
