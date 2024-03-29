package contract

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/login"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/contract/request"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/response"
)

type WemixonServicer interface {
	MintToken(loginInfo *login.Information, reqM *request.MintingContract) (*response.SimpleUser, error)

	ExchangeContract(loginInfo *login.Information, reqE *request.ExchangeContract) error
}
