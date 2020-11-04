package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) LRANGE(k string, s, e interface{}) ([]interface{}, error) {
	buf, err := redis.Values(this.Do("LRANGE", k, s, e))
	switch err {
	case nil:
	case redis.ErrNil:
		return nil, nil
	default:
		return nil, err
	}
	return buf, nil
}
