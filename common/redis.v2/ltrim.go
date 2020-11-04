package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) LTRIM(k string, s, e interface{}) ([]interface{}, error) {
	buf, err := redis.Values(this.Do("LTRIM", k, s, e))
	switch err {
	case nil:
	case redis.ErrNil:
		return nil, nil
	default:
		return nil, err
	}
	return buf, nil
}
