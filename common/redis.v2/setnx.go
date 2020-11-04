package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) SETNX(k string, v interface{}) (int64, error) {
	nlen, err := redis.Int64(this.Do("SETNX", k, v))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return 0, nil
	default:
		return 0, err
	}
	return nlen, nil
}
