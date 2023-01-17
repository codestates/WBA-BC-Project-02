package contract

import (
	"fmt"
	"github.com/codestates/WBA-BC-Project-02/common/model/entity"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	"github.com/codestates/WBA-BC-Project-02/was/config"
	"github.com/codestates/WBA-BC-Project-02/was/model/contract"
	contract2 "github.com/codestates/WBA-BC-Project-02/was/protocol/contract"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/contract/response"
	"math"
	"strconv"
)

type contractService struct {
	contactModel     contract.ContractModeler
	dexContractModel contract.DexContractModeler
}

var instance *contractService

func NewContractService(con contract.ContractModeler, dex contract.DexContractModeler) *contractService {
	if instance != nil {
		return instance
	}
	instance = &contractService{
		contactModel:     con,
		dexContractModel: dex,
	}
	return instance
}

func (c *contractService) GetContractByName(contractName string) (*response.Contract, error) {
	foundCon, err := c.contactModel.FindContractByName(contractName)
	if err != nil {
		return nil, err
	}

	resC := response.FromContractEntity(foundCon)
	return resC, nil
}

func (c *contractService) GetContracts() ([]*response.SimpleContract, error) {
	foundsContracts, err := c.contactModel.FindNonTxContracts()
	if err != nil {
		return nil, err
	}
	simpleContracts := response.FromSimpleContracts(foundsContracts)
	return simpleContracts, nil
}

func (c *contractService) GetRatioTokenAndCredit(contractName string) (*response.RatioContract, error) {
	dexContract, err := c.dexContractModel.FindDexContract(config.ContractConfig.DexAddr)
	if err != nil {
		return nil, err
	}

	return getRatioContract(dexContract, contractName)
}

func getRatioContract(dexContract *entity.DexContract, name string) (*response.RatioContract, error) {
	switch name {
	case enum.DracoContractName:
		return createResponseRatioContract(dexContract.DracoPoolToken, dexContract.DracoPoolCredit)
	case enum.TigContractName:
		return createResponseRatioContract(dexContract.TigPoolToken, dexContract.TigPoolCredit)
	}
	return nil, nil
}

func createResponseRatioContract(token, credit string) (*response.RatioContract, error) {
	price, err := getPrice(token, credit)
	if err != nil {
		return nil, err
	}
	priceSTR := fmt.Sprintf("%f", price)
	ratio := &contract2.TokenAndCredit{
		Token: priceSTR, Credit: enum.RatioCreditDefaultValue,
	}
	detail := &contract2.TokenAndCredit{
		Token: token, Credit: credit,
	}
	return &response.RatioContract{
		Price:  priceSTR,
		Ratio:  ratio,
		Detail: detail,
	}, nil
}

func getPrice(token, credit string) (float64, error) {
	tokenF, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return 0, err
	}
	creditF, err := strconv.ParseFloat(credit, 64)
	if err != nil {
		return 0, err
	}

	price := math.Round((tokenF/creditF)*10000) / 10000
	return price, nil
}
