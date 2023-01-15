package cache

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	wasError "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/gofrs/uuid"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Tokens struct {
	AccessToken  TokenDetail
	RefreshToken TokenDetail
}

type TokenDetail struct {
	CacheID  string
	TokenID  string
	Token    string
	Duration int64
}

func CreateToken(userID, accessKey, refreshKey string) (*Tokens, error) {
	//Creating Access Token
	acToken, err := createAccessToken(userID, accessKey)
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rfToken, err := createRefreshToken(userID, refreshKey)
	if err != nil {
		return nil, err
	}

	return &Tokens{
		AccessToken:  *acToken,
		RefreshToken: *rfToken,
	}, nil
}

func DecryptToken(signedToken, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(
		signedToken,
		func(JWTtoekn *jwt.Token) (interface{}, error) {
			if _, ok := JWTtoekn.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, wasError.UnauthorizedError
			}
			return []byte(secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenUUID(token *jwt.Token, filed string) (string, string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	cacheID := ""
	tokenID := ""
	if ok && token.Valid {
		cacheID, ok = claims[filed].(string)
		if !ok {
			return "", "", wasError.UnauthorizedError
		}
		tokenID, ok = claims[enum.JWTTokenID].(string)
		if !ok {
			return "", "", wasError.UnauthorizedError
		}
	}
	return cacheID, tokenID, nil
}

func createAccessToken(userID, accessKey string) (*TokenDetail, error) {
	td := &TokenDetail{}
	td.Duration = time.Now().Add(time.Minute * 15).Unix()
	td.CacheID = enum.AccessToken + userID
	UUID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	td.TokenID = UUID.String()

	atClaims := jwt.MapClaims{}
	atClaims[enum.JWTAuthorized] = true
	atClaims[enum.JWTAccessUUID] = td.CacheID
	atClaims[enum.JWTTokenID] = td.TokenID
	atClaims[enum.JWTEXP] = td.Duration
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	signedToken, err := at.SignedString([]byte(accessKey))
	if err != nil {
		return nil, err
	}
	td.Token = signedToken
	return td, nil
}

func createRefreshToken(userID, refreshKey string) (*TokenDetail, error) {
	td := &TokenDetail{}
	td.Duration = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.CacheID = enum.RefreshToken + userID
	UUID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	td.TokenID = UUID.String()

	rtClaims := jwt.MapClaims{}
	rtClaims[enum.JWTRefreshUUID] = td.CacheID
	rtClaims[enum.JWTTokenID] = td.TokenID
	rtClaims[enum.JWTEXP] = td.Duration
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	signedToken, err := rt.SignedString([]byte(refreshKey))
	if err != nil {
		return nil, err
	}
	td.Token = signedToken
	return td, nil
}
