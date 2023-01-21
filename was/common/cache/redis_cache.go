package cache

import (
	"encoding/json"
	"errors"
	wasCommon "github.com/codestates/WBA-BC-Project-02/was/common"
	wasError "github.com/codestates/WBA-BC-Project-02/was/common/error"
	"github.com/go-redis/redis/v9"
	"time"
)

var Redis *RedisProxy

type RedisProxy struct {
	cache Cacher
}

type RedisData interface {
	MarshalBinary() (data []byte, err error)
}

func LoadRedis(DNS string) error {
	if Redis != nil {
		return nil
	}

	rCache, err := newRedisCache(DNS)
	if err != nil {
		return err
	}
	Redis = &RedisProxy{cache: rCache}
	return nil
}

func (r *RedisProxy) Cache(key string, data RedisData, du time.Duration) error {
	if err := r.cache.Cache(key, data, du); err != nil {
		return err
	}
	return nil
}

// Update duration 값이 -1 이면 기존의 TTL(Time To Live) 유지, 0 은 지속 유지
func (r *RedisProxy) Update(key string, data RedisData, du time.Duration) error {
	if err := r.cache.Update(key, data, du); err != nil {
		return err
	}
	return nil
}

func (r *RedisProxy) Delete(keys ...string) error {
	if err := r.cache.Delete(keys...); err != nil {
		return err
	}
	return nil
}

func (r *RedisProxy) Get(key string, t RedisData) (any, error) {
	data, err := r.cache.Get(key, t)
	if err != nil {
		return nil, err
	}
	return data, nil
}

var instance *redisCache

type redisCache struct {
	client *redis.Client
}

func newRedisCache(DNS string) (*redisCache, error) {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	if instance == nil {
		instance = &redisCache{
			client: redis.NewClient(&redis.Options{Addr: DNS}),
		}
	}

	if _, err := instance.client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return instance, nil
}

func (rc *redisCache) Cache(key string, data any, du time.Duration) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	if err := rc.client.Set(ctx, key, data, du).Err(); err != nil {
		return err
	}
	return nil
}

// Update duration 값이 -1 이면 기존의 TTL(Time To Live) 유지, 0 은 지속 유지
func (rc *redisCache) Update(key string, data any, du time.Duration) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	if err := rc.client.Set(ctx, key, data, du).Err(); err != nil {
		return err
	}
	return nil
}

func (rc *redisCache) Delete(keys ...string) error {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	if count, err := rc.client.Del(ctx, keys...).Result(); err != nil || count == 0 {
		return errors.New(wasError.RedisDelZeroCount + err.Error())
	}
	return nil

}

func (rc *redisCache) Get(key string, t any) (any, error) {
	ctx, cancel := wasCommon.NewContext(wasCommon.ServiceContextTimeOut)
	defer cancel()

	data, err := rc.client.Get(ctx, key).Result()
	if err != nil {
		return nil, errors.New(wasError.RedisGetError + err.Error())
	}

	if err := json.Unmarshal([]byte(data), t); err != nil {
		return nil, err
	}

	return t, nil
}
