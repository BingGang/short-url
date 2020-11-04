package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) ZADD(k string, c, v interface{}) (bool, error) {
	result, err := redis.Int(this.Do("ZADD", k, c, v))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return false, nil
	default:
		return false, err
	}

	return result != 0, nil
}
