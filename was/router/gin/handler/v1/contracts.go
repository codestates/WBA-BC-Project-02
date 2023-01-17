package v1

import (
	"github.com/codestates/WBA-BC-Project-02/was/controller"
	"github.com/gin-gonic/gin"
)

// AuthContracts ("app/v1/auth/contracts")
func AuthContracts(authUrl *gin.RouterGroup) {
	authUrl.GET("contract", controller.ContractControl.GetContractByName)
	authUrl.GET("", controller.ContractControl.GetContracts)
}
