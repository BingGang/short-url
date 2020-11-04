package redis

import (
	"github.com/garyburd/redigo/redis"
)

func  (this *Redis) HDEL(k string, f ...interface{}) (bool, error) {
	arg := []interface{}{k}
	arg = append(arg,f...)
	result, err := redis.Int(this.Do("HDEL",arg...))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return false, nil
	default:
		return false, err
	}

	return result == 1, nil
}