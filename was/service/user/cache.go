package user

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/logger"
)

func saveCacheLoginInfos(loginInfo *cache.LoginInformation, token *cache.Tokens) error {
	if err := cache.CacheLoginInfos(loginInfo, token); err != nil {
		return err
	}
	return nil
}

func updateAccessCacheInfo(loginInfo *cache.LoginInformation) error {
	key := enum.AccessCache + loginInfo.UserID
	if err := cache.UpdateLoginInfo(key, loginInfo, -1); err != nil {
		return err
	}
	return nil
}

func deleteCachedLoginInfos(token *cache.Tokens) {
	if err := cache.Delete(token.AccessToken.CacheID, token.RefreshToken.CacheID); err != nil {
		logger.AppLog.Error(err)
	}
}

func getToken(userID string) (*cache.Tokens, error) {
	accessKey := config.JWTConfig.AccessKey
	refreshKey := config.JWTConfig.RefreshKey
	token, err := cache.CreateToken(userID, accessKey, refreshKey)
	if err != nil {
		return nil, err
	}

	return token, nil
}
