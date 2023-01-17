package user

import (
	"github.com/codestates/WBA-BC-Project-02/common/util/validator"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	error2 "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/protocol"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/request"
	"github.com/codestates/WBA-BC-Project-02/was/service/user"
	"github.com/gin-gonic/gin"
)

var instance *userControl

type userControl struct {
	userService user.UserServicer
}

func NewUserControl(svc user.UserServicer) *userControl {
	if instance != nil {
		return instance
	}
	instance = &userControl{
		userService: svc,
	}
	return instance
}

func (u *userControl) CreateWallet(c *gin.Context) {
	reqP := &request.Password{}
	if err := c.ShouldBindJSON(reqP); err != nil {
		protocol.Fail(error2.BadRequestError).Response(c)
		return
	}

	device := c.GetHeader(enum.HeaderUserAgent)
	if err := validator.CheckBlank(device); err != nil {
		protocol.Fail(error2.BadRequestError).Response(c)
		return
	}

	resMnemonic, err := u.userService.CreateWallet(reqP.Password, device)
	if err != nil {
		protocol.Fail(error2.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(resMnemonic).Response(c)
}

func (u *userControl) BringWallet(c *gin.Context) {

}
