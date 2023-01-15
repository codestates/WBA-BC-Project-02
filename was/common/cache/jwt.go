package cache

import (
	"github.com/codestates/WBA-BC-Project-02/was/common/enum"
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

func createAccessToken(userID, accessKey string) (*TokenDetail, error) {
	td := &TokenDetail{}
	td.Duration = time.Now().Add(time.Minute * 15).Unix()
	td.ID = enum.AccessToken + userID

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.ID
	atClaims["user_id"] = userID
	atClaims["exp"] = td.Duration
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
	rtClaims["refresh_uuid"] = td.ID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.Duration
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	signedToken, err := rt.SignedString([]byte(refreshKey))
	if err != nil {
		return nil, err
	}
	td.Token = signedToken
	return td, nil
}
