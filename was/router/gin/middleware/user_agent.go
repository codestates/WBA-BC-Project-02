package middleware

import (
	enum2 "github.com/codestates/WBA-BC-Project-02/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	error2 "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/protocol"
	"github.com/gin-gonic/gin"
	"strings"
)

var (
	exploreList = []string{
		"Chrome",
		"Edge",
		"Firefox",
		"Opera",
		"Safari",
		"Samsung",
		"WebView",
	}
)

func UserAgent() gin.HandlerFunc {
	return func(c *gin.Context) {
		device := c.GetHeader(enum.HeaderUserAgent)
		if err := check(device); err != nil {
			protocol.Fail(error2.BadUserAgentError).Response(c)
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func check(ua string) error {
	if config.ServerConfig.Mode == enum2.DevMode {
		exploreList = append(exploreList, "Postman")
	}
	for _, exp := range exploreList {
		if strings.Contains(ua, exp) {
			return nil
		}
	}
	return error2.BadUserAgentError
}
