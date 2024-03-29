package error

const (
	Unauthorized      = "인증 오류"
	DifferentDevice   = "다른 환경에서 로그인 시도"
	DifferentToken    = "변조된 인증 시도"
	RestAccessFail    = "로그인 필요"
	BadRequest        = "부적절한 요청 오류"
	DataNotFound      = "존재하지 않는 데이터"
	DuplicatedData    = "이미 존재하는 데이터"
	UserNotFound      = "가입되지 않은 사용자"
	RedisDelZeroCount = "삭제된 캐시가 없음"
	RedisGetError     = "캐시에서 찾을 수 없음"

	InternalServerErr = "서버 로직 오류"
	FailRequestErr    = "시스템 오류"
	UnKnown           = "알 수 없는 오류"

	NonInjected = "의존성 주입 오류"
)

const (
	UnauthorizedCode     = 401
	DifferentRequestCode = 404
	RestAccessFailCode   = 442
	BadRequestCode       = 443
	DataNotFoundCode     = 444
	DuplicatedDataCode   = 445
	UserNotFoundCode     = 446

	InternalServerErrCode = 551
	SystemErrCode         = 552
	UnKnownCode           = 553

	NonInjectedCode = 554
)
