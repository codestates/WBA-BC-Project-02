package user

import (
	"fmt"
	commonEnum "github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	wasError "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/config"
)

func ValidateTokenAndUserAgent(token, userAgent, uuidKey string) (*cache.LoginInformation, error) {
	jwtToken, err := cache.DecryptToken(token, config.JWTConfig.AccessKey)
	if err != nil {
		return nil, err
	}

	Key, tokenID, err := cache.ExtractTokenUUID(jwtToken, uuidKey)
	if err != nil {
		return nil, err
	}

	info, err := cache.GetLoginInfo(Key)
	if err != nil || info.UserID == commonEnum.BlankSTR {
		return nil, wasError.BadUserAgentError
	}

	fmt.Println("before device :: ", info.Device)
	fmt.Println("user Agent :: ", userAgent)
	if info.Device != userAgent {
		return nil, wasError.DifferentDeviceError
	}

	if info.TokenID != tokenID {
		return nil, wasError.UnauthorizedError
	}

	return info, nil
}
