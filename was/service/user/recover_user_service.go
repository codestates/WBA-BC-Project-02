package user

import (
	"github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/request"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/response"
	"strings"
)

func (u *userService) RecoverWallet(recovery *request.Recovery, userAgent string) (*response.Token, error) {
	mnemonicSTR := strings.Join(recovery.Mnemonic, enum.SpaceSTR)
	wallet, err := NewWallet(mnemonicSTR)
	if err != nil {
		return nil, err
	}

	encryptedPWD := BcryptHashPassword(recovery.Password)

	user, err := u.userModel.FindUserAndPWDUpdate(wallet.Address, encryptedPWD)
	if err != nil {
		return nil, err
	}

	tokens, err := u.getToken(user.ID.Hex())
	if err != nil {
		return nil, err
	}

	u.deleteCachedLoginInfos(tokens)

	loginInfo := cache.NewLoginInfo(userAgent, user)
	if err := u.saveCacheLoginInfos(loginInfo, userAgent, tokens); err != nil {
		return nil, err
	}

	return &response.Token{
		AccessToken:  tokens.AccessToken.Token,
		RefreshToken: tokens.RefreshToken.Token,
	}, nil
}
