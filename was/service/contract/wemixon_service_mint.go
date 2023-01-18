package contract

import (
	"errors"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/login"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/contract/request"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/response"
	userService "github.com/codestates/WBA-BC-Project-02/was/service/user"
)

func (w *wemixonService) MintToken(loginInfo *login.Information, reqM *request.MintingContract) (*response.SimpleUser, error) {
	userNonTx, err := w.userModel.FindUserNonTx(loginInfo.Address)
	if err != nil {
		return nil, err
	}

	if err := userService.BcryptVerifyPassword(userNonTx.Password, reqM.Password); err != nil {
		return nil, err
	}

	if loginInfo.BlackIron < reqM.BlackIronBurnAmount {
		// TODO ERROR
		return nil, errors.New("충분한 흑철이 없습니다")
	}

	switch reqM.MintingName {
	case enum.DracoContractName:
		if _, err := mintDraco(loginInfo.Address, reqM.BlackIronBurnAmount); err != nil {
			return nil, err
		}
	case enum.TigContractName:
		if _, err := mintTig(loginInfo.Address, reqM.BlackIronBurnAmount); err != nil {
			return nil, err
		}
	}

	// 흑철 감소
	updateIron := loginInfo.BlackIron - reqM.BlackIronBurnAmount
	updatedUser, err := w.userModel.FindUserAndSetIron(loginInfo.Address, updateIron)
	if err != nil {
		return nil, err
	}

	loginInfo.BlackIron = updatedUser.BlackIron
	if err := userService.UpdateAccessCacheInfo(loginInfo); err != nil {
		return nil, err
	}

	return response.FromCache(loginInfo), nil
}
