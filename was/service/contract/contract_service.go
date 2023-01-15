package contract

import (
	"github.com/codestates/WBA-BC-Project-02/was/model/contract"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/contract/response"
)

type contractService struct {
	contactModel contract.ContractModeler
}

var instance *contractService

func NewContractService(con contract.ContractModeler) *contractService {
	if instance != nil {
		return instance
	}
	instance = &contractService{
		contactModel: con,
	}
	return instance
}

func (c *contractService) GetContractByName(name string) (*response.Contract, error) {
	foundCon, err := c.contactModel.FindContractByName(name)
	if err != nil {
		return nil, err
	}

	resC := response.FromContractEntity(foundCon)
	return resC, nil
}

func (c *contractService) GetContracts() ([]*response.SimpleContract, error) {
	foundsContracts, err := c.contactModel.FindNonTxContracts()
	if err != nil {
		return nil, err
	}
	simpleContracts := response.FromSimpleContracts(foundsContracts)
	return simpleContracts, nil
}
