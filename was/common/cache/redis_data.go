package cache

type RedisData interface {
	MarshalBinary() (data []byte, err error)
}
