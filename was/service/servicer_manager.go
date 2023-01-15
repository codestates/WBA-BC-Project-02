package service

import (
	"github.com/codestates/WBA-BC-Project-02/was/model"
	"github.com/codestates/WBA-BC-Project-02/was/service/contract"
	"github.com/codestates/WBA-BC-Project-02/was/service/user"
)

var UserService user.UserServicer
var ContractService contract.ContractServicer

func InjectServicesDependency() {
	UserService = user.NewUserService(model.UserModel)
	ContractService = contract.NewContractService(model.ContractModel)
}
