package cache

import (
	"encoding/json"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
)

type LoginInformation struct {
	UserID       string `json:"user_id"`
	Device       string `json:"device"`
	Address      string `json:"address"`
	WemixAmount  string `json:"wemix_amount"`
	DracoAmount  string `json:"draco_amount"`
	TigAmount    string `json:"tig_amount"`
	CreditAmount string `json:"credit_amount"`
	TokenID      string `json:"token_id"`
}

func NewLoginInfo(device string, user *entity.User) *LoginInformation {
	return &LoginInformation{
		UserID:       user.ID.Hex(),
		Device:       device,
		Address:      user.Address,
		WemixAmount:  user.WemixAmount,
		DracoAmount:  user.DracoAmount,
		TigAmount:    user.TigAmount,
		CreditAmount: user.CreditAmount,
	}
}

func (l *LoginInformation) MarshalBinary() (data []byte, err error) {
	return json.Marshal(l)
}
