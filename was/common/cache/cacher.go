package cache

import "time"

//var AppCacher Cacher

type Cacher interface {
	Cache(key string, data any, du time.Duration) error

	Update(key string, data any, du time.Duration) error

	Delete(keys ...string) error

	Get(key string, t any) (any, error)
}

//func SetAppCacher(cache Cacher) {
//	AppCacher = cache
//}
