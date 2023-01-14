package user

import (
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/response"
)

type UserServicer interface {
	CreateWallet(PWD, device string) (*response.Mnemonic, error)
}
