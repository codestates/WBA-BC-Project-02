package middleware

import (
	"fmt"
	commonEnum "github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	error2 "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/protocol"
	"github.com/codestates/WBA-BC-Project-02/was/service/user"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("JWT :: middleware")
		headerAuth := c.GetHeader(enum.HeaderAuthorization)
		token := extractToken(headerAuth)

		userAgent := c.GetHeader(enum.HeaderUserAgent)
		loginInfo, err := user.ValidateTokenAndUserAgent(token, userAgent, enum.JWTAccessUUID, config.JWTConfig.AccessKey)
		if err != nil {
			fmt.Println(err)
			protocol.Fail(error2.NewAppError(err)).Response(c)
			c.AbortWithStatus(204)
			return
		}

		c.Set(enum.LoginInformation, loginInfo)
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
