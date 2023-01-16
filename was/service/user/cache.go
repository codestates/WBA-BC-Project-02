package user

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/login"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/token"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/logger"
	"time"
)

func saveCacheLoginInfos(loginInfo *login.Information, token *token.Tokens) error {
	if err := cacheLoginInfos(loginInfo, token); err != nil {
		return err
	}
	return nil
}

func updateAccessCacheInfo(loginInfo *login.Information) error {
	key := enum.AccessCache + loginInfo.UserID
	if err := cache.AppCacher.Update(key, loginInfo, -1); err != nil {
		return err
	}
	return nil
}

func deleteCachedLoginInfos(token *token.Tokens) {
	if err := cache.AppCacher.Delete(token.AccessToken.CacheID, token.RefreshToken.CacheID); err != nil {
		logger.AppLog.Error(err)
	}
}

func getToken(userID string) (*token.Tokens, error) {
	accessKey := config.JWTConfig.AccessKey
	refreshKey := config.JWTConfig.RefreshKey
	token, err := token.CreateTokens(userID, accessKey, refreshKey)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func cacheLoginInfos(loginInfo *login.Information, token *token.Tokens) error {
	loginInfo.TokenID = token.AccessToken.TokenID

	at := time.Unix(token.AccessToken.Duration, 0)
	rt := time.Unix(token.RefreshToken.Duration, 0)
	now := time.Now()

	loginInfo.TokenID = token.AccessToken.TokenID
	if err := cache.AppCacher.Cache(token.AccessToken.CacheID, loginInfo, at.Sub(now)); err != nil {
		return err
	}

	loginInfo.TokenID = token.RefreshToken.TokenID
	if err := cache.AppCacher.Cache(token.RefreshToken.CacheID, loginInfo, rt.Sub(now)); err != nil {
		return err
	}

	return nil
}

func getLoginInfo(key string) (*login.Information, error) {
	info := &login.Information{}
	if _, err := cache.AppCacher.Get(key, info); err != nil {
		return nil, err
	}

	return info, nil
}
