package v1

import (
	"github.com/codestates/WBA-BC-Project-02/was/controller"
	"github.com/gin-gonic/gin"
)

// AuthContracts ("app/v1/auth/contracts")
func AuthContracts(authUrl *gin.RouterGroup) {
	authUrl.GET("/contract", controller.ContractControl.GetContractByName)
	authUrl.GET("", controller.ContractControl.GetContracts)

	authUrl.POST("/minting", controller.ContractControl.MintContract)
	authUrl.PUT("/exchange", controller.ContractControl.ExchangeContract)
	authUrl.GET("/ratio", controller.ContractControl.GetRatioTokenAndCredit)
}
