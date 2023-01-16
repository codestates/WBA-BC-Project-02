package enum

import "time"

const (
	// JWT TOKEN
	JWTAuthorized = "authorized"
	JWTAccessID   = "access_id"
	JWTRefreshID  = "refresh_id"
	JWTTokenID    = "token_id"
	JWTEXP        = "exp"

	JWTAccessDuration  = time.Minute * 15
	JWTRefreshDuration = time.Hour * 24 * 7
)
