package contract

import "github.com/codestates/WBA-BC-Project-02/common/model/entity"

type DexContractModeler interface {
	FindDexContract(dexAddr string) (*entity.DexContract, error)
}
