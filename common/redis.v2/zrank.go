package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) ZRANK(k string, v interface{}) (int64, error) {
	result, err := redis.Int64(this.Do("ZRANK", k, v))
	switch err {
	case nil:
	case redis.ErrNil:
		return -1, nil
	default:
		return 0, err
	}
	return result, nil
}
