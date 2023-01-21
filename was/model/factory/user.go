package factory

import (
	"github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity/dom"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func NewCreateUser(password, address, priKey, pubKey string) *entity.User {
	return &entity.User{
		ID:           primitive.NewObjectID(),
		Password:     password,
		WemixAmount:  enum.ZeroSTR,
		DracoAmount:  enum.ZeroSTR,
		TigAmount:    enum.ZeroSTR,
		CreditAmount: enum.ZeroSTR,
		PublicKey:    priKey,
		PrivateKey:   pubKey,
		Address:      address,
		BaseTime: &dom.BaseTime{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
}
