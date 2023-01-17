package contract

import (
	wasError "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/protocol"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/contract/request"
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
	reqN := &request.ContractName{}
	if err := c.ShouldBindQuery(reqN); err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

	resC, err := co.contractService.GetContractByName(reqN.Name)
	if err != nil {
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(resC).Response(c)
}

func (co *contractControl) GetContracts(c *gin.Context) {
	simpleContracts, err := co.contractService.GetContracts()
	if err != nil {
		return
	}
	protocol.SuccessData(simpleContracts).Response(c)
}

func (co *contractControl) MintContract(c *gin.Context) {
	reqM := &request.MintingContract{}
	if err := c.ShouldBindJSON(reqM); err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

}

func (co *contractControl) ExchangeContract(c *gin.Context) {
	reqE := &request.ExchangeContract{}
	if err := c.ShouldBindJSON(reqE); err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

}

func (co *contractControl) GetRatioTokenAndCredit(c *gin.Context) {
	reqN := &request.ContractName{}
	if err := c.ShouldBindQuery(reqN); err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

	responseRatio, err := co.contractService.GetRatioTokenAndCredit(reqN.Name)
	if err != nil {
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(responseRatio).Response(c)
}
