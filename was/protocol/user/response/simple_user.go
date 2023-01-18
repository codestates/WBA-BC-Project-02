package response

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/common/util/convertor"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/login"
)

type SimpleUser struct {
	Address      string `json:"address"`
	BlackIron    int    `json:"black_iron"`
	WemixAmount  string `json:"wemix_amount"`
	DracoAmount  string `json:"draco_amount"`
	TigAmount    string `json:"tig_amount"`
	CreditAmount string `json:"credit_amount"`
}

func FromCache(info *login.Information) *SimpleUser {
	return &SimpleUser{
		Address:      info.Address,
		BlackIron:    info.BlackIron,
		WemixAmount:  convertor.DefaultAmount(info.WemixAmount),
		DracoAmount:  convertor.DefaultAmount(info.DracoAmount),
		TigAmount:    convertor.DefaultAmount(info.TigAmount),
		CreditAmount: convertor.DefaultAmount(info.CreditAmount),
	}
}

func GetSimpleFromUser(usr *entity.User) *SimpleUser {
	return &SimpleUser{
		Address:      usr.Address,
		BlackIron:    usr.BlackIron,
		WemixAmount:  convertor.DefaultAmount(usr.WemixAmount),
		DracoAmount:  convertor.DefaultAmount(usr.DracoAmount),
		TigAmount:    convertor.DefaultAmount(usr.TigAmount),
		CreditAmount: convertor.DefaultAmount(usr.CreditAmount),
	}
}
