package error

import (
	"errors"
)

type AppError struct {
	Code int
	MSG  string
	Name string
}

func (e AppError) Error() string {
	return e.MSG
}

func (e AppError) New() error {
	return errors.New(e.MSG)
}

func NewAppErrorAndMSG(MSG string) AppError {
	return NewAppErrorAndCode(UnKnownCode, MSG)
}

func NewAppError(e error) AppError {
	return NewAppErrorAndCode(UnKnownCode, e.Error())
}

func NewAppErrorAndCode(code int, MSG string) AppError {
	return NewAppErrorCustom(code, MSG, FailRequestErr)
}

func NewAppErrorCustom(code int, MSG string, name string) AppError {
	return AppError{
		Code: code,
		MSG:  MSG,
		Name: name,
	}
}

/* 401 ~ 499 error status */
var (
	UnauthorizedError    = NewAppErrorCustom(UnauthorizedCode, "허가되지 않은 사용자 입니다.", Unauthorized)
	DifferentDeviceError = NewAppErrorCustom(DifferentRequestCode, "이전에 사용된 환경과 다릅니다. 인증을 재시도 해주세요.", DifferentDevice)
	DifferentTokenID     = NewAppErrorCustom(DifferentRequestCode, "인증이 변조된 것 같습니다. 인증을 재시도 해주세요.", DifferentToken)
	RestAccessFailError  = NewAppErrorCustom(RestAccessFailCode, "로그인이 필요합니다.", RestAccessFail)
	BadRequestError      = NewAppErrorCustom(BadRequestCode, "부적절한 요청입니다.", BadRequest)
	DataNotFoundError    = NewAppErrorCustom(DataNotFoundCode, "데이터를 찾을 수 없습니다.", DataNotFound)
	DuplicatedDataError  = NewAppErrorCustom(DuplicatedDataCode, "이미 존재하는 데이터 입니다.", DuplicatedData)
	BadUserAgentError    = NewAppErrorCustom(BadRequestCode, "식별할 수 없는 요청입니다.", BadRequest)
	UserNotFoundError    = NewAppErrorCustom(UserNotFoundCode, "찾을 수 없는 사용자 입니다.", UserNotFound)
)

/* 501 ~ 599 서버 에러 */
var (
	NonInjectedError    = NewAppErrorCustom(InternalServerErrCode, "의존성 주입이 이루어지지 않았습니다.", NonInjected)
	InternalServerError = NewAppErrorCustom(SystemErrCode, "서버 로직을 수행하지 못했습니다.", InternalServerErr)
	FailRequestError    = NewAppErrorCustom(UnKnownCode, "요청을 정상적으로 처리하지 못했습니다.", FailRequestErr)
	UnKnownError        = NewAppErrorCustom(NonInjectedCode, "알 수 없는 오류입니다.", UnKnown)
)
