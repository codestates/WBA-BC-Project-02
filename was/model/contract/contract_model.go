package contract

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/model/query"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var instance *contractModel

type contractModel struct {
	collection *mongo.Collection
}

func NewContractModel(col *mongo.Collection) *contractModel {
	if instance != nil {
		return instance
	}
	instance = &contractModel{
		collection: col,
	}
	return instance
}

func (c *contractModel) FindContractByName(name string) (*entity.Contract, error) {
	filter := query.GetNameFilter(name)

	con := &entity.Contract{}
	if err := query.NewFindAction(
		con, c.collection,
	).InjectFilter(filter).FindOne(nil); err != nil {
		return nil, err
	}

	return con, nil
}

func (c *contractModel) FindNonTxContracts() ([]*entity.Contract, error) {
	opt := options.Find().SetProjection(bson.M{enum.Transactions: 0})

	var contracts []*entity.Contract
	if err := query.NewFindAction(
		contracts, c.collection,
	).Find(opt); err != nil {
		return nil, err
	}

	return contracts, nil
}
