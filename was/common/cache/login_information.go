package cache

import (
	"encoding/json"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
)

type LoginInformation struct {
	UserID       string
	Device       string
	Address      string
	WemixAmount  string
	DracoAmount  string
	TigAmount    string
	CreditAmount string
	PublicKey    string
	PrivateKey   string
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
		PublicKey:    user.PublicKey,
		PrivateKey:   user.PrivateKey,
	}
}

func (l *LoginInformation) MarshalBinary() (data []byte, err error) {
	return json.Marshal(l)
}
