package contract

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/login"
	"github.com/codestates/WBA-BC-Project-02/was/common/contract/util"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/contract/request"
	userService "github.com/codestates/WBA-BC-Project-02/was/service/user"
)

func (w *wemixonService) ExchangeContract(loginInfo *login.Information, reqE *request.ExchangeContract) error {
	userNonTx, err := w.userModel.FindUserNonTx(loginInfo.Address)
	if err != nil {
		return err
	}

	if err := userService.BcryptVerifyPassword(userNonTx.Password, reqE.Password); err != nil {
		return err
	}

	switch {
	case reqE.From == enum.DracoContractName && reqE.To == enum.CreditContractName:
		if _, err := ActContract(loginInfo.Address, reqE.Amount, util.BuyCreditbyDracoTx); err != nil {
			return err
		}
	case reqE.From == enum.CreditContractName && reqE.To == enum.DracoContractName:
		if _, err := ActContract(loginInfo.Address, reqE.Amount, util.BuyDracoByCreditTx); err != nil {
			return err
		}
	case reqE.From == enum.TigContractName && reqE.To == enum.CreditContractName:
		if _, err := ActContract(loginInfo.Address, reqE.Amount, util.BuyCreditByTigTx); err != nil {
			return err
		}
	case reqE.From == enum.CreditContractName && reqE.To == enum.TigContractName:
		if _, err := ActContract(loginInfo.Address, reqE.Amount, util.BuyTigByCreditTx); err != nil {
			return err
		}
	case reqE.From == enum.CreditContractName && reqE.To == enum.WemixName:
		if _, err := ActContract(loginInfo.Address, reqE.Amount, util.CreditToWemixTx); err != nil {
			return err
		}
	}

	return nil
}
