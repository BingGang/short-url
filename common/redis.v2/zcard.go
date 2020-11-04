package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) ZCARD(k string) (int64, error) {
	result, err := redis.Int64(this.Do("ZCARD", k))
	switch err {
	case nil:
	case redis.ErrNil:
		return -1, nil
	default:
		return 0, err
	}
	return result, nil
}
