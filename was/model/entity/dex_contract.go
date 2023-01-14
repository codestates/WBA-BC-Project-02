package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DexContract struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ContractAddress  string             `bson:"contract_addrss,omitempty"`
	TransactionHashs []string           `bson:"transaction_hashs,omitempty"`
	DracoPoolToken   string             `bson:"draco_pool_token,omitempty"`
	DracoPoolCredit  string             `bson:"draco_pool_credit"`
	TigPoolToken     string             `bson:"tig_pool_token"`
	TigPoolCredit    string             `bson:"tig_pool_token"`
}