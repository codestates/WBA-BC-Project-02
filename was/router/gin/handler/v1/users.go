package v1

import (
	"github.com/codestates/WBA-BC-Project-02/was/controller"
	"github.com/gin-gonic/gin"
)

// Users ("app/v1/users")
func Users(usersUrl *gin.RouterGroup) {
	usersUrl.POST("/wallet", controller.UserControl.CreateWallet)
	usersUrl.POST("/wallet/recovery", controller.UserControl.RecoverWallet)
}
