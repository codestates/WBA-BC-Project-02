package user

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/logger"
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

	tokens, err := u.getToken(info.UserID)
	if err != nil {
		return nil, err
	}

	u.deleteCachedLoginInfos(tokens)

	if err := u.saveCacheLoginInfos(info, ua, tokens); err != nil {
		return nil, err
	}

	return &response.Token{
		AccessToken:  tokens.AccessToken.Token,
		RefreshToken: tokens.RefreshToken.Token,
	}, nil
}

func (u *userService) getToken(userID string) (*cache.Tokens, error) {
	accessKey := config.JWTConfig.AccessKey
	refreshKey := config.JWTConfig.RefreshKey
	token, err := cache.CreateToken(userID, accessKey, refreshKey)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (u *userService) saveCacheLoginInfos(loginInfo *cache.LoginInformation, userAgent string, token *cache.Tokens) error {
	if err := cache.CacheLoginInfos(loginInfo, token); err != nil {
		return err
	}
	return nil
}

func (u *userService) deleteCachedLoginInfos(token *cache.Tokens) {
	if err := cache.Delete(token.AccessToken.CacheID, token.RefreshToken.CacheID); err != nil {
		logger.AppLog.Error(err)
	}
}
