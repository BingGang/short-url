package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) DECR(k string) (int64, error) {
	result, err := redis.Int64(this.Do("DECR", k))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return 0, nil
	default:
		return 0, err
	}
	return result, nil
}
