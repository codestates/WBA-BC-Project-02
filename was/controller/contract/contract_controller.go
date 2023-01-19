package contract

import "github.com/gin-gonic/gin"

type ContractController interface {
	GetContractByName(c *gin.Context)

	GetContracts(c *gin.Context)

	MintContract(c *gin.Context)

	ExchangeContract(c *gin.Context)

	GetRatioTokenAndCredit(c *gin.Context)

	GetNonce(c *gin.Context)
}
