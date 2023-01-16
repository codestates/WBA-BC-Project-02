package enum

import "time"

const (
	// JWT TOKEN
	JWTAuthorized = "authorized"
	JWTAccessID   = "access_uuid"
	JWTRefreshID  = "refresh_uuid"
	JWTTokenID    = "token_id"
	JWTEXP        = "exp"

	JWTAccessDuration  = time.Minute * 15
	JWTRefreshDuration = time.Hour * 24 * 7
)
