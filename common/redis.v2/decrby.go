package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) DECRBY(k string, n int64) (int64, error) {
	result, err := redis.Int64(this.Do("DECR", k, n))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return 0, nil
	default:
		return 0, err
	}
	return result, nil
}
