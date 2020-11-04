package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) LLEN(k string) (int64, error) {

	value, err := redis.Int64(this.Do("LLEN", k))
	switch err {
	case nil:
	case redis.ErrNil:
		return 0, nil
	default:
		return 0, err
	}
	return value, nil
}
