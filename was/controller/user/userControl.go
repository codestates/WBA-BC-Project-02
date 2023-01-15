package user

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	error2 "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/codestates/WBA-BC-Project-02/was/protocol"
	"github.com/codestates/WBA-BC-Project-02/was/protocol/user/request"
	"github.com/codestates/WBA-BC-Project-02/was/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO 비밀번호는 bcrypt 쓰고 있는데 이는 항상 해시값이 달라 찾기힘들다
// 	생각해보니 메타마스크는 동일한 비밀번호로 게속 계정생성이 가능하다. 즉 유니크가아니다
// 	내 생각에는 지갑주소로 찾아 비밀번호로 일치하는지 보는것 같다
// 	리프레시 토큰을 7일로 설정해 사용자는 비밀키로 로그인을 따로하지 않아도 계정주소를 7일간 리프레시 토큰안에 갖고 있을것임으로 JWT 를 들고있다면 비밀번호만으로 가능
// 	그렇기에 추후 요청 UserAgent 를 저장해 장치가 조금이라도 바뀌면 JWT 를 폐기하도록 하자. 또한 암호도 좀더 강력한 Argon 방식으로 바꾸는게 나을꺼 같기도하다.

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

	ua := c.GetHeader(enum.HeaderUserAgent)

	resMnemonic, err := u.userService.CreateWallet(reqP.Password, ua)
	if err != nil {
		protocol.Fail(error2.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessCodeAndData(http.StatusCreated, resMnemonic).Response(c)
}

func (u *userControl) RecoverWallet(c *gin.Context) {
	reqR := &request.Recovery{}
	if err := c.ShouldBindJSON(reqR); err != nil {
		protocol.Fail(error2.BadRequestError).Response(c)
		return
	}

	ua := c.GetHeader(enum.HeaderUserAgent)

	token, err := u.userService.RecoverWallet(reqR, ua)
	if err != nil {
		protocol.Fail(error2.NewAppError(err)).Response(c)
		return
	}

	protocol.SuccessData(token).Response(c)
}
