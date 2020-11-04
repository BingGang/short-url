package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (this *Redis) SMEMBERS(k string) ([]interface{}, error) {
	buf, err := redis.Values(this.Do("SMEMBERS", k))
	switch err {
	case nil:
	case redis.ErrNil:
		return nil, nil
	default:
		return nil, err
	}
	return buf, nil
}
