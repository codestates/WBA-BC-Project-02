package cache

import (
	"encoding/json"
	"errors"
	wasCommon "github.com/codestates/WBA-BC-Project-02/was/common"
	wasError "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/go-redis/redis/v9"
	"time"
)

var instance *RedisCache

type RedisCache struct {
	client *redis.Client
}

type RedisData interface {
	MarshalBinary() (data []byte, err error)
}

func NewRedisCache(DNS string) (*RedisCache, error) {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	if instance == nil {
		instance = &RedisCache{
			client: redis.NewClient(&redis.Options{Addr: DNS}),
		}
	}

	if _, err := instance.client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (r *RedisCache) Cache(key string, data any, du time.Duration) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	if err := r.client.Set(ctx, key, data, du).Err(); err != nil {
		return err
	}
	return nil
}

// Update duration 값이 -1 이면 기존의 TTL(Time To Live) 유지, 0 은 지속 유지
func (r *RedisCache) Update(key string, data any, du time.Duration) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	if err := r.client.Set(ctx, key, data, du).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisCache) Delete(keys ...string) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	if count, err := r.client.Del(ctx, keys...).Result(); err != nil || count == 0 {
		return errors.New(wasError.RedisDelZeroCount + err.Error())
	}
	return nil

}

func (r *RedisCache) Get(key string, t any) (any, error) {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	data, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return nil, errors.New(wasError.RedisGetError + err.Error())
	}

	if err := json.Unmarshal([]byte(data), t); err != nil {
		return nil, err
	}

	return t, nil
}
