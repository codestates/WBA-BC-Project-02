package user

import (
	commonEnum "github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/login"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/token"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	wasError "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/logger"
	"strings"
)

func ValidateTokenAndUserAgent(JWTToken, userAgent, JWTFiled, JWTCipherKey string) (*login.Information, error) {
	jwtToken, err := token.DecryptToken(JWTToken, JWTCipherKey)
	if err != nil {
		return nil, err
	}

	Key, tokenID, err := token.ExtractCacheIdAndFiled(jwtToken, JWTFiled)
	if err != nil {
		return nil, err
	}

	info, err := getLoginInfo(Key)
	if err != nil || info.UserID == commonEnum.BlankSTR {
		return nil, wasError.UnauthorizedError
	}

	if info.Device != userAgent {
		checkAccessCacheThenDelete(Key)
		return nil, wasError.DifferentDeviceError
	}

	if info.TokenID != tokenID {
		checkAccessCacheThenDelete(Key)
		return nil, wasError.UnauthorizedError
	}

	return info, nil
}

func checkAccessCacheThenDelete(Key string) {
	if strings.Contains(Key, enum.AccessCache) {
		if err := cache.Redis.Delete(Key); err != nil {
			logger.AppLog.Error(err)
		}
	}
}
