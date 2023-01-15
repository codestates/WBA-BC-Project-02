package middleware

import (
	"fmt"
	commonEnum "github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	wasError "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/logger"
	"github.com/codestates/WBA-BC-Project-02/was/protocol"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerAuth := c.GetHeader(enum.HeaderAuthorization)
		token := extractToken(headerAuth)

		jwtToken, err := cache.DecryptToken(token, config.JWTConfig.AccessKey)
		if err != nil {
			logger.AppLog.Error(err)
			protocol.Fail(wasError.BadUserAgentError).Response(c)
			c.AbortWithStatus(204)
			return
		}

		accessKey, err := cache.ExtractAccessTokenUUID(jwtToken)
		if err != nil {
			logger.AppLog.Error(err)
			protocol.Fail(wasError.BadUserAgentError).Response(c)
			c.AbortWithStatus(204)
		}

		info, err := cache.GetLoginInfo(accessKey)
		if err != nil || info.UserID == commonEnum.BlankSTR {
			logger.AppLog.Error(err)
			protocol.Fail(wasError.BadUserAgentError).Response(c)
			c.AbortWithStatus(204)
			return
		}

		userAgent := c.GetHeader(enum.HeaderUserAgent)
		fmt.Println("before device :: ", info.Device)
		fmt.Println("user Agent :: ", userAgent)
		if info.Device != userAgent {
			protocol.Fail(wasError.DifferentDeviceError).Response(c)
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func extractToken(authorization string) string {
	strArr := strings.Split(authorization, commonEnum.SpaceSTR)
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
