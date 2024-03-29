package controller

import (
	"github.com/codestates/WBA-BC-Project-02/was/controller/contract"
	"github.com/codestates/WBA-BC-Project-02/was/controller/info"
	"github.com/codestates/WBA-BC-Project-02/was/controller/user"
	"github.com/codestates/WBA-BC-Project-02/was/service"
)

var InfoControl info.InfoController
var UserControl user.UserController
var ContractControl contract.ContractController

func InjectControllerDependency() {
	InfoControl = info.NewInfoControl()
	UserControl = user.NewUserControl(service.UserService)
	ContractControl = contract.NewContractControl(service.ContractService, service.WemixService)
}
