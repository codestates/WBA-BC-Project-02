package contract

import "github.com/gin-gonic/gin"

type ContractController interface {
	GetContractByName(c *gin.Context)

	GetContracts(c *gin.Context)
}
