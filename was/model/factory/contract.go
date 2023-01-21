package factory

import (
	commonEnum "github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity/dom"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func NewInitDexContracts(dexContractAddr string) *entity.DexContract {
	return &entity.DexContract{
		ID:                primitive.NewObjectID(),
		ContractAddress:   dexContractAddr,
		TransactionHashes: []string{},
		DracoPoolToken:    commonEnum.ZeroSTR,
		DracoPoolCredit:   commonEnum.ZeroSTR,
		TigPoolToken:      commonEnum.ZeroSTR,
		TigPoolCredit:     commonEnum.ZeroSTR,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func NewInitDracoContracts(dracoAddr string) *entity.Contract {
	return &entity.Contract{
		ID:              primitive.NewObjectID(),
		Name:            enum.DracoContractName,
		ContractAddress: dracoAddr,
		Transactions:    []*dom.Transaction{},
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func NewInitCreditContracts(creditAddr string) *entity.Contract {
	return &entity.Contract{
		ID:              primitive.NewObjectID(),
		Name:            enum.CreditContractName,
		ContractAddress: creditAddr,
		Transactions:    []*dom.Transaction{},
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}

func NewInitTigContracts(tigAddr string) *entity.Contract {
	return &entity.Contract{
		ID:              primitive.NewObjectID(),
		Name:            enum.TigContractName,
		ContractAddress: tigAddr,
		Transactions:    []*dom.Transaction{},
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
