package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) INCRBY(k string, n int64) (int64, error) {
	result, err := redis.Int64(this.Do("INCRBY", k, n))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return -1, nil
	default:
		return 0, err
	}
	return result, nil
}
