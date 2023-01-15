package user

import (
	"github.com/codestates/WBA-BC-Project-02/common/ciper"
	"github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/model/factory"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/response"
	"strings"
)

func (u *userService) CreateWallet(PWD, userAgent string) (*response.Mnemonic, error) {
	mnemonic, err := NewMnemonic()
	if err != nil {
		return nil, err
	}

	wallet, err := NewWallet(mnemonic)
	if err != nil {
		return nil, err
	}

	hashPassword := BcryptHashPassword(PWD)

	encryptPK, err := ciper.AESEncrypt(ciper.GetCipherBlock(), []byte(wallet.PrivateKey))
	if err != nil {
		return nil, err
	}

	newUser := factory.NewCreateUser(hashPassword, wallet.Address, encryptPK, wallet.PublicKey)

	token, err := u.getToken(newUser)
	if err != nil {
		return nil, err
	}

	if err := u.saveCache(newUser, userAgent, token); err != nil {
		return nil, err
	}

	if err := u.userModel.InsertUser(newUser); err != nil {
		u.deleteCachedToken(token)
		return nil, err
	}

	slicedMnemonic := strings.Split(mnemonic, enum.SpaceSTR)
	return &response.Mnemonic{
		Mnemonic: slicedMnemonic,
		Token: response.Token{
			AccessToken:  token.AccessToken.Token,
			RefreshToken: token.RefreshToken.Token,
		},
	}, nil
}
