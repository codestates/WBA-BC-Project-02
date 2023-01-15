package user

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	error2 "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/config"
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

func (u *userService) ReissueToken(refreshToken string, ua string) (*response.Token, error) {
	info, err := ValidateTokenAndUserAgent(refreshToken, ua, enum.JWTRefreshUUID, config.JWTConfig.RefreshKey)
	if err != nil {
		return nil, err
	}

	tokens, err := getToken(info.UserID)
	if err != nil {
		return nil, err
	}

	deleteCachedLoginInfos(tokens)

	info.Device = ua
	if err := saveCacheLoginInfos(info, tokens); err != nil {
		return nil, err
	}

	return &response.Token{
		AccessToken:  tokens.AccessToken.Token,
		RefreshToken: tokens.RefreshToken.Token,
	}, nil
}

func (u *userService) GetUser(address string) (*response.User, error) {
	foundUser, err := u.userModel.FindUser(address)
	if err != nil {
		return nil, error2.UserNotFoundError
	}
	resU := response.FromUserEntity(foundUser)
	return resU, err
}

func (u *userService) IncreaseBlackIron(info *cache.LoginInformation) (*response.SimpleUser, error) {
	updatedUser, err := u.userModel.FindUserAndIncreaseIron(info.Address)
	if err != nil {
		return nil, err
	}

	info.BlackIron = updatedUser.BlackIron
	if err := updateAccessCacheInfo(info); err != nil {
		return nil, err
	}

	return response.FromCache(info), nil
}
