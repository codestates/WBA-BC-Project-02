package response

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity/dom"
	"github.com/codestates/WBA-BC-Project-02/common/util/convertor"
)

type User struct {
	ID           string             `json:"user_id"`
	BlackIron    int                `json:"black_iron"`
	WemixAmount  string             `json:"wemix_amount"`
	DracoAmount  string             `json:"draco_amount"`
	TigAmount    string             `json:"tig_amount"`
	CreditAmount string             `json:"credit_amount"`
	Address      string             `json:"address"`
	Transactions []*dom.Transaction `json:"transactions"`
}

func FromUserEntity(usr *entity.User) *User {
	return &User{
		ID:           usr.ID.Hex(),
		BlackIron:    usr.BlackIron,
		WemixAmount:  convertor.DefaultAmount(usr.WemixAmount),
		DracoAmount:  convertor.DefaultAmount(usr.DracoAmount),
		TigAmount:    convertor.DefaultAmount(usr.TigAmount),
		CreditAmount: convertor.DefaultAmount(usr.CreditAmount),
		Address:      usr.Address,
		Transactions: usr.Transactions,
	}
}
