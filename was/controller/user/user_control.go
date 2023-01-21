package user

import (
	"github.com/codestates/WBA-BC-Project-02/common/util/validator"
	"github.com/codestates/WBA-BC-Project-02/was/common/cache/login"
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
	setCookieWithRefreshToken(resMnemonic.Token.RefreshToken, c)
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

	setCookieWithRefreshToken(token.RefreshToken, c)
	protocol.SuccessData(token).Response(c)
}

func (u *userControl) ReissueToken(c *gin.Context) {
	refreshToken, err := c.Cookie(enum.CookieRefreshToken)
	if err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}
	if err := validator.CheckBlank(refreshToken); err != nil {
		protocol.Fail(wasError.BadRequestError).Response(c)
		return
	}

	ua := c.GetHeader(enum.HeaderUserAgent)

	token, err := u.userService.ReissueToken(refreshToken, ua)
	if err != nil {
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}

	setCookieWithRefreshToken(token.RefreshToken, c)
	protocol.SuccessData(token).Response(c)
}

func (u *userControl) GetUserSimpleInformation(c *gin.Context) {
	loginInfo, exists := c.Keys[enum.LoginInformation].(*login.Information)
	if !exists {
		protocol.Fail(wasError.InternalServerError).Response(c)
		return
	}
	userInfo := response.FromCache(loginInfo)
	resU, err := u.userService.GetSimpleUser(userInfo.Address)
	if err != nil {
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(resU).Response(c)
}

func (u *userControl) GetUserInformation(c *gin.Context) {
	loginInfo, exists := c.Keys[enum.LoginInformation].(*login.Information)
	if !exists {
		protocol.Fail(wasError.InternalServerError).Response(c)
		return
	}

	userInfo := response.FromCache(loginInfo)
	resU, err := u.userService.GetUser(userInfo.Address)
	if err != nil {
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(resU).Response(c)
}

func (u *userControl) IncreaseBlackIron(c *gin.Context) {
	loginInfo, exists := c.Keys[enum.LoginInformation].(*login.Information)
	if !exists {
		protocol.Fail(wasError.InternalServerError).Response(c)
		return
	}

	simpleUser, err := u.userService.IncreaseBlackIron(loginInfo)
	if err != nil {
		protocol.Fail(wasError.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(simpleUser).Response(c)
}

func setCookieWithRefreshToken(refreshToken string, c *gin.Context) {
	c.SetCookie(
		enum.CookieRefreshToken,
		refreshToken,
		enum.CookieRefreshTokenDuration,
		enum.CookiePath,
		enum.CookieDomain,
		enum.CookieSecure,
		enum.CookieHttpOnly,
	)
}
