package response

import "github.com/codestates/WBA-BC-Project-02/was/common/cache"

type SimpleUser struct {
	Address      string `json:"address"`
	BlackIron    int    `json:"black_iron"`
	WemixAmount  string `json:"wemix_amount"`
	DracoAmount  string `json:"draco_amount"`
	TigAmount    string `json:"tig_amount"`
	CreditAmount string `json:"credit_amount"`
}

func FromCache(info *cache.LoginInformation) *SimpleUser {
	return &SimpleUser{
		Address:      info.Address,
		BlackIron:    info.BlackIron,
		WemixAmount:  info.WemixAmount,
		DracoAmount:  info.DracoAmount,
		TigAmount:    info.TigAmount,
		CreditAmount: info.CreditAmount,
	}
}
