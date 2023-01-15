package user

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	error2 "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/model/user"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/response"
	"github.com/codestates/WBA-BC-Project-02/was/service/user/util"
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
	info, err := util.ValidateTokenAndUserAgent(refreshToken, ua, enum.JWTRefreshUUID, config.JWTConfig.RefreshKey)
	if err != nil {
		return nil, err
	}

	tokens, err := util.GetToken(info.UserID)
	if err != nil {
		return nil, err
	}

	util.DeleteCachedLoginInfos(tokens)

	info.Device = ua
	if err := util.SaveCacheLoginInfos(info, tokens); err != nil {
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
	resU := response.FromEntity(foundUser)
	return resU, err
}
