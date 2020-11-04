package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) INCR(k string) (int64, error) {
	result, err := redis.Int64(this.Do("INCR", k))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return -1, nil
	default:
		return 0, err
	}
	return result, nil
}
