package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis)ZREMRANGEBYSCORE (k string,v...interface{}) (bool, error) {
	arg := []interface{}{k}
	arg = append(arg,v...)
	result, err := redis.Int(this.Do("ZREMRANGEBYSCORE", arg...))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return false, nil
	default:
		return false, err
	}

	return result == 1, nil
}
