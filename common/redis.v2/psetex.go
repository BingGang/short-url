package redis

import (
	"github.com/garyburd/redigo/redis"
	"strings"
)

func (this *Redis) PSETEX(k string, v interface{}, ttl int) (bool, error) {
	result, err := redis.String(this.Do("PSETEX", k, ttl, v))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return false, nil
	default:
		return false, err
	}

	return strings.EqualFold(result, "ok"), nil
}
