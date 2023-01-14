package entity

import "github.com/codestates/WBA-BC-Project-02/was/model/entity/dom"

type Contract struct {
	ID              string             `bson:"_id,omitempty"`
	Name            string             `bson:"name,omitempty"`
	ContractAddress string             `bson:"contract_address,omitempty"`
	Transactions    []*dom.Transaction `bson:"transactions,omitempty"`
}
