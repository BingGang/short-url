package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) HSET(k string, f interface{}, v interface{}) (bool, error) {
	result, err := redis.Int(this.Do("HSET", k, f, v))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return false, nil
	default:
		return false, err
	}

	return result == 1, nil
}
