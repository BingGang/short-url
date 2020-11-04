package redis

import (
	"github.com/garyburd/redigo/redis"
	"strings"
)

func (this *Redis) SETEX(k string, v interface{}, ttl int64) (bool, error) {
	result, err := redis.String(this.Do("SETEX", k, ttl, v))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return false, nil
	default:
		return false, err
	}

	return strings.EqualFold(result, "ok"), nil
}
