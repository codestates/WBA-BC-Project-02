package v1

import (
	"github.com/codestates/WBA-BC-Project-02/was/controller"
	"github.com/gin-gonic/gin"
)

// TODO jwt 리프래시 토큰을 통해 엑세스 토큰을 재발급 받는 컨트롤러도 필요 (핵심 기능이 아니기에 우선순위 일단 낮음)

// Users ("app/v1/users")
func Users(usersUrl *gin.RouterGroup) {
	usersUrl.POST("/wallet", controller.UserControl.CreateWallet)
	usersUrl.PUT("/wallet", controller.UserControl.RecoverWallet)
	usersUrl.GET("/tokens", controller.UserControl.ReissueToken)
}

// AuthUsers ("app/v1/auth/users")
func AuthUsers(authUrl *gin.RouterGroup) {
	authUrl.GET("", controller.UserControl.GetUserSimpleInformation)
	authUrl.GET("/transactions", controller.UserControl.GetUserInformation)
	authUrl.PUT("/black-iron", controller.UserControl.IncreaseBlackIron)
}
