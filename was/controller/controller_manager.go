package controller

import (
	"github.com/codestates/WBA-BC-Project-02/was/controller/info"
	"github.com/codestates/WBA-BC-Project-02/was/controller/user"
)

var InfoControl info.InfoController
var UserControl user.UserController

func InjectControllerDependency() {
	InfoControl = info.NewInfoControl()

}
