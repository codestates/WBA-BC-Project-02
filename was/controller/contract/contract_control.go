package contract

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	cacheContract "github.com/codestates/WBA-BC-Project-02/was/common/cache/contract"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/login"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	wasError "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/protocol"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/contract/request"
	"github.com/codestates/WBA-BC-Project-02/was/service/contract"
	"github.com/gin-gonic/gin"
	"net/http"
)

var instance *contractControl

type contractControl struct {
	contractService contract.ContractServicer
	wemixonService  contract.WemixonServicer
}

func NewContractControl(
	contractService contract.ContractServicer,
	wemixonService contract.WemixonServicer,
) *contractControl {
	if instance != nil {
		return instance
	}
	instance = &contractControl{
		contractService: contractService,
		wemixonService:  wemixonService,
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
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}
	protocol.SuccessData(simpleContracts).Response(c)
}

func (co *contractControl) MintContract(c *gin.Context) {
	loginInfo, exists := c.Keys[enum.LoginInformation].(*login.Information)
	if !exists {
		protocol.Fail(wasError.InternalServerError).Response(c)
		return
	}

	reqM := &request.MintingContract{}
	if err := c.ShouldBindJSON(reqM); err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

	simpleUser, err := co.wemixonService.MintToken(loginInfo, reqM)
	if err != nil {
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(simpleUser).Response(c)
}

func (co *contractControl) ExchangeContract(c *gin.Context) {
	loginInfo, exists := c.Keys[enum.LoginInformation].(*login.Information)
	if !exists {
		protocol.Fail(wasError.InternalServerError).Response(c)
		return
	}

	reqE := &request.ExchangeContract{}
	if err := c.ShouldBindJSON(reqE); err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

	if err := co.wemixonService.ExchangeContract(loginInfo, reqE); err != nil {
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessCodeAndData(http.StatusOK, gin.H{"exchange": "ok"}).Response(c)
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

func (co *contractControl) GetNonce(c *gin.Context) {
	reqM := &request.Nonce{}
	if err := c.ShouldBindQuery(reqM); err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

	nonce := &cacheContract.Nonce{}
	_, err := cache.Redis.Get(enum.NonceCacheKey, nonce)
	if err != nil {
		protocol.SuccessCodeAndData(wasError.DataNotFoundCode, gin.H{"nonce": false}).Response(c)
		return
	}

	for _, n := range nonce.NonceValues {
		if n == uint64(reqM.Value) {
			cache.Redis.Delete(enum.NonceCacheKey)
			protocol.SuccessCodeAndData(http.StatusOK, gin.H{"nonce": true}).Response(c)
			return
		}
	}

	protocol.SuccessCodeAndData(wasError.DataNotFoundCode, gin.H{"nonce": false}).Response(c)
	return
}
