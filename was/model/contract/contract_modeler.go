package contract

import "github.com/codestates/WBA-BC-Project-02/common/model/entity"

type ContractModeler interface {
	FindContractByName(name string) (*entity.Contract, error)
	FindNonTxContracts() ([]*entity.Contract, error)
}
