package cache

import (
	wasCommon "github.com/codestates/WBA-BC-Project-02/was/common"
	"github.com/go-redis/redis/v9"
	"time"
)

var client *redis.Client

func LoadRedisClient(DNS string) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr: DNS,
		})
	}

	if _, err := client.Ping(ctx).Result(); err != nil {
		return err
	}

	return nil
}

func GetClient() *redis.Client {
	return client
}

func CacheLoginInfo(loginInfo *LoginInformation, token *wasCommon.Token) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	at := time.Unix(token.AccessToken.Duration, 0)
	rt := time.Unix(token.RefreshToken.Duration, 0)
	now := time.Now()

	if err := client.Set(ctx, token.AccessToken.UUID, loginInfo, at.Sub(now)).Err(); err != nil {
		return err
	}
	if err := client.Set(ctx, token.RefreshToken.UUID, loginInfo, rt.Sub(now)).Err(); err != nil {
		return err
	}

	return nil
}
