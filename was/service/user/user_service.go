package user

import (
	"github.com/codestates/WBA-BC-Project-02/was/common"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/model/factory"
	"github.com/codestates/WBA-BC-Project-02/was/model/user"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/response"
)

type userService struct {
	userModel user.UserModeler
}

var instance *userService

func NewUserService(modeler user.UserModeler) *userService {
	if instance != nil {
		return instance
	}
	instance = &userService{
		userModel: modeler,
	}
	return instance
}

func (u *userService) CreateWallet(PWD, device string) (*response.Mnemonic, error) {
	mnemonic, err := NewMnemonic()
	if err != nil {
		return nil, err
	}

	wallet, err := NewWallet(mnemonic)
	if err != nil {
		return nil, err
	}

	hashPassword := HashPassword(PWD)

	newUser := factory.NewCreateUser(hashPassword, wallet.Address, wallet.PrivateKey, wallet.PublicKey)

	if err := u.userModel.InsertUser(newUser); err != nil {
		return nil, err
	}

	accessKey := config.JWTConfig.AccessKey
	refreshKey := config.JWTConfig.RefreshKey
	token, err := common.CreateToken(newUser.ID.Hex(), accessKey, refreshKey)
	if err != nil {
		return nil, err
	}

	loginInfo := cache.NewLoginInfo(device, newUser)
	if err := cache.CacheLoginInfo(loginInfo, token); err != nil {
		return nil, err
	}

	return &response.Mnemonic{
		Mnemonic: mnemonic,
		Token: response.Token{
			AccessToken:  token.AccessToken.Token,
			RefreshToken: token.RefreshToken.Token,
		},
	}, nil
}
