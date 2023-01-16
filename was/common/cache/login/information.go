package login

import (
	"encoding/json"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
)

type Information struct {
	UserID       string `json:"user_id"`
	Device       string `json:"device"`
	Address      string `json:"address"`
	BlackIron    int    `json:"black_iron"`
	WemixAmount  string `json:"wemix_amount"`
	DracoAmount  string `json:"draco_amount"`
	TigAmount    string `json:"tig_amount"`
	CreditAmount string `json:"credit_amount"`
	TokenID      string `json:"token_id"`
}

func NewLoginInfo(device string, user *entity.User) *Information {
	return &Information{
		UserID:       user.ID.Hex(),
		Device:       device,
		Address:      user.Address,
		BlackIron:    user.BlackIron,
		WemixAmount:  user.WemixAmount,
		DracoAmount:  user.DracoAmount,
		TigAmount:    user.TigAmount,
		CreditAmount: user.CreditAmount,
	}
}

func (l *Information) MarshalBinary() (data []byte, err error) {
	return json.Marshal(l)
}
