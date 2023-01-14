package v1

import (
	"github.com/codestates/WBA-BC-Project-02/was/controller"
	"github.com/gin-gonic/gin"
)

// Info ("app/v1/info")
func Info(infoUrl *gin.RouterGroup) {
	infoUrl.GET("", controller.InfoControl.GetInformation)
}
