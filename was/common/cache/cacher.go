package cache

import "time"

type Cacher interface {
	Cache(key string, data any, du time.Duration) error

	Update(key string, data any, du time.Duration) error

	Delete(keys ...string) error

	Get(key string, t any) (any, error)
}
