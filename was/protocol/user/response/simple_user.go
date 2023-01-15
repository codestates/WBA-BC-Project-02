package response

import "github.com/codestates/WBA-BC-Project-02/was/common/cache"

type SimpleUser struct {
	Address      string `json:"address"`
	WemixAmount  string `json:"wemix_amount"`
	DracoAmount  string `json:"draco_amount"`
	TigAmount    string `json:"tig_amount"`
	CreditAmount string `json:"credit_amount"`
}

func FromCache(info *cache.LoginInformation) *SimpleUser {
	return &SimpleUser{
		Address:      info.Address,
		WemixAmount:  info.WemixAmount,
		DracoAmount:  info.DracoAmount,
		TigAmount:    info.TigAmount,
		CreditAmount: info.CreditAmount,
	}
}
