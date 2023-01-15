package cache

import (
	"encoding/json"
	"errors"
	wasCommon "github.com/codestates/WBA-BC-Project-02/was/common"
	error2 "github.com/codestates/WBA-BC-Project-02/was/common/error"
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

func CacheLoginInfo(loginInfo *LoginInformation, token *Token) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	at := time.Unix(token.AccessToken.Duration, 0)
	rt := time.Unix(token.RefreshToken.Duration, 0)
	now := time.Now()

	if err := client.Set(ctx, token.AccessToken.ID, loginInfo, at.Sub(now)).Err(); err != nil {
		return err
	}
	if err := client.Set(ctx, token.RefreshToken.ID, loginInfo, rt.Sub(now)).Err(); err != nil {
		return err
	}

	return nil
}

func Delete(keys ...string) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	if count, err := client.Del(ctx, keys...).Result(); err != nil || count == 0 {
		return errors.New(error2.RedisDelZeroCount + err.Error())
	}
	return nil
}

func GetLoginInfo(key string) (*LoginInformation, error) {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	loginInfoSTR, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, errors.New(error2.RedisGetError + err.Error())
	}

	info := &LoginInformation{}
	if err := json.Unmarshal([]byte(loginInfoSTR), info); err != nil {
		return nil, err
	}
	return info, nil
}
