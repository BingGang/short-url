package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) SISMEMBERS(k string,  v interface{}) (bool, error) {
	result, err := redis.Int(this.Do("SISMEMBERS", k, v))
	switch err {
	case nil:
	case redis.ErrNil:
		return false, nil
	default:
		return false, err
	}
	return result == 1, nil
}
