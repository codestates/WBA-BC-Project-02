package util

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/logger"
)

func SaveCacheLoginInfos(loginInfo *cache.LoginInformation, token *cache.Tokens) error {
	if err := cache.CacheLoginInfos(loginInfo, token); err != nil {
		return err
	}
	return nil
}

func DeleteCachedLoginInfos(token *cache.Tokens) {
	if err := cache.Delete(token.AccessToken.CacheID, token.RefreshToken.CacheID); err != nil {
		logger.AppLog.Error(err)
	}
}

func GetToken(userID string) (*cache.Tokens, error) {
	accessKey := config.JWTConfig.AccessKey
	refreshKey := config.JWTConfig.RefreshKey
	token, err := cache.CreateToken(userID, accessKey, refreshKey)
	if err != nil {
		return nil, err
	}

	return token, nil
}
