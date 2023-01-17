package response

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/was/protocol"
)

type Contract struct {
	ID              string                  `json:"id"`
	Name            string                  `json:"name"`
	ContractAddress string                  `json:"contract_address"`
	Transactions    []*protocol.Transaction `json:"transactions"`
}

func FromContractEntity(con *entity.Contract) *Contract {
	return &Contract{
		ID:              con.ID.Hex(),
		Name:            con.Name,
		ContractAddress: con.ContractAddress,
		Transactions:    protocol.FromTransactionDoms(con.Transactions),
	}
}
