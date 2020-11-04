package redis

import (
	"github.com/garyburd/redigo/redis"
	"strings"
)

func (this *Redis) SET(k string, v interface{}) (bool, error) {
	result, err := redis.String(this.Do("SET", k, v))
	switch err {
	case nil:
	case redis.ErrPoolExhausted:
		return false, nil
	default:
		return false, err
	}

	return strings.EqualFold(result, "ok"), nil
}
