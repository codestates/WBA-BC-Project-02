package contract

import "github.com/codestates/WBA-BC-Project-02/was/protocol/contract/response"

type ContractServicer interface {
	GetContractByName(name string) (*response.Contract, error)
	GetContracts() ([]*response.SimpleContract, error)
}
