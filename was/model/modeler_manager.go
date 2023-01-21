package model

import (
	"github.com/codestates/WBA-BC-Project-02/common/enum"
	wasEnum "github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/model/contract"
	"github.com/codestates/WBA-BC-Project-02/was/model/factory"
	"github.com/codestates/WBA-BC-Project-02/was/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var AppModel Modeler

var UserModel user.UserModeler
var ContractModel contract.ContractModeler
var DexContractModel contract.DexContractModeler

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
	DexContractModel = contract.NewDexContractModel(m[enum.DexContractName])
}

func CreateIndexesInModels() {
	AppModel.CreateIndexes(enum.UserCollectionName, true, wasEnum.PrivateKey, wasEnum.Address)
	AppModel.CreateIndexes(enum.ContractCollectionName, true, wasEnum.Name)
	AppModel.CreateIndexes(enum.DexContractName, true, wasEnum.ContractAddress)
}

func InsertInitContracts() {
	if _, err := DexContractModel.FindDexContractNonTxHashes(config.ContractConfig.DexAddr); err != nil {
		dexCon := factory.NewInitDexContracts(config.ContractConfig.DexAddr)
		DexContractModel.InsertOne(dexCon)
	}

	if _, err := ContractModel.FindContractByName(enum.DracoContract); err != nil {
		dracoCon := factory.NewInitDracoContracts(config.ContractConfig.DracoAddr)
		ContractModel.InsertOne(dracoCon)
	}

	if _, err := ContractModel.FindContractByName(enum.CreditContract); err != nil {
		creditCon := factory.NewInitCreditContracts(config.ContractConfig.CreditAddr)
		ContractModel.InsertOne(creditCon)
	}

	if _, err := ContractModel.FindContractByName(enum.TigContract); err != nil {
		tigCon := factory.NewInitTigContracts(config.ContractConfig.TigAddr)
		ContractModel.InsertOne(tigCon)
	}
}
