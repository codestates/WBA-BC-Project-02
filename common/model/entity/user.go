package entity

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Password     string             `bson:"password,omitempty"`
	BlackIron    int                `bson:"black_iron,omitempty"`
	WemixAmount  string             `bson:"wemix_amount,omitempty"`
	DracoAmount  string             `bson:"draco_amount,omitempty"`
	TigAmount    string             `bson:"tig_amount,omitempty"`
	CreditAmount string             `bson:"credit_amount,omitempty"`
	PrivateKey   string             `bson:"private_key,omitempty"`
	PublicKey    string             `bson:"public_key,omitempty"`
	Address      string             `bson:"address,omitempty"`
	Transactions []*dom.Transaction `bson:"transactions,omitempty"`
	BaseTime     *dom.BaseTime      `bson:"base_time,omitempty"`
}
