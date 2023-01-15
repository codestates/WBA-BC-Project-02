package cache

import (
	"fmt"
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
	wasError "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	AccessToken  TokenDetail
	RefreshToken TokenDetail
}

type TokenDetail struct {
	ID       string
	Token    string
	Duration int64
}

func CreateToken(userID, accessKey, refreshKey string) (*Token, error) {
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

	return &Token{
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

func ExtractAccessTokenUUID(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println("##### ok 와 claims :: ", ok, claims)
	accessUUID := ""
	if ok && token.Valid {
		accessUUID, ok = claims[enum.JWTAccessUUID].(string)
		if !ok {
			return "", wasError.UnauthorizedError
		}
	}
	return accessUUID, nil
}

func createAccessToken(userID, accessKey string) (*TokenDetail, error) {
	td := &TokenDetail{}
	td.Duration = time.Now().Add(time.Minute * 15).Unix()
	td.ID = enum.AccessToken + userID

	atClaims := jwt.MapClaims{}
	atClaims[enum.JWTAuthorized] = true
	atClaims[enum.JWTAccessUUID] = td.ID
	atClaims[enum.JWTUserID] = userID
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
	td.ID = enum.RefreshToken + userID

	rtClaims := jwt.MapClaims{}
	rtClaims[enum.JWTRefreshUUID] = td.ID
	rtClaims[enum.JWTUserID] = userID
	rtClaims[enum.JWTEXP] = td.Duration
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	signedToken, err := rt.SignedString([]byte(refreshKey))
	if err != nil {
		return nil, err
	}
	td.Token = signedToken
	return td, nil
}
