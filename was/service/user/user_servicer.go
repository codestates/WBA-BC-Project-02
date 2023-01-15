package user

import (
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/request"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/response"
)

type UserServicer interface {
	CreateWallet(PWD, userAgent string) (*response.Mnemonic, error)
	RecoverWallet(recovery *request.Recovery, userAgent string) (*response.Token, error)
}
