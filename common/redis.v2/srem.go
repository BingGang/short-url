package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) SREM(k string, v interface{}) (bool, error) {
	result, err := redis.Int(this.Do("SREM", k, v))
	switch err {
	case nil:
	case redis.ErrNil:
		return false, nil
	default:
		return false, err
	}
	return result == 1, nil
}
