package contract

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/login"
	"github.com/codestates/WBA-BC-Project-02/was/common/contract/util"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/contract/request"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/response"
	userService "github.com/codestates/WBA-BC-Project-02/was/service/user"
)

func (w *wemixonService) ExchangeContract(loginInfo *login.Information, reqE *request.ExchangeContract) (*response.SimpleUser, error) {
	userNonTx, err := w.userModel.FindUserNonTx(loginInfo.Address)
	if err != nil {
		return nil, err
	}

	if err := userService.BcryptVerifyPassword(userNonTx.Password, reqE.Password); err != nil {
		return nil, err
	}

	switch {
	case reqE.From == enum.DracoContractName && reqE.To == enum.CreditContractName:
		if _, err := ActContract(loginInfo.Address, reqE.Amount, util.BuyCreditbyDracoTx); err != nil {
			return nil, err
		}
	case reqE.From == enum.CreditContractName && reqE.To == enum.DracoContractName:
		if _, err := ActContract(loginInfo.Address, reqE.Amount, util.BuyDracoByCreditTx); err != nil {
			return nil, err
		}
	case reqE.From == enum.TigContractName && reqE.To == enum.CreditContractName:
		if _, err := ActContract(loginInfo.Address, reqE.Amount, util.BuyCreditByTigTx); err != nil {
			return nil, err
		}
	case reqE.From == enum.CreditContractName && reqE.To == enum.TigContractName:
		if _, err := ActContract(loginInfo.Address, reqE.Amount, util.BuyTigByCreditTx); err != nil {
			return nil, err
		}
	case reqE.From == enum.CreditContractName && reqE.To == enum.WemixName:
		if _, err := ActContract(loginInfo.Address, reqE.Amount, util.CreditToWemixTx); err != nil {
			return nil, err
		}
	}
	// TODO 가지고 있는 교환후 감소는 어떻게???
	// TODO Redis 와 DB 에 저장 전 loginInfo 감소 동기화 해줘야함
	if _, err := w.userModel.FindUserAndSet(loginInfo); err != nil {
		return nil, err
	}
	if err := userService.UpdateAccessCacheInfo(loginInfo); err != nil {
		return nil, err
	}

	return response.FromCache(loginInfo), nil
}
