package contract

import (
	"github.com/codestates/WBA-BC-Project-02/common/util/validator"
	wasError "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/protocol"
	"github.com/codestates/WBA-BC-Project-02/was/service/contract"
	"github.com/gin-gonic/gin"
)

var instance *contractControl

type contractControl struct {
	contractService contract.ContractServicer
}

func NewContractControl(contractService contract.ContractServicer) *contractControl {
	if instance != nil {
		return instance
	}
	instance = &contractControl{
		contractService: contractService,
	}
	return instance
}

func (co *contractControl) GetContractByName(c *gin.Context) {
	name := c.Query("name")
	if err := validator.CheckBlank(name); err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

	co.contractService.GetContractByName(name)
}
