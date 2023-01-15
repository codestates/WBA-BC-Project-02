package user

import (
	"github.com/codestates/WBA-BC-Project-02/common/ciper"
	"github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/model/factory"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/response"
	"github.com/codestates/WBA-BC-Project-02/was/service/user/util"
	"strings"
)

func (u *userService) CreateWallet(PWD, userAgent string) (*response.Mnemonic, error) {
	mnemonic, err := util.NewMnemonic()
	if err != nil {
		return nil, err
	}

	wallet, err := util.NewWallet(mnemonic)
	if err != nil {
		return nil, err
	}

	hashPassword := util.BcryptHashPassword(PWD)

	encryptPK, err := ciper.AESEncrypt(ciper.GetCipherBlock(), []byte(wallet.PrivateKey))
	if err != nil {
		return nil, err
	}

	newUser := factory.NewCreateUser(hashPassword, wallet.Address, encryptPK, wallet.PublicKey)

	tokens, err := util.GetToken(newUser.ID.Hex())
	if err != nil {
		return nil, err
	}

	loginInfo := cache.NewLoginInfo(userAgent, newUser)
	if err := util.SaveCacheLoginInfos(loginInfo, tokens); err != nil {
		return nil, err
	}

	if err := u.userModel.InsertUser(newUser); err != nil {
		util.DeleteCachedLoginInfos(tokens)
		return nil, err
	}

	slicedMnemonic := strings.Split(mnemonic, enum.SpaceSTR)
	return &response.Mnemonic{
		Mnemonic: slicedMnemonic,
		Token: response.Token{
			AccessToken:  tokens.AccessToken.Token,
			RefreshToken: tokens.RefreshToken.Token,
		},
	}, nil
}
