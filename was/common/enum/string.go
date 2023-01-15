package enum

const (
	// WEB
	HeaderUserAgent     = "User-Agent"
	HeaderAuthorization = "Authorization"
	LoginInformation    = "loginInformation"

	// JWT TOKEN
	JWTAuthorized  = "authorized"
	JWTAccessUUID  = "access_uuid"
	JWTRefreshUUID = "refresh_uuid"
	JWTTokenID     = "token_id"
	JWTEXP         = "exp"

	// cache
	AccessCache  = "ACCESSTOEKN"
	RefreshCache = "REFRESHTOEKN"
)
