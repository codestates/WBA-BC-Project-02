package user

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/cache"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	wasError "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/protocol"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/request"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/response"
	"github.com/codestates/WBA-BC-Project-02/was/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
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
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

	ua := c.GetHeader(enum.HeaderUserAgent)

	resMnemonic, err := u.userService.CreateWallet(reqP.Password, ua)
	if err != nil {
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessCodeAndData(http.StatusCreated, resMnemonic).Response(c)
}

func (u *userControl) RecoverWallet(c *gin.Context) {
	reqR := &request.Recovery{}
	if err := c.ShouldBindJSON(reqR); err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

	ua := c.GetHeader(enum.HeaderUserAgent)

	token, err := u.userService.RecoverWallet(reqR, ua)
	if err != nil {
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(token).Response(c)
}

func (u *userControl) ReissueToken(c *gin.Context) {
	reqR := &request.ReissueToken{}
	if err := c.ShouldBindJSON(reqR); err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

	ua := c.GetHeader(enum.HeaderUserAgent)

	token, err := u.userService.ReissueToken(reqR.RefreshToken, ua)
	if err != nil {
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(token).Response(c)
}

func (u *userControl) GetUserSimpleInformation(c *gin.Context) {
	value, exists := c.Keys[enum.LoginInformation].(*cache.LoginInformation)
	if !exists {
		protocol.Fail(wasError.InternalServerError).Response(c)
		return
	}
	userInfo := response.FromCache(value)
	protocol.SuccessData(userInfo).Response(c)
}

func (u *userControl) GetUserInformation(c *gin.Context) {

}
