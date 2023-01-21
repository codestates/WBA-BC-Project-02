package entity

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DexContract struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	ContractAddress   string             `bson:"contract_address,omitempty"`
	TransactionHashes []string           `bson:"transaction_hashes,omitempty"`
	DracoPoolToken    string             `bson:"draco_pool_token,omitempty"`
	DracoPoolCredit   string             `bson:"draco_pool_credit"`
	TigPoolToken      string             `bson:"tig_pool_token"`
	TigPoolCredit     string             `bson:"tig_pool_credit"`
	BaseTime          *dom.BaseTime      `bson:"base_time,omitempty"`
}
