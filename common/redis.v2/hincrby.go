package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) HINCRBY(k string, v interface{},n int64) (int64, error) {
	result, err := redis.Int64(this.Do("HINCRBY", k,v, n))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return -1, nil
	default:
		return 0, err
	}
	return result, nil
}
