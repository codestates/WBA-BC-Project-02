package user

import (
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/logger"
	"github.com/codestates/WBA-BC-Project-02/was/model/user"
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

func (u *userService) getToken(newUser *entity.User) (*cache.Token, error) {
	accessKey := config.JWTConfig.AccessKey
	refreshKey := config.JWTConfig.RefreshKey
	token, err := cache.CreateToken(newUser.ID.Hex(), accessKey, refreshKey)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (u *userService) saveCache(user *entity.User, userAgent string, token *cache.Token) error {
	loginInfo := cache.NewLoginInfo(userAgent, user)
	if err := cache.CacheLoginInfo(loginInfo, token); err != nil {
		return err
	}
	return nil
}

func (u *userService) deleteCachedToken(token *cache.Token) {
	if err := cache.Delete(token.AccessToken.ID, token.RefreshToken.ID); err != nil {
		logger.AppLog.Error(err)
	}
}
