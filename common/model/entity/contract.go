package entity

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contract struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name,omitempty"`
	ContractAddress string             `bson:"contract_address,omitempty"`
	Transactions    []*dom.Transaction `bson:"transactions,omitempty"`
	BaseTime        *dom.BaseTime      `bson:"base_time,omitempty"`
}
