package contract

import (
	"github.com/codestates/WBA-BC-Project-02/was/protocol/contract/response"
)

type ContractServicer interface {
	GetContractByName(contractName string) (*response.Contract, error)

	GetContracts() ([]*response.SimpleContract, error)

	GetRatioTokenAndCredit(contractName string) (*response.RatioContract, error)
}
