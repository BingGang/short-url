package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) EXISTS(k string) (bool, error) {
	result, err := redis.Int(this.Do("EXISTS", k))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return false, nil
	default:
		return false, err
	}

	return result == 1, nil
}
