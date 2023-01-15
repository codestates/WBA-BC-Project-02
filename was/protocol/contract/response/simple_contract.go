package response

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
)

type SimpleContract struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	ContractAddress string `json:"contract_address"`
}

func FromSimpleContracts(cons []*entity.Contract) []*SimpleContract {
	sContracts := make([]*SimpleContract, len(cons))
	for _, con := range cons {
		c := FromSimpleContract(con)
		sContracts = append(sContracts, c)
	}
	return sContracts
}

func FromSimpleContract(con *entity.Contract) *SimpleContract {
	return &SimpleContract{
		ID:              con.ID.Hex(),
		Name:            con.Name,
		ContractAddress: con.ContractAddress,
	}
}
